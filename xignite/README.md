[Documentation|https://www.xignite.com/Products]

API Explorer for each of the APIs is linked to above


## How to get started

### Step 1
Explore their API to see what's available and run some sample queries.

Their docs are in better state than most (linked above): interactive API explorer can render sample data requests live. There are two caveats: some of them are broken and some require special permissions. Once you've signed a contract with them you can view the same API explorer with permissions. 

### Step 2
Contct sale person and sign the agreement (Kerry Langstaff)

### Step 3
Make sure you email your rep to ask for the entire asset universe list - their coverage isn't complete, so make sure they have the exact data points you need. 

## Caveats
Their Edgar API claims to be realtime, but is only showing previous days data. We went a few rounds with support, but there was no resolution. Can take a week or two to work out permissions to various APIS in your account

## API Clients
 Ruby support is pretty good with several gems available, we've forked our own because the ones we found were built for hitting one specific API endpoint. We're generalizing what's out there to allow it to be easily expanded for various of their API products (https://github.com/DataSimply/xignite/)

## Asset Universe

### Assets https://www.xignite.com/product/XigniteGlobalHistorical/

* All constituents of Russell 3000
* All US-traded ETFs 

### Fund Assets https://www.xignite.com/product/XigniteNAVs/
* All US-provider mutual funds

### Asset Attributes

* Ticker(s) 
* Company Name 
* Asset Name 
* ISIN 
* Cusip 
* Sedol(s) 
* Industry Exposure(s) 
* Sector Exposure(s) 
* Dividend Yield 

Do Not Have:
* Country Exposure(s)
* Currency Exposure(s)

###https://www.xignite.com/product/us-etf-mutual-fund-market-data/

**See listfundfundamentals_20150210231715.csv for full list of available fields here
* Expense Ratio 
* Asset Type 

Don't Have
* Duration ??
* Credit Quality ??
* Assets Under Management (yes-top 10)
* Number of Constituents

## Market Data https://www.xignite.com/product/XigniteGlobalHistorical/

* Closing Prices
* Closing Total Returns
* Volume

Don't Have
* Traded Range
* Market Capitalization (can be calculated) (available via fundamentals service. https://www.xignite.com/product/XigniteGlobalFundamentals/api/)
 

## Functions To Pull Data

### GetGlobalHistoricalQuotesRange

This operation returns a complete range of historical stock quotes for an equity based on a date range (start date, end date). This includes the adjusted price as specified. 

     var objectPassed = {
		"requestName":'hist_price',
		"optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS','AdjustmentMethod':'SplitOnly','StartDate':'6/12/2014','EndDate':"6/11/2015"}}
	};

    historical(objectPassed);

### GetCompanyFundamentalList

Returns values for multiple fundamental data types for one company.

     var objectPassed = {
		"requestName":'get_company_fundamental_list',
		"optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS', 'FundamentalTypes':'TotalAssets,HighPriceLTM,LowPriceLTM,AverageDailyVolumeLastTwelveMonths,PERatio,DividendRate,LastDividendYield,MarketCapitalization', 'UpdatedSince':''}}
    };

    fundamental(objectPassed);

### GetLogo

Get the logo for a company.

     var objectPassed = {
		"requestName":'logo',
		"optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS'}}
    };

    logos(objectPassed);

### GetGlobalDelayedQuote

Returns a delayed quote for a global security.

     var objectPassed = {
		"requestName":'get_technical',
		"optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS'}}
	};

    technical(objectPassed);

### GetAllEquityOptionChain

Returns the complete option chain for an equity. (Also called with technicals and fundamentals to get implied vol, dividends etc).

     var objectPassed = {
	        "requestName":'get_all_equity_option_chain',
	        "optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS', 'SymbologyType':'', 'OptionExchange':''}}
    };

    var objectPassedTech = {
		"requestName":'get_technical',
		"optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS'}}
	};

    var objectPassedFund = {
	       "requestName":'get_company_fundamental_list',
	       "optionsObject":{'options':{'IdentifierType':'Symbol','Identifier':'GS','FundamentalTypes':'DividendRate', 'UpdatedSince':''}}
    };

    optionsTech(objectPassed, objectPassedTech, objectPassedFund);

## Manipulating Data

All of the modify functions in the lib/modify.js file are called with data from the functions above. They are called only when all of the data needed is loaded.

[http://www.xignite.com/product/global-security-master-data/api/ListExchanges/]


## Some commonly used US exchange codes

#### NYSE = XASE
#### NASDAQ = XNAS
#### AMEX = AMXO (AMEX is also known as NYSE MKT, which is different from NYSE) 

Xignite uses the standard ISO 10383 Market Identifier Codes (MIC) for the various exchanges. They are all listed there and in the ISO 10383 document, but you have to find them. Posting here for convenience.


