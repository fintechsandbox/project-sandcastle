
import requests
import json


class XigniteGlobalQuotes:

    def __init__(self, base_url, access_token):
        """
        Creates an XigniteGlobalQuotes object to communicate fetch JSON data for a list of symbols

        Args:
            base_url: Base URL for querying Xignite
            access_token: API Access Token
        Returns:
            None
        """
        self.access_token = access_token
        self.base_url = base_url

    def construct_url(self, symbols):
        """
        Formats and constructs a URL used for querying Xignite based on a list of symbols. 

        Args:
            symbols: ["GOOG","AAPL"]
        Returns:
            string: formatted URL 
        """
        if symbols:
            url_parameters = {"Identifiers": symbols,"IdentifierType": ["Symbol"]}
            formatted_parameters = '&'.join('{0}={1}'.format(key, ','.join([str(var) for var in val])) for key, val in sorted(url_parameters.items()))

            return "%s?_token=%s&%s" % (self.base_url, self.access_token, formatted_parameters)

    def query(self, symbols):
        """
        Queries Xignite for a list of symbols. 

        Uses the Requests library: http://docs.python-requests.org/en/latest/

        Args:
            params: ["GOOG", "AAPL"]
        Returns:    
            JSON
        """
        if not symbols: return

        url = self.construct_url(symbols)
        try:
            new_request = requests.get(url)
            if new_request.status_code == 200:
                return json.loads(new_request.text)

            if new_request.status_code == 401:
                print "In get: new request status code is %s" % (new_request.status_code)
                return

            print "In get: new request status code is %s" % (new_request.status_code)

        except Exception as e:
            print "In get: Exception is %s" % (e)
        return None


if __name__ == "__main__":
    ## EXAMPLE USAGE
    symbols = ["GOOG", "AAPL"]
    xignite = XigniteGlobalQuotes("https://www.xignite.com/xGlobalQuotes.json/GetGlobalDelayedQuotes", "access_token")

    print xignite.query(symbols)
    
