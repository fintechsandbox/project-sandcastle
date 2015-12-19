import json
import requests

class XigniteEconomicCalendar:

    def __init__(self, base_url, access_token):
        """
        Creates an XigniteEconomicCalendar object to communicate & fetch JSON data for a list of economic events

        Args:
            base_url: Base URL for querying Xignite
            access_token: API Access Token
        Returns:
            None
        """
        self.access_token = access_token
        self.base_url = base_url

    def construct_url(self, country, date_range):
        """
        Formats and constructs a URL used for querying Xignite based on a date range and country. 

        Args:
            date_range: ["start_date","end_date"] # date format: "12/01/2015"
            country: "US"
        Returns:
            string: formatted URL 
        """
        if date_range:
            url_parameters = {"ReleasedOnStart": date_range[0], "ReleasedOnEnd": date_range[1], "CountryCode": country}
            formatted_parameters = '&'.join('{0}={1}'.format(key, ''.join([str(var) for var in val])) for key, val in sorted(url_parameters.items()))
            return "%s?_token=%s&%s" % (self.base_url, self.access_token, formatted_parameters)

    def query(self, country, date_range):
        """
        Queries Xignite for a list of economic events by country code and date range. 

        Uses the Requests library: http://docs.python-requests.org/en/latest/

        Args:
            date_range: ["start_date","end_date"] # date format: "12/01/2015"
            country: "US"
        Returns:
            JSON
        """
        if not date_range and country: return

        url = self.construct_url(country, date_range)
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
    date_range = ["12/01/2015", "12/15/2015"]
    country = "US"
    xignite = XigniteEconomicCalendar("https://www.xignite.com/xCalendar.json/GetEventsByCountryCode", "access_token")
    print xignite.query(country, date_range)
