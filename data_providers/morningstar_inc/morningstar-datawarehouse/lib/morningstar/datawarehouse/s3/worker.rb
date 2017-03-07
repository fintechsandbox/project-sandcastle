require 'aws-sdk'

module Morningstar
  module Datawarehouse
    module S3
      class  Worker
        def initialize(bucket)
          config = Morningstar::Datawarehouse::Configuration.options

          @client = Aws::S3::Client.new(
              credentials: Aws::Credentials.new(config.aws_access_key, config.aws_access_secret),
              region: config.aws_region)

          @resource = Aws::S3::Resource.new(client: @client)

          @bucket = @resource.bucket(bucket)
          @index = 1
        end

        def client
          @client
        end

        def refresh_index
          @index = 1
        end

        def resume(key)
          @i_upload.part(@index).copy_from(copy_source: "#{@object.bucket_name}/#{key}")
          @index +=  1
        end

        def get_object
          @object
        end

        def object(path)
          @object = @bucket.object(path)
        end

        def is_object_exists?(object_path)
          @bucket.object(object_path).exists?
        end

        def initialize_upload
          @i_upload = @object.initiate_multipart_upload
          @index = 1
        end

        def complete_upload
          @i_upload.complete(compute_parts: true)
        end

        def upload(data)
          @i_upload.part(@index).upload(body: data, content_length: data.size)
          @index +=  1
        end
      end
    end
  end
end
