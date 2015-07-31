How their data is arranged:
Quandl has a different away of thinking about datasets that you might expect.
For example, for EOD (end of day) price data, a dataset is available for AAPL. Another, separate dataset is available for eod data for MSFT.

What this means is that if you want to get prices for many securities for each day, you need to pull for each of those databases. To get eod prices for 20 securites you mneed to pull from 20 databases.

Getting CUSIP data:
Quandl is not oriented around CUSIP identifiers. It is oriented around ticker symbols. There are no CUSIP identtifiers in any of the datasets (this is according to Qunandl hep desk in spring '15). Also, Quandl does not have a security master.