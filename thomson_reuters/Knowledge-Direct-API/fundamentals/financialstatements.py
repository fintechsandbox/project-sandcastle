import json
from companyidentifiers import CompanyIdentifiers
from financialstatement import FinancialStatement
from companygeneralinfo import CompanyGeneralInfo
from chartofaccount import ChartOfAccount

from statement import Statement
from issuers import Issue
from collections import namedtuple

__all__ = ['parse_json']


def convertJsonTuple(jsonData):
	data = json.loads(jsonData, object_hook=lambda d: namedtuple('data', d.keys())(*d.values()))
	return data

##  Company Identifiers Conversion
##  -- Convert the jsonData to a Company Identifier
##
def company_identifiers(data):

  coids = data.GetFinancialStatementsReports_Response_1.FundamentalReports.ReportFinancialStatements.CoIDs

  # RepNo
  rep_no = None
  # CompanyName
  company_name = None
  # MXID
  mx_id = None
  # CIKNo
  cik_number = None
  #IRSNo
  irs_number = None
  
  for coid in coids.CoID:
    co_id_type = coid.Type
    co_id_value = coid.Value
    if co_id_type == "RepNo":
      rep_no = co_id_value
    elif co_id_type == "CompanyName":
      company_name = co_id_value
    elif co_id_type == "MXID":
      mx_id = co_id_value
    elif co_id_type == "CIKNo":
      cik_number = co_id_value
    elif co_id_type == "IRSNo":
      irs_number = co_id_value
    else:
    	print "Unsupported COID: " + col_id_type

  company_ids = CompanyIdentifiers(rep_no, company_name, mx_id, cik_number, irs_number)

  return company_ids


##  Issue
##  -- Convert the jsonData to a Company Identifier
##
def issuers(data):

  issueData = data.GetFinancialStatementsReports_Response_1.FundamentalReports.ReportFinancialStatements.Issues

  issues = []

  for issue in issueData.Issue:

    name = None
    ticker = None
    ric = None
    displayRic = None
    instrumentPI = None
    quotePI = None

    for issueID in issue.IssueID:
    	issue_type = issueID.Type
    	issue_value = issueID.Value

    	if issue_type == 'Name':
    		name = issue_value
    	elif issue_type == 'Ticker':
    		ticker = issue_value
    	elif issue_type == 'RIC':
    		ric = issue_value
    	elif issue_type == 'DisplayRIC':
    		displayRic = issue_value
    	elif issue_type == 'InstrumentPI':
    		instrumentPI = issue_value
    	elif issue_type == 'QuotePI':
    		quotePI == issue_value
    	else:
    		print "Unsupported Issue TYPE: " + issue_type
    
    global_listing_type_value = None
    global_listing_type_shares = None

    if hasattr(issue,'GlobalListingType'):

    	global_listing_type_value = issue.GlobalListingType.Value

    	if hasattr(issue.GlobalListingType, 'SharesPerListing'):
    		global_listing_type_shares = issue.GlobalListingType.SharesPerListing


    recent_split_date = None
    recent_split_amount = None
    if hasattr(issue, 'MostRecentSplit'):
    	recent_split_amount = issue.MostRecentSplit.Value
    	recent_split_date =  issue.MostRecentSplit.Date

    issuer = Issue(issue.ID, issue.Type, issue.Order, issue.Desc,
		         name, ticker, ric, displayRic, instrumentPI, quotePI,
		         issue.Exchange.Country, 
		         issue.Exchange.Code, 
		         issue.Exchange.Value,
		         global_listing_type_value,
		         global_listing_type_shares, 
		         recent_split_date, recent_split_amount)

    issues.append(issuer)

  return issues

