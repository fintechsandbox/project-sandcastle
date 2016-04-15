require 'ostruct'

module Morningstar
  module Datawarehouse
    module Configuration
      module_function

      def configure
        yield(options) if block_given?
      end

      def options
        @configuration ||= OpenStruct.new
      end

    end
  end
end
