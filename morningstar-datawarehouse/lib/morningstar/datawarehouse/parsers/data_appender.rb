module Morningstar
  module Datawarehouse
    module Parsers
      class DataAppender
        def append(parsed_data)
          raise NotImplementedError
        end
      end
    end
  end
end
