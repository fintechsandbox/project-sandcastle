#WARNING:Not brilliant code
require 'quandl/client'

class QuandlEODAdapter

  def self.fetch_end_of_day_prices_for symbol, from_date="default", to_date="default"

    dataset = Quandl::Client::Dataset.find('EOD/' + symbol )

    from_date = dataset.from_date if from_date == "default"
    to_date = dataset.to_date if to_date == "default"

    closing_prices = {}

    if dataset && dataset.data
      dataset.data.trim_start(from_date).trim_end(to_date).each do |entry|
        closing_prices[entry[0]] = entry[11]
      end
    end

    return closing_prices
  end
end
