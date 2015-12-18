Requirements:
pip install requests
pip install pytz


Usage:
from global_historical import GluobalHistorical
g = GlobalHistorical(access_token) # your access_token as a string
g.fetch_new_data(["GOOG"]) # ["GOOG","AAPL"]

from xignite_global_quotes import XigniteGlobalQuotes
x = XigniteGlobalQuotes(access_token) # your access_token
x.fetch_new_data(["GOOG”]) # ["GOOG","AAPL"]
