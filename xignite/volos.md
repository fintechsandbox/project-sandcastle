Volos
=====

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

