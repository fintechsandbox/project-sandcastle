# Data Simply POV:
We chose Quandle because it was fast and easy to get the basics going (EOD via Symbol). We're on a 
Ruby platform, and a gem is availabe form Quandl (and others):

* [https://github.com/quandl/quandl-ruby] <--- pre-release version for new v3 API

* [https://github.com/quandl/quandl_client] <---This is the old deprecated client

the quandl_client it out of date and requires you to manually specify the HTTPS version of the API
URL (see line 2 of quandl_initializer.rb)

The gem has the advantage that you never need to see the raw data from their API. Caveats: The old version isn't 
updated, and the new one is of unknown quality and recency. 

We're not updating to the new one because we don't want to have to make a separate call for each 
symbol we're intetested in (i.e. all of them). Also they don't have CUSIP data.
 
## How their data is arranged
Quandl has a different away of thinking about datasets that you might expect.
For example, for EOD (end of day) price data, a dataset is available for AAPL. Another, separate dataset is available for eod data for MSFT.

What this means is that if you want to get prices for many securities for each day, you need to pull for each of those databases. To get eod prices for 20 securites you mneed to pull from 20 databases.

## Getting CUSIP data:
Quandl is not oriented around CUSIP identifiers. It is oriented around ticker symbols. There are no CUSIP identtifiers in any of the datasets (this is according to Qunandl hep desk in spring '15). Also, Quandl does not have a security master.