##  Company General Info
##  -- Convert the jsonData to a Company General Info
##
def company_general_info(data):

  # Get Company INfo
  coGenInfoData = data.GetFinancialStatementsReports_Response_1.FundamentalReports.ReportFinancialStatements.CoGeneralInfo
  
  # Determine if we have Employee information
  employees = None
  employeesLastUpdate = None

  if hasattr(coGenInfoData, 'Employees'):
  	employees = coGenInfoData.Employees.Value
  	employeesLastUpdate = coGenInfoData.Employees.LastUpdated

  # Determine if we have Shares Out
  sharesOutDate = None
  sharesOut = None
  sharesOutFloat = None

  if hasattr(coGenInfoData, 'SharesOut'):
  		sharesOutDate = coGenInfoData.SharesOut.Date
  		sharesOut = coGenInfoData.ShareOut.Value
  		sharesOutFloat = coGenInfoData.SharesOut.TotalFloat

  # Determine if we have the lastest interim price
  latestAvailableInterim = None
  if hasattr(coGenInfoData, 'LatestAvailableInterim'):
  	latestAvailableInterim = coGenInfoData.LatestAvailableInterim

  # Determine if we have the lastest annual price
  latestAvailableAnnual = None
  if hasattr(coGenInfoData, 'LatestAvailableAnnual'):
  	latestAvailableAnnual = coGenInfoData.LatestAvailableAnnual	

  # Create Company Info Object
  company_general_info = CompanyGeneralInfo(coGenInfoData.CoStatus.Code, coGenInfoData.CoType.Code, 
  		coGenInfoData.CoType.Value,
  		coGenInfoData.LastModified, 
  		latestAvailableInterim,
  		latestAvailableAnnual,
  		employees, employeesLastUpdate,
  		sharesOutDate, sharesOut, sharesOutFloat,
  		coGenInfoData.ReportingCurrency.Code, coGenInfoData.ReportingCurrency.Value,
  		coGenInfoData.MostRecentExchange.Date, coGenInfoData.MostRecentExchange.Value)

  # Return our new object 
  return company_general_info

##  Financial Statements
##  
##

def build_statements(fiscal_periods):

    statements = []

    for fiscal_period in fiscal_periods:

		for financial_statement in fiscal_period.Statement:

			header = financial_statement.FPHeader

			periodTypeCode = None
			periodTypeValue = None
			periodLength = None

			auditorNameCode = None
			auditorNameValue = None

			auditorOpinionCode = None
			auditorOpinionValue = None 

			sourceDate = None

			if hasattr(header, 'periodType'):
				periodTypeCode = header.periodType.Code
				periodTypeValue = header.periodType.Value
				periodLength = header.PeriodLength

			if hasattr(header, 'AuditorName'):
				auditorNameCode = header.AuditorName.Code
				auditorNameValue = header.AuditorName.Value

			if hasattr(header, 'AuditorOpinion'):
				auditorOpinionCode = header.AuditorOpinion.Code
				auditorOpinionValue = header.AuditorOpinion.Value				

			if hasattr(header.Source, 'Date'):
				sourceDate = header.Source.Date

			statement = Statement(fiscal_period.EndDate, fiscal_period.Type, 
								  fiscal_period.FiscalYear, 
								  financial_statement.Type,
								  periodTypeCode, periodTypeValue, periodLength, 
								  auditorNameCode, 
								  auditorNameValue,
								  header.Source.Value, sourceDate,
								  header.StatementDate,
								  header.UpdateType.Code, header.UpdateType.Value,
								  auditorOpinionCode, auditorOpinionValue, 
								  financial_statement.lineItem)
			statements.append(statement)

    return statements

def financial_statements(data):
  
  statements = []

  fin_stmt = data.GetFinancialStatementsReports_Response_1.FundamentalReports.ReportFinancialStatements.FinancialStatements

  if hasattr(fin_stmt.AnnualPeriods, 'FiscalPeriod'):
  	annual_fiscal_periods = fin_stmt.AnnualPeriods.FiscalPeriod
	a_statements = build_statements(annual_fiscal_periods)
	statements += a_statements

  if hasattr(fin_stmt.InterimPeriods, 'FiscalPeriod'):
  	interim_fiscal_periods = fin_stmt.InterimPeriods.FiscalPeriod
	i_statements = build_statements(interim_fiscal_periods)
	statements += i_statements	

  return statements

##  Chart of Accounts
##  
##
def chart_of_accounts(data):

	chart_of_accounts = []

	coas = data.GetFinancialStatementsReports_Response_1.FundamentalReports.ReportFinancialStatements.FinancialStatements.COAMap

	for coa in coas.mapItem:
		chart_of_account = ChartOfAccount(coa.coaItem, coa.Value, coa.statementType, coa.lineID, coa.precision)
		chart_of_accounts.append(chart_of_account)

	return chart_of_accounts


## Parse JSON Data
##
##
def parse_json(jsonData):

  data = convertJsonTuple(jsonData)
  financial_statement = FinancialStatement()

  company_ids = company_identifiers(data)
  financial_statement.add_company_identifier(company_ids)

  issues = issuers(data)
  financial_statement.add_issuers(issues)

  co_general_info = company_general_info(data)
  financial_statement.add_company_general_info(co_general_info)

  statements = financial_statements(data)
  financial_statement.add_statements(statements)

  coa_array = chart_of_accounts(data)
  financial_statement.add_chart_of_accounts(coa_array)

  return financial_statement

