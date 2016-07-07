class CompanyGeneralInfo(object):

	def __init__(self, statusCode, coType, coTypeDescription, lastModified, 
					latestAvailableInterim, latestAvailableAnnual, 
					employeeCount, employeeLastUpdate, 
					sharesOutDate, sharesOut, totalFloat, 
					reportingCurrency, reportingCurrencyDescription, 
					mostRecentExchangeDate, mostRecentExchange):

		self.statusCode = statusCode
		self.coType = coType
		self.coTypeDescription = coTypeDescription
		self.lastModified = lastModified
		self.latestAvailableInterim = latestAvailableInterim
		self.latestAvailableAnnual = latestAvailableAnnual
		self.employeeCount = employeeCount
		self.employeeLastUpdate = employeeLastUpdate
		self.sharesOutDate = sharesOutDate
		self.sharesOut = sharesOut
		self.totalFloat = totalFloat
		self.reportingCurrency = reportingCurrency
		self.reportingCurrencyDescription = reportingCurrencyDescription
		self.mostRecentExchangeDate = mostRecentExchangeDate
		self.mostRecentExchange = mostRecentExchange