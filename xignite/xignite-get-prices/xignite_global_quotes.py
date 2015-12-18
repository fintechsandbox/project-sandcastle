import json
import time
from data_provider import DataProvider

class XigniteGlobalQuotes(DataProvider):
    """
    This class responsible for getting and parsing
    global delayed quotes of Xignite's API
    """

    def __init__(self, access_token, max_number_of_symbols=200):
        """
        You need an access token and how many symbols you want to request at any given point in time
        """
        DataProvider.__init__(self,"https://www.xignite.com/xGlobalQuotes.json/GetGlobalDelayedQuotes",access_token)
        self.max_number_of_symbols = max_number_of_symbols

    def parse_data(self, data):
        if not data:
            print "In parse_data for symbols %s data is: %s" % (self.symbols, data)
            return

        data_format = {}
        for row in data:
            if not row:
                print "In parse_data for symbols %s row is: %s" % (self.symbols, row)
                continue

            if "You are not authorized" in row["Message"]:
                print "IN parse_data row[Message] is %s" % row["Message"]
                continue

            now = self.get_time()
            str_now = str(now)
            (symbol,row_new_format) = self.convert_to_trigger_format(str_now,row) # convert function to the type you need 

            if row_new_format is None:
                print "IN parse_data can't convert row is %s" % row
                continue

            if symbol is not None and row_new_format is not None:
                trigger_data_format[symbol] = row_new_format

        return data_format

    def get_data(self,var):
        """

        """
        symbols = var
        # making the api call to get the new request
        symbols_lists = [symbols[i:i+self.max_number_of_symbols] for i  in range(0, len(symbols), self.max_number_of_symbols)]
        data = []
        for symbols in symbols_lists:
            symbols = [ symbol.replace("#","%23") for symbol in symbols ]
            params = {"Identifiers":symbols,"IdentifierType":["Symbol"]}
            new_data = self.get(params=params)
            if new_data is None :
                print "In get_data : for params %s new_data is %s" % (params,new_data)
                continue
            data += new_data
        return data


if __name__ == "__main__":
    pass
