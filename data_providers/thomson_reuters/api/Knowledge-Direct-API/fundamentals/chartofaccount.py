class ChartOfAccount(object):

	def __init__(self, coa_item, coa_description, statement_type, line_id, precision):
		self.coa_item = coa_item
		self.coa_description = coa_description
		self.statement_type = statement_type
		self.line_id = line_id
		self.precision = precision