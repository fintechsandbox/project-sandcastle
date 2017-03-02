module Xignite
  class CurrencyService < Xignite::BaseService
    DEFAULT_SYMBOLS = 'USDEUR'

    def base_url
      'http://globalcurrencies.xignite.com/xGlobalCurrencies'
    end

    def get_historical_rates_ranges(params, &block)
      default_params = {
          Symbol: DEFAULT_SYMBOLS,
          FixingTime: '22:00',
          PeriodType: 'Daily',
          PriceType: 'Mid',
          StartDate: Date.today.prev_day.strftime('%m/%d/%Y'),
          EndDate: Date.today.strftime('%m/%d/%Y'),
      }

      send_request('GetLondonHistoricalRatesRange', default_params.merge(params)){|data| block.call(data)}
    end

    def list_currencies(&block)
      default_params = { FixingTime: '22:00'}

      send_request('ListCurrencies', :json, default_params){|data| block.call(data)}
    end
  end
end
