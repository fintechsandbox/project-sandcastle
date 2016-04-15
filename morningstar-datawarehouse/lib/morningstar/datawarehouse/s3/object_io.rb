require 'aws-sdk'
require 'zip'
require 'tempfile'

module Morningstar
  module Datawarehouse
    module S3
      #TODO: This class was implemented just in testing purpose. It is unstable and not tested. So, use it on own risk
      class ObjectIO < ::IO

        def initialize(bucket, path)
          config = Morningstar::Datawarehouse::Configuration.options
          @bucket = bucket
          @path = path
          @pos = 0
          @s3 = Aws::S3::Client.new(
              access_key_id: config.aws_access_key,
              secret_access_key: config.aws_access_secret,
              region: config.aws_region
          )

          @content_length = @s3.head_object(bucket: @bucket, key: @path).content_length
        end


        def seek(amount, whence=SEEK_SET)
          case whence
            when IO::SEEK_CUR then
              @pos += amount
            else
              @pos = amount
          end
        end

        def read(length = nil, buff = nil)
          return '' if length == 0
          range = to_range(length)
          @pos += length ? length : 0
          #TODO: this is workaround. Need to use tmp file fr chunks because StringIO will corrupt output in other case
          tmp = Tempfile.new('foo')
          r = @s3.get_object(bucket: @bucket,
                             key: @path, range: range, response_target: tmp)

          tmp.rewind
          result = tmp.read
          tmp.close
          tmp.unlink
          result
        end

        def rewind
          @pos = 0
        end

        def pos
          @pos
        end

        def binmode

        end

        def dup
          ObjectIO.new(@bucket, @path)
        end

        def binmode?
          true
        end

        def path
          self
        end

        def tell
          pos
        end

        def reopen(path)
          raise IOError, 'not opened for writing'
        end

        private
        def to_range(length)
          if @pos < 0
            "bytes=#{@content_length + @pos}-#{@content_length - 1}"
          elsif length.nil?
            "bytes=#{@pos}-#{@content_length - 1}"
          else
            "bytes=#{@pos}-#{length ? @pos + length - 1 : ''}"
          end
        end
      end
    end
  end
end
