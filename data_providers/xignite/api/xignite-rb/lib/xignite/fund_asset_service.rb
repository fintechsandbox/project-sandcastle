module Xignite
  class FundAssetService < Xignite::BaseService

    def base_url
      'http://fundfundamentals.xignite.com/xfundfundamentals'
    end

    def get_fund_asset_allocation(params, &block)
      default_params = {
          IdentifierType: IDENTIFIER_TYPE,
          Identifier: '',
          UpdatedSince: Date.today.strftime('%m/%d/%Y'),
      }
      send_request('GetFundAssetAllocation', default_params.merge(params)){|data| block.call(data)}
    end

    def get_fund_fundamental_list(params, &block)
      #TODO clarify update date
      default_params = {
          IdentifierType: IDENTIFIER_TYPE,
          Identifier: '',
          FundamentalType: '',
          UpdatedSince: '01/01/2014',
      }

      send_request('GetFundFundamentalList', default_params.merge(params)){|data| block.call(data)}
    end
  end
end
