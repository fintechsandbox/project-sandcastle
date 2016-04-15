# Xignite

Xignite RESTful web services client implemented in Ruby.

## Configuration

Simply set API token provided by Xignite:
 
```ruby
Xignite::Configuration.configure do |config|
  config.apitoken = 'XIGNITE_API_TOKEN'
  config.request_timeout = 5 #Specifies request timeout. Default value is 10 sec. 
end
```

## Usage

### XigniteGlobalMaster

Global Security Master API List.

The following end-points are available:

**GetMasterByIdentifier**

```ruby
Xignite::MasterDataService.new.get_master_by_identifier(params) {|response| ... }
```

Default parameters values:

* _IdentifierType_ is _ISIN_

**GetMasterByIdentifiers**

```ruby
Xignite::MasterDataService.new.get_master_by_identifiers(params) {|response| ... }
```

Default parameters values:

* _IdentifierType_ is _ISIN_

**ListExchanges**

```ruby
Xignite::MasterDataService.new.list_exchanges(params) {|response| ... }
```

**ListMICToLegacyExchange**

```ruby
Xignite::MasterDataService.new.list_mic_to_legacy_exchange(params) {|response| ... }
```

**GetMasterByExchangeChanges**

```ruby
Xignite::MasterDataService.new.get_master_by_exchange_changes(params) {|response| ... }
```

Default parameters values:

* _InstrumentClass_ is _All_

**GetMasterByExchange**

```ruby
Xignite::MasterDataService.new.get_master_by_exchange(params) {|response| ... }
```

Default parameters values:

* _InstrumentClass_ is _All_

**GetInstruments**

```ruby
Xignite::MasterDataService.new.get_instruments(params) {|response| ... }
```

Default parameters values:

* _IncludeRelated_ is _Securities_
* _IdentifierType_ is _ISIN_

**GetInstrument**

```ruby
Xignite::MasterDataService.new.get_instrument(params) {|response| ... }
```

Default parameters values:

* _IncludeRelated_ is _None_
* _IdentifierType_ is _ISIN_

### XigniteGlobalHistorical

Global Historical Stocks API List

**GetGlobalHistoricalQuotesAsOf**

```ruby
Xignite::HistoryDataService.new.get_global_historical_quotes_as_of(params) {|response| ... }
```

Default parameters values:

* _AdjustmentMethod_ is _SplitOnly_
* _PeriodType_ is _Day_
* _Periods_ is _180_

**GetGlobalHistoricalQuotesRange**

```ruby
Xignite::HistoryDataService.new.get_global_historical_quotes_range(params) {|response| ... }
```

Default parameters values:

* _AdjustmentMethod_ is _SplitOnly_
* _IdentifierType_ is _ISIN_

### XigniteFundFundamentals

ETF, Mutual Fund, Money Market API List

**GetFundAssetAllocation**

```ruby
Xignite::FundAssetService.new.get_fund_asset_allocation(params) {|response| ... }
```

Default parameters values:

* _UpdatedSince_ is _Date.today_
* _IdentifierType_ is _ISIN_

**GetFundFundamentalList**

```ruby
Xignite::FundAssetService.new.get_fund_fundamental_list(params) {|response| ... }
```

Default parameters values:

* _UpdatedSince_ is _01/01/2014_
* _IdentifierType_ is _ISIN_

### XigniteGlobalCurrencies

Real-Time and Historical Foreign Currency Exchange Rates API (Forex/FX)

**GetLondonHistoricalRatesRange**

```ruby
Xignite::CurrencyService.new.get_historical_rates_ranges(params) {|response| ... }
```

Default parameters values:

* _Symbol_  is _USDEUR_
* _FixingTime_ is _22:00_
* _PeriodType_ is _Daily_
* _PriceType_ is _Mid_
* _StartDate_ is _Date.today_
* _EndDate_ is _Date.today_

