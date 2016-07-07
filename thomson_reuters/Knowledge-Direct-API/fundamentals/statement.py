class Statement(object):

	def __init__(self, endDate, fiscalType, fiscalYear, statementType,
				 periodType, periodDescription, periodLength,
		         auditorCode, auditorName, 
				 source, sourceDate, statementDate, updateType, updateTypeDescription,
				 auditorOpinion, auditorOpinionDescription, lineItems):

		self.endDate = endDate
		self.fiscalType = fiscalType
		self.fiscalYear = fiscalYear
		self.statementType = statementType
		self.periodType = periodType
		self.periodDescription = periodDescription
		self.periodLength = periodLength
		self.auditorName = auditorName
		self.auditorCode = auditorCode
		self.source = source
		self.sourceDate = sourceDate
		self.statementDate = statementDate
		self.updateType = updateType
		self.updateTypeDescription = updateTypeDescription
		self.auditorOpinion = auditorOpinion
		self.auditorOpinionDescription = auditorOpinionDescription
		self.lineItems = lineItems
