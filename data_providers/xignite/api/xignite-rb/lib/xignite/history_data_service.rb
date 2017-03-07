module Xignite
  class HistoryDataService < Xignite::BaseService

    def base_url
      'http://www.xignite.com/xGlobalHistorical'
    end

    def get_global_historical_quotes_as_of(params, &block)
      default_params = {
          IdentifierType: IDENTIFIER_TYPE,
          Identifier: '',
          AdjustmentMethod: 'SplitOnly',
          EndDate: '',
          PeriodType: 'Day',
          Periods: '180',
      }
      send_request('GetGlobalHistoricalQuotesAsOf', default_params.merge(params)){|data| block.call(data)}
    end


    def get_global_historical_quotes_range(params, &block)
      default_params = {
          IdentifierType: IDENTIFIER_TYPE,
          Identifier: '',
          AdjustmentMethod: 'SplitOnly',
          StartDate: '',
          EndDate: ''
      }
      send_request('GetGlobalHistoricalQuotesRange', default_params.merge(params)){|data| block.call(data)}
    end
  end
end
