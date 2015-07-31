require 'quandl/client'
Quandl::Client.use 'https://www.quandl.com/api/'
Quandl::Client.token = ENV['QUANDL_AUTH_TOKEN']