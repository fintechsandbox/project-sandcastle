class Issue(object):

	def __init__(self, issueId, issueType, issueOrder, description,
		         name, ticker, ric, displayRic, instrumentPI, quotePI,
		         exchangeCountry, exchangeCode, exchangeDescription,
		         globalListingType, sharesPerListing,
		         recentSplitDate, recentSplitAmount):

		self.issueId = issueId
		self.issueType = issueType
		self.description = description
		self.issueOrder = issueOrder
		
		self.name = name
		self.ticker = ticker
		self.ric = ric
		self.displayRic = displayRic
		self.instrumentPI = instrumentPI
		self.quotePI = quotePI

		self.exchangeCountry = exchangeCountry
		self.exchangeCode = exchangeCode
		self.exchangeDescription = exchangeDescription

		self.globalListingType = globalListingType
		self.sharesPerListing = sharesPerListing

		self.recentSplitDate = recentSplitDate
		self.recentSplitAmount = recentSplitAmount

"""
		"Issues":{
			"Issue":[
				{"ID":"2","Type":"C","Desc":"Common Stock","Order":"1",
					"IssueID":[
						{"Type":"Name","Value":"Ordinary Shares"},
						{"Type":"Ticker","Value":"ADK"},
						{"Type":"RIC","Value":"ADK"},
						{"Type":"DisplayRIC","Value":"ADK.A"},
						{"Type":"InstrumentPI","Value":"30030923"},
						{"Type":"QuotePI","Value":"30038014"}],
					"Exchange":{"Code":"AMEX","Country":"USA","Value":"NYSE MKT LLC"},
					"MostRecentSplit":{"Date":"2012-10-03","Value":"1.05"}
				}
			]
		},
"""