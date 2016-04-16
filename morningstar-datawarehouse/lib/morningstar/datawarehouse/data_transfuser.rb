module Morningstar
  module Datawarehouse
    class DataProcessingError < StandardError
      attr_reader :code

      def with_code(code)
        @code = code
        self
      end

      def with_backtrace(backtrace)
        set_backtrace(backtrace)
        self
      end

      def to_json
        {code: self.code, message: self.message, backtrace: self.backtrace.join("\n")}
      end

    end

    class DataDownloadError < DataProcessingError; end

    class DataUploadError < DataProcessingError; end

    class DataTransfuser

      def initialize(ftp_path, bucket)
        config = Morningstar::Datawarehouse::Configuration.options
        @logger = config.logger
        @ftp_path = ftp_path
        @ftp_crawler = Morningstar::Datawarehouse::FtpCrawler.new(ftp_path)
        @s3_worker = Morningstar::Datawarehouse::S3::Worker.new(bucket)
      end

      def logger
        @logger ||= ::Logger.new('morningstar_data_warehouse.log')
      end

      def transfer_to_s3(file, folder)
        file_path = "#{folder}/#{file}"
        object_path = "MorningStar/#{@ftp_path}/#{file_path}"

        object = @s3_worker.object(object_path)

        ftp_file_size = @ftp_crawler.size(file_path)
        min_chunk = 6 * 1024 * 1024

        if ftp_file_size > 0
          begin
            @s3_worker.initialize_upload
            unless object.exists?
              @s3_worker.refresh_index
              @ftp_crawler.load(file_path, min_chunk){|data|
                @s3_worker.upload(data)
              }
              @s3_worker.complete_upload
            else
              if object.content_length != @ftp_crawler.size(file_path)
                @s3_worker.object(object_path)
                @s3_worker.resume(@s3_worker.get_object.key)

                @ftp_crawler.restore(file_path, min_chunk, @s3_worker.get_object.content_length){|data|
                  @s3_worker.upload(data)
                }
                @s3_worker.complete_upload
              end
            end
          rescue Net::FTPConnectionError => error
            logger.error("FTP connection error #{error}")
            # Upload to S3 already uploaded part (to assemble part of file in s3)
            @s3_worker.complete_upload
          end
        end
      end

      def has_trigger?(files)
        files.select{ |i| i[/trigger/] }.any?
      end

      def get_date_file_name(file)
        file[/\d{8}/]
      end

      def group_by_date(files)
        files.group_by{|file| get_date_file_name(file)}
      end

      def transfer(files_type = '')
        list = @ftp_crawler.list(nil)
        list = list.slice(2, list.length)
        begin
          list.each do |el|
            folder = el.split(' ').last
            if @ftp_crawler.list(folder).size > 2

              files = @ftp_crawler.list_of_files("#{folder}/*#{files_type}*")
              group_files = group_by_date(files)
              files.each do |file|
                date = get_date_file_name(file)
                if has_trigger? group_files[date]
                  logger.info "Upload #{file}/#{folder} \n"
                  transfer_to_s3(file, folder)
                  logger.info "END #{file}/#{folder} \n"
                else
                  logger.info "#{date} doesn't have trgigger yet \n"
                end
              end
            end
          end
          @ftp_crawler.close
        rescue Net::FTPPermError => error
          if error.message == "550 No files found.\n"
            logger.error error
          end
          raise Morningstar::Datawarehouse::DataDownloadError.new(error.message).with_code('FTPPermError').with_backtrace(error.backtrace)
        rescue Net::FTPReplyError => error
          raise Morningstar::Datawarehouse::DataDownloadError.new(error.message).with_code('FTPReplyError').with_backtrace(error.backtrace)
        rescue Aws::S3::MultipartUploadError => error
          raise Morningstar::Datawarehouse::DataUploadError.new(error.message).with_code('S3MultipartUploadError').with_backtrace(error.backtrace)
        rescue Aws::Errors::ServiceError => error
          raise Morningstar::Datawarehouse::DataUploadError.new(error.message).with_code('AWSError').with_backtrace(error.backtrace)
        end
      end
    end
  end
end
