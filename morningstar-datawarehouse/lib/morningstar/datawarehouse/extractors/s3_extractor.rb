require 'zlib'
require 'zip'
require 'tempfile'

module Morningstar
  module Datawarehouse
    module Extractors
      class S3Extractor
        def initialize(bucket_name, s3_client)
          @bucket_name = bucket_name
          @client = s3_client
        end

        def list_objects
          @client.list_objects({bucket: @bucket_name})
        end

        def translate_object(key, &block)

          if ["zip", "gzip", "gz"].include? key.split('.').last
            if key.split('.').last == "zip"
              get_zip(key) do |io|
                block.call(io)
              end
            else
              print "Start load #{key} \n"
              get_gzip(key) do |chunk|
                block.call(chunk)
              end
              print "End LOAD \n"
            end
          else
            @client.get_object(bucket: @bucket_name, key: key) do |chunk|
              block.call(chunk)
            end
          end
        end

        def get_zip(key, &block)
          tmp_zip = Tempfile.new('morningstar_data')
          begin
            @client.get_object(bucket: @bucket_name,
                               key: key, response_target: tmp_zip)
            tmp_zip.rewind
            Zip::File.open_buffer(tmp_zip) do |zip_file|
              zip_file.each do |entry|
                block.call(entry.get_input_stream) if entry.file?
              end
            end
          ensure
            tmp_zip.close(true)
          end
        end


        def get_gzip(key, &block)
          reader, writer = IO.pipe
          fork do
            reader.close

            @client.get_object(bucket: @bucket_name, key: key) do |chunk|
              writer.write chunk
            end
          end

          writer.close

          gz = Zlib::GzipReader.new(reader)

          while line = gz.gets
            block.call(line)
          end

          gz.close
        end
      end
    end
  end
end
