require 'net/ftp'
require 'aws-sdk'

module Morningstar
  module Datawarehouse
    class  FtpCrawler
      def init_ftp_connection
        @config = Morningstar::Datawarehouse::Configuration.options

        @ftp = Net::FTP.new(@config.ftp_source, @config.ftp_username, @config.ftp_password)

        @ftp.resume = true
        @ftp.passive = true
      end

      def initialize(source)
        init_ftp_connection
        @source = source

        @ftp.resume = true
        @ftp.passive = true
        @ftp.login(@config.ftp_username, @config.ftp_password)
        @ftp.chdir(source)
      end

      def list_of_files(folder)
        @ftp.nlst(folder)
      end

      def list(folder)
        @ftp.list(folder)
      end

      def size(file_path)
        @ftp.size(file_path)
      end

      # TODO find better solution
      # if there are more then 7 chunks throw 421 error (421 Connection timed out - closing)
      # current solution reopen connection
      def load(file_path, min_chunk = 6 * 1024 * 1024, &block)
        index = 0
        @ftp.getbinaryfile(file_path, nil, min_chunk){|data|
          block.call(data)
          index += 1
        }

        if index > 9
          close
          init_ftp_connection
          @ftp.chdir(@source)
        end
      end

      # TODO find better solution
      # if there are more then 7 chunks throw 421 error (421 Connection timed out - closing)
      # current solution reopen connection
      def restore(file_path, min_chunk = 6 * 1024 * 1024, start = nil, &block)
        index = 0
        @ftp.retrbinary("RETR " + file_path.to_s, min_chunk, start) do |data|
          yield(data) if block_given?
        end
      end

      def close
        @ftp.close
      end

    end
  end
end
