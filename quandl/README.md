## Data Simply POV:

We chose Quandle because it was fast and easy to get the basics going (EOD via Symbol). We're on a 
Ruby platform, and a gem is availabe form Quandl (and others):

*[https://github.com/quandl/quandl-ruby] <--- pre-release version for new v3 API
*[https://github.com/quandl/quandl_client] <---This is the old deprecated client

the quandl_client it out of date and requires you to manually specify the HTTPS version of the API
URL (see line 2 of quandl_initializer.rb)

 
