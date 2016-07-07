## Financial Statements
class FinancialStatement(object):

  def ticker(self):
  	return self.issuers[0].ticker

  def add_company_identifier(self, company_identifier):
  	self.company_identifier = company_identifier

  def add_issuers(self, issuers):
  	self.issuers = issuers  	

  def add_company_general_info(self, company_general_info):
  	self.company_general_info = company_general_info

  def add_statements(self, statements):
  	self.statements = statements

  def add_chart_of_accounts(self, chart_of_accounts):
  	self.chart_of_accounts = chart_of_accounts


