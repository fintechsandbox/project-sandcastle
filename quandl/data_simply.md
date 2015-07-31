## Data Simply POV:

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

 
