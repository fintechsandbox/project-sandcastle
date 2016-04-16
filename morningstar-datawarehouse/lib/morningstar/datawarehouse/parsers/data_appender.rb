module Morningstar
  module Datawarehouse
    module Parsers
      class DataAppender

        def deleted_shares(shares=[], universe=nil)
          raise NotImplementedError
        end

        def under_review_shares(shares=[], universe=nil)
          raise NotImplementedError
        end

        def updated_shares(shares=[], universe=nil)
          raise NotImplementedError
        end

        def append(parsed_data)
          raise NotImplementedError
        end
      end
    end
  end
end
