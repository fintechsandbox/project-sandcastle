package capiq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	apiEndpoint = "https://sdk.gds.standardandpoors.com/gdssdk/rest/v2/clientservice.json"
)

type CapIQData struct {
	RawData map[string]interface{}
}

type PeriodType string

const (
	Annual    PeriodType = "IQ_FY"
	Quarterly            = "IQ_FQ"
	TTM                  = "IQ_LTM"
)

type CapIQMnemonic struct {
	Name           string
	Code           string
	AdditionalInfo string
}

func GetUniqueId(ticker string, code string, periodType PeriodType) string {
	return ticker + "/" + code + "/" + string(periodType)
}

var GetCapIQCodes []CapIQMnemonic = []CapIQMnemonic{
	CapIQMnemonic{"Accounts Payable, Total", "IQ_AP", "Balance Sheet"},
	CapIQMnemonic{"Accounts Receivable Long-Term", "IQ_AR_LT", "Balance Sheet"},
	CapIQMnemonic{"Accounts Receivable, Total", "IQ_AR", "Balance Sheet"},
	CapIQMnemonic{"Accounts Receivables - Unbilled", "IQ_AR_UNBILLED", "Balance Sheet"},
	CapIQMnemonic{"Accrued Expenses, Total", "IQ_AE", "Balance Sheet"},
	CapIQMnemonic{"Accumulated Amortization of Goodwill", "IQ_ACCUM_AMORT_GW", "Balance Sheet"},
	CapIQMnemonic{"Accumulated Amortization of Intangible Assets", "IQ_ACCUM_AMORT_INTAN_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Accumulated Depreciation", "IQ_AD", "Balance Sheet"},
	CapIQMnemonic{"Additional Paid In Capital", "IQ_APIC", "Balance Sheet"},
	CapIQMnemonic{"Assets Available for Sale", "IQ_ASSETS_AVAILABLE_SALE", "Balance Sheet"},
	CapIQMnemonic{"Assets Held to Maturity", "IQ_ASSETS_HELD_MATURITY", "Balance Sheet"},
	CapIQMnemonic{"Book Value / Share", "IQ_BV_SHARE", "Balance Sheet"},
	CapIQMnemonic{"Book Value Per Share (As Reported)", "IQ_BV_SHARE_REPORTED", "Balance Sheet"},
	CapIQMnemonic{"Buildings, Total", "IQ_BUILDINGS", "Balance Sheet"},
	CapIQMnemonic{"Capital Leases", "IQ_CAPITAL_LEASES", "Balance Sheet"},
	CapIQMnemonic{"Cash And Equivalents", "IQ_CASH_EQUIV", "Balance Sheet"},
	CapIQMnemonic{"Cash Per Share - (Ratio)", "IQ_CASH_SHARE", "Balance Sheet"},
	CapIQMnemonic{"Common Stock & APIC", "IQ_COMMON_APIC", "Balance Sheet"},
	CapIQMnemonic{"Common Stock, Total", "IQ_COMMON", "Balance Sheet"},
	CapIQMnemonic{"Comprehensive Income and Other", "IQ_OTHER_EQUITY", "Balance Sheet"},
	CapIQMnemonic{"Construction In Progress, Total", "IQ_CIP", "Balance Sheet"},
	CapIQMnemonic{"Contingent Liabilities – Guarantees", "IQ_CONTINGENT_LIABILITIES", "Balance Sheet"},
	CapIQMnemonic{"Cost of Borrowings (Calculated - Non Banks)", "IQ_COST_BORROWING", "Balance Sheet"},
	CapIQMnemonic{"Current Financing Obligations", "IQ_TOTAL_DEBT_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Current Income Taxes Payable", "IQ_INC_TAX_PAY_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Current Portion of Capital Lease Obligations", "IQ_CURRENT_PORT_LEASES", "Balance Sheet"},
	CapIQMnemonic{"Current Portion of Long Term Debt Derivative", "IQ_CURRENT_PORT_DEBT_DERIVATIVES", "Balance Sheet"},
	CapIQMnemonic{"Current Portion of Long-Term Debt", "IQ_CURRENT_PORT_DEBT", "Balance Sheet"},
	CapIQMnemonic{"Debt Equivalent of Unfunded Proj. Benefit Obligation", "IQ_DEBT_EQUIV_NET_PBO", "Balance Sheet"},
	CapIQMnemonic{"Debt Securities in Issue", "IQ_DEBT_SECURITIES_IN_ISSUE", "Balance Sheet"},
	CapIQMnemonic{"Deferred Charges Long-Term", "IQ_DEF_CHARGES_LT", "Balance Sheet"},
	CapIQMnemonic{"Deferred Tax Assets Current", "IQ_DEF_TAX_ASSETS_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Deferred Tax Assets Long-Term", "IQ_DEF_TAX_ASSETS_LT", "Balance Sheet"},
	CapIQMnemonic{"Deferred Tax Liability Current", "IQ_DEF_TAX_LIAB_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Deferred Tax Liability Non Current", "IQ_DEF_TAX_LIAB_LT", "Balance Sheet"},
	CapIQMnemonic{"Derivative Assets Current", "IQ_DERIVATIVE_ASSETS_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Derivative Liabilities – Current", "IQ_DERIVATIVE_LIAB_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Derivative Liabilities – Non Current", "IQ_DERIVATIVE_LIAB_NON_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Derivative Trading Asset Securities", "IQ_DERIVATIVE_TRADING_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Equity Method Investments, Total", "IQ_EQUITY_METHOD", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Cash and Cash Equivalents", "IQ_FIN_DIV_CASH_EQUIV", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Debt Current", "IQ_FIN_DIV_DEBT_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Debt Non Current", "IQ_FIN_DIV_DEBT_LT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Loans and Leases Current", "IQ_FIN_DIV_LOANS_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Loans and Leases Long-Term", "IQ_FIN_DIV_LOANS_LT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Long Term Debt, Total", "IQ_FIN_DIV_LT_DEBT_TOTAL", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Minority Interest", "IQ_FIN_DIV_MINORITY_INTEREST", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Other Current Assets, Total", "IQ_FIN_DIV_ASSETS_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Other Current Liabilities, Total", "IQ_FIN_DIV_LIAB_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Other Long-Term Assets, Total", "IQ_FIN_DIV_ASSETS_LT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Other Non Current Liabilities, Total", "IQ_FIN_DIV_LIAB_LT", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Redeemable Minority Interest", "IQ_FIN_DIV_MINORITY_INT_REDEEM", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Short Term Debt, Total", "IQ_FIN_DIV_ST_DEBT_TOTAL", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Short Term Investments", "IQ_FIN_DIV_ST_INVEST", "Balance Sheet"},
	CapIQMnemonic{"Finance Division Total Debt", "IQ_FIN_DIV_DEBT_TOTAL", "Balance Sheet"},
	CapIQMnemonic{"Full Time Employees", "IQ_FULL_TIME", "Balance Sheet"},
	CapIQMnemonic{"Goodwill", "IQ_GW", "Balance Sheet"},
	CapIQMnemonic{"Gross Goodwill", "IQ_GROSS_GW", "Balance Sheet"},
	CapIQMnemonic{"Gross Intangible Assets", "IQ_GROSS_INTAN_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Gross Property Plant And Equipment", "IQ_GPPE", "Balance Sheet"},
	CapIQMnemonic{"Inventories - Finished Goods, Total", "IQ_FINISHED_INV", "Balance Sheet"},
	CapIQMnemonic{"Inventories - Others", "IQ_OTHER_INV", "Balance Sheet"},
	CapIQMnemonic{"Inventories - Raw Materials, Total", "IQ_RAW_INV", "Balance Sheet"},
	CapIQMnemonic{"Inventories - Work In Process, Total", "IQ_WIP_INV", "Balance Sheet"},
	CapIQMnemonic{"Inventory", "IQ_INVENTORY", "Balance Sheet"},
	CapIQMnemonic{"Investment Liabilities to Clients", "IQ_INVESTMENT_LIABILITIES_CLIENTS", "Balance Sheet"},
	CapIQMnemonic{"Land - (BS)", "IQ_LAND", "Balance Sheet"},
	CapIQMnemonic{"Leasehold Improvements - Gross", "IQ_LEASEHOLD_IMPROVEMENT", "Balance Sheet"},
	CapIQMnemonic{"LIFO Reserve, Total", "IQ_LIFOR", "Balance Sheet"},
	CapIQMnemonic{"Loans Held For Sale", "IQ_LOANS_FOR_SALE", "Balance Sheet"},
	CapIQMnemonic{"Loans Receivable Long-Term", "IQ_LOANS_RECEIV_LT", "Balance Sheet"},
	CapIQMnemonic{"Long-Term Debt", "IQ_LT_DEBT", "Balance Sheet"},
	CapIQMnemonic{"Long-term Investments", "IQ_LT_INVEST", "Balance Sheet"},
	CapIQMnemonic{"Machinery, Total", "IQ_MACHINERY", "Balance Sheet"},
	CapIQMnemonic{"Minority Interest, Total", "IQ_MINORITY_INTEREST", "Balance Sheet"},
	CapIQMnemonic{"Minority Interest, Total (Incl. Fin. Div)", "IQ_MINORITY_INTEREST_TOTAL", "Balance Sheet"},
	CapIQMnemonic{"Natural Resources, at Cost", "IQ_NATURAL_RESOURCES_COST", "Balance Sheet"},
	CapIQMnemonic{"Net Debt", "IQ_NET_DEBT", "Balance Sheet"},
	CapIQMnemonic{"Net Property Plant And Equipment", "IQ_NPPE", "Balance Sheet"},
	CapIQMnemonic{"Net Working Capital", "IQ_NET_WORKING_CAP", "Balance Sheet"},
	CapIQMnemonic{"Non-Current Financing Obligations", "IQ_TOTAL_DEBT_NON_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Notes Receivable", "IQ_ST_NOTE_RECEIV", "Balance Sheet"},
	CapIQMnemonic{"Other Assets, Total", "IQ_OTHER_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Other Current Assets, Total", "IQ_OTHER_CA_SUPPL", "Balance Sheet"},
	CapIQMnemonic{"Other Current Liabilities, Total", "IQ_OTHER_CL", "Balance Sheet"},
	CapIQMnemonic{"Other Intangibles, Total", "IQ_OTHER_INTAN", "Balance Sheet"},
	CapIQMnemonic{"Other Liabilities, Total", "IQ_OTHER_LIAB", "Balance Sheet"},
	CapIQMnemonic{"Other Long-Term Assets, Total", "IQ_OTHER_LT_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Other Non Current Liabilities", "IQ_OTHER_LIAB_LT", "Balance Sheet"},
	CapIQMnemonic{"Other Receivables", "IQ_OTHER_RECEIV", "Balance Sheet"},
	CapIQMnemonic{"Pension & Other Post Retirement Benefits", "IQ_PENSION", "Balance Sheet"},
	CapIQMnemonic{"Preferred Stock - Others", "IQ_PREF_OTHER", "Balance Sheet"},
	CapIQMnemonic{"Preferred Stock Convertible", "IQ_PREF_CONVERT", "Balance Sheet"},
	CapIQMnemonic{"Preferred Stock Non Redeemable", "IQ_PREF_NON_REDEEM", "Balance Sheet"},
	CapIQMnemonic{"Preferred Stock Redeemable", "IQ_PREF_REDEEM", "Balance Sheet"},
	CapIQMnemonic{"Total Debt", "IQ_TOTAL_DEBT", "Balance Sheet"},
	CapIQMnemonic{"Prepaid Expenses", "IQ_PREPAID_EXP", "Balance Sheet"},
	CapIQMnemonic{"Restricted Cash", "IQ_RESTRICTED_CASH", "Balance Sheet"},
	CapIQMnemonic{"Restricted Cash (Non-Current)", "IQ_RESTRICTED_CASH_NON_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Retained Earnings", "IQ_RE", "Balance Sheet"},
	CapIQMnemonic{"Short Term Investments", "IQ_ST_INVEST", "Balance Sheet"},
	CapIQMnemonic{"Short-Term Borrowing Derivatives", "IQ_ST_DEBT_DERIVATIVES", "Balance Sheet"},
	CapIQMnemonic{"Short-term Borrowings", "IQ_ST_DEBT", "Balance Sheet"},
	CapIQMnemonic{"Tangible Book Value", "IQ_TBV", "Balance Sheet"},
	CapIQMnemonic{"Tangible Book Value Per Share", "IQ_TBV_SHARE", "Balance Sheet"},
	CapIQMnemonic{"Tangible Book Value Per Share (Supple)", "IQ_TBV_SHARE_REPORTED", "Balance Sheet"},
	CapIQMnemonic{"Total Assets", "IQ_TOTAL_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Total Capital", "IQ_TOTAL_CAP", "Balance Sheet"},
	CapIQMnemonic{"Total Cash And Short Term Investments", "IQ_CASH_ST_INVEST", "Balance Sheet"},
	CapIQMnemonic{"Total Common Equity", "IQ_TOTAL_COMMON_EQUITY", "Balance Sheet"},
	CapIQMnemonic{"Total Common Shares Outstanding (ECS)", "IQ_TOTAL_OUTSTANDING_BS_DATE", "Balance Sheet"},
	CapIQMnemonic{"Total Current Assets", "IQ_TOTAL_CA", "Balance Sheet"},
	CapIQMnemonic{"Total Current Liabilities", "IQ_TOTAL_CL", "Balance Sheet"},
	CapIQMnemonic{"Total Debt (excl. Fin. Div. Debt)", "IQ_TOTAL_DEBT_EXCL_FIN", "Balance Sheet"},
	CapIQMnemonic{"Total Employees", "IQ_TOTAL_EMPLOYEES", "Balance Sheet"},
	CapIQMnemonic{"Total Equity", "IQ_TOTAL_EQUITY", "Balance Sheet"},
	CapIQMnemonic{"Total Intangibles", "IQ_GW_INTAN", "Balance Sheet"},
	CapIQMnemonic{"Total Liabilities - (Standard / Utility Template)", "IQ_TOTAL_LIAB", "Balance Sheet"},
	CapIQMnemonic{"Total Liabilities And Equity", "IQ_TOTAL_LIAB_EQUITY", "Balance Sheet"},
	CapIQMnemonic{"Total Receivables", "IQ_TOTAL_RECEIV", "Balance Sheet"},
	CapIQMnemonic{"Total Redeemable Minority Interest", "IQ_MINORITY_INT_REDEEM_TOT", "Balance Sheet"},
	CapIQMnemonic{"Total Restricted Cash", "IQ_RESTRICTED_CASH_TOTAL", "Balance Sheet"},
	CapIQMnemonic{"Total Shares Outstanding on Filing Date", "IQ_TOTAL_OUTSTANDING_FILING_DATE", "Balance Sheet"},
	CapIQMnemonic{"Trading Asset Securities", "IQ_TRADING_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Trading Portfolio Assets", "IQ_TRADING_PORTFOLIO_ASSETS", "Balance Sheet"},
	CapIQMnemonic{"Trading Portfolio Liabilities", "IQ_TRADING_PORTFOLIO_LIABILITIES", "Balance Sheet"},
	CapIQMnemonic{"Treasury Stock", "IQ_TREASURY", "Balance Sheet"},
	CapIQMnemonic{"Treasury Stock & Other", "IQ_TREASURY_OTHER_EQUITY", "Balance Sheet"},
	CapIQMnemonic{"Unearned Revenue Current, Total", "IQ_UNEARN_REV_CURRENT", "Balance Sheet"},
	CapIQMnemonic{"Unearned Revenue Non Current", "IQ_UNEARN_REV_LT", "Balance Sheet"},
	CapIQMnemonic{"Working Capital", "IQ_WORKING_CAP", "Balance Sheet"},
	CapIQMnemonic{"Total Preferred Equity", "IQ_PREF_EQUITY", "Balance Sheet"},
	CapIQMnemonic{"Net Income Margin %", "IQ_NI_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"(Gain) Loss From Sale Of Asset", "IQ_GAIN_ASSETS_CF", "Cash Flow"},
	CapIQMnemonic{"(Gain) Loss on Sale of Investments - (CF)", "IQ_GAIN_INVEST_CF", "Cash Flow"},
	CapIQMnemonic{"(Income) Loss On Equity Investments - (CF)", "IQ_INC_EQUITY_CF", "Cash Flow"},
	CapIQMnemonic{"Amortization of Deferred Charges, Total - (CF)", "IQ_OTHER_AMORT", "Cash Flow"},
	CapIQMnemonic{"Amortization of Goodwill and Intangible Assets - (CF) - (Template Specific)", "IQ_GW_INTAN_AMORT_CF", "Cash Flow"},
	CapIQMnemonic{"Amortization of Negative Goodwill - (CF)", "IQ_NEGATIVE_GW_AMORT_CF", "Cash Flow"},
	CapIQMnemonic{"Asset Writedown & Restructuring Costs", "IQ_ASSET_WRITEDOWN_CF", "Cash Flow"},
	CapIQMnemonic{"Capital Expenditure", "IQ_CAPEX", "Cash Flow"},
	CapIQMnemonic{"Cash Acquisitions", "IQ_CASH_ACQUIRE_CF", "Cash Flow"},
	CapIQMnemonic{"Cash from Financing", "IQ_CASH_FINAN", "Cash Flow"},
	CapIQMnemonic{"Cash from Investing", "IQ_CASH_INVEST", "Cash Flow"},
	CapIQMnemonic{"Cash from Operations", "IQ_CASH_OPER", "Cash Flow"},
	CapIQMnemonic{"Cash Income Tax Paid - Financing Activities", "IQ_CASH_TAXES_FINAN", "Cash Flow"},
	CapIQMnemonic{"Cash Income Tax Paid - Investing Activities", "IQ_CASH_TAXES_INVEST", "Cash Flow"},
	CapIQMnemonic{"Cash Income Tax Paid - Operating Activities", "IQ_CASH_TAXES_OPER", "Cash Flow"},
	CapIQMnemonic{"Cash Income Tax Paid (Refund)", "IQ_CASH_TAXES", "Cash Flow"},
	CapIQMnemonic{"Cash Interest Paid", "IQ_CASH_INTEREST", "Cash Flow"},
	CapIQMnemonic{"Cash Interest Paid - Financing Activities", "IQ_CASH_INTEREST_FINAN", "Cash Flow"},
	CapIQMnemonic{"Cash Interest Paid - Investing Activities", "IQ_CASH_INTEREST_INVEST", "Cash Flow"},
	CapIQMnemonic{"Cash Interest Paid - Operating Activities", "IQ_CASH_INTEREST_OPER", "Cash Flow"},
	CapIQMnemonic{"Change In Accounts Payable", "IQ_CHANGE_AP", "Cash Flow"},
	CapIQMnemonic{"Change In Accounts Receivable", "IQ_CHANGE_AR", "Cash Flow"},
	CapIQMnemonic{"Change In Deferred Taxes", "IQ_CHANGE_DEF_TAX", "Cash Flow"},
	CapIQMnemonic{"Change In Income Taxes", "IQ_CHANGE_INC_TAX", "Cash Flow"},
	CapIQMnemonic{"Change In Inventories", "IQ_CHANGE_INVENTORY", "Cash Flow"},
	CapIQMnemonic{"Change in Net Operating Assets", "IQ_CHANGE_NET_OPER_ASSETS", "Cash Flow"},
	CapIQMnemonic{"Change In Net Working Capital", "IQ_CHANGE_NET_WORKING_CAPITAL", "Cash Flow"},
	CapIQMnemonic{"Change in Other Net Operating Assets", "IQ_CHANGE_OTHER_NET_OPER_ASSETS", "Cash Flow"},
	CapIQMnemonic{"Change in Trading Asset Securities", "IQ_CHANGE_TRADING_ASSETS", "Cash Flow"},
	CapIQMnemonic{"Change in Unearned Revenues", "IQ_CHANGE_UNEARN_REV", "Cash Flow"},
	CapIQMnemonic{"Common & Preferred Stock Dividends Paid", "IQ_TOTAL_DIV_PAID_CF", "Cash Flow"},
	CapIQMnemonic{"Common Dividends Paid", "IQ_COMMON_DIV_CF", "Cash Flow"},
	CapIQMnemonic{"Depreciation & Amortization - CF", "IQ_DA_SUPPL_CF", "Cash Flow"},
	CapIQMnemonic{"Depreciation & Amortization, Total - CF", "IQ_DA_CF", "Cash Flow"},
	CapIQMnemonic{"Depreciation of Rental Assets", "IQ_DEPRECIATION_RENTAL_ASSETS_CF", "Cash Flow"},
	CapIQMnemonic{"Divestitures", "IQ_DIVEST_CF", "Cash Flow"},
	CapIQMnemonic{"Foreign Exchange Rate Adjustments", "IQ_FX", "Cash Flow"},
	CapIQMnemonic{"Free Cash Flow / Share", "IQ_CF_SHARE", "Cash Flow"},
	CapIQMnemonic{"Impairment of Oil, Gas & Mineral Properties - (CF)", "IQ_OIL_IMPAIR", "Cash Flow"},
	CapIQMnemonic{"Investment in Marketable and Equity Securities, Total", "IQ_INVEST_SECURITY_CF", "Cash Flow"},
	CapIQMnemonic{"Issuance of Common Stock", "IQ_COMMON_ISSUED", "Cash Flow"},
	CapIQMnemonic{"Issuance of Preferred Stock", "IQ_PREF_ISSUED", "Cash Flow"},
	CapIQMnemonic{"Levered Free Cash Flow", "IQ_LEVERED_FCF", "Cash Flow"},
	CapIQMnemonic{"Long-Term Debt Issued, Total", "IQ_LT_DEBT_ISSUED", "Cash Flow"},
	CapIQMnemonic{"Long-Term Debt Repaid, Total", "IQ_LT_DEBT_REPAID", "Cash Flow"},
	CapIQMnemonic{"Minority Interest in Earnings - (CF)", "IQ_MINORITY_INTEREST_CF", "Cash Flow"},
	CapIQMnemonic{"Net (Increase) Decrease in Loans Originated / Sold - Investing", "IQ_INVEST_LOANS_CF", "Cash Flow"},
	CapIQMnemonic{"Net (Increase) Decrease in Loans Originated / Sold - Operating", "IQ_LOANS_CF", "Cash Flow"},
	CapIQMnemonic{"Net Cash From Discontinued Operations", "IQ_DO_CF", "Cash Flow"},
	CapIQMnemonic{"Net Cash From Discontinued Operations - (Financing Activities)", "IQ_DO_FINANCE_CF", "Cash Flow"},
	CapIQMnemonic{"Net Cash From Discontinued Operations - (Investing Activities)", "IQ_DO_INVEST_CF", "Cash Flow"},
	CapIQMnemonic{"Net Change in Cash", "IQ_NET_CHANGE", "Cash Flow"},
	CapIQMnemonic{"Net Debt Issued / Repaid", "IQ_NET_DEBT_ISSUED", "Cash Flow"},
	CapIQMnemonic{"Other Financing Activities, Total", "IQ_OTHER_FINANCE_ACT_SUPPL", "Cash Flow"},
	CapIQMnemonic{"Other Investing Activities, Total", "IQ_OTHER_INVEST_ACT_SUPPL", "Cash Flow"},
	CapIQMnemonic{"Other Non-Cash Items, Total", "IQ_NON_CASH_ITEMS", "Cash Flow"},
	CapIQMnemonic{"Other Operating Activities, Total", "IQ_OTHER_OPER_ACT", "Cash Flow"},
	CapIQMnemonic{"Preferred Dividends Paid", "IQ_PREF_DIV_CF", "Cash Flow"},
	CapIQMnemonic{"Provision and Write-off of Bad Debts", "IQ_PROV_BAD_DEBTS_CF", "Cash Flow"},
	CapIQMnemonic{"Provision for Credit Losses", "IQ_CREDIT_LOSS_CF", "Cash Flow"},
	CapIQMnemonic{"Repurchase of Common Stock", "IQ_COMMON_REP", "Cash Flow"},
	CapIQMnemonic{"Repurchase of Preferred Stock", "IQ_PREF_REP", "Cash Flow"},
	CapIQMnemonic{"Sale (Purchase) of Intangible assets", "IQ_SALE_INTAN_CF", "Cash Flow"},
	CapIQMnemonic{"Sale (Purchase) of Real Estate Properties", "IQ_SALE_REAL_ESTATE_CF", "Cash Flow"},
	CapIQMnemonic{"Sale of Property, Plant, and Equipment", "IQ_SALE_PPE_CF", "Cash Flow"},
	CapIQMnemonic{"Sale Proceeds from Rental Assets", "IQ_SALE_PROCEEDS_RENTAL_ASSETS", "Cash Flow"},
	CapIQMnemonic{"Short Term Debt Issued", "IQ_ST_DEBT_ISSUED", "Cash Flow"},
	CapIQMnemonic{"Short Term Debt Repaid", "IQ_ST_DEBT_REPAID", "Cash Flow"},
	CapIQMnemonic{"Special Dividend Paid", "IQ_SPECIAL_DIV_CF", "Cash Flow"},
	CapIQMnemonic{"Stock-Based Compensation (CF)", "IQ_STOCK_BASED_CF", "Cash Flow"},
	CapIQMnemonic{"Tax Benefit from Stock Options", "IQ_TAX_BENEFIT_OPTIONS", "Cash Flow"},
	CapIQMnemonic{"Total Debt Issued", "IQ_TOTAL_DEBT_ISSUED", "Cash Flow"},
	CapIQMnemonic{"Total Debt Repaid", "IQ_TOTAL_DEBT_REPAID", "Cash Flow"},
	CapIQMnemonic{"Unlevered Free Cash Flow", "IQ_UNLEVERED_FCF", "Cash Flow"},
	CapIQMnemonic{"(EBITDA - Capex) / Interest Expense", "IQ_EBITDA_CAPEX_INT", "Financial Ratios"},
	CapIQMnemonic{"Altman Z Score Using the Average Stock Information for a Period", "IQ_Z_SCORE", "Financial Ratios"},
	CapIQMnemonic{"Asset Turnover", "IQ_ASSET_TURNS", "Financial Ratios"},
	CapIQMnemonic{"Average Days Payable Outstanding", "IQ_DAYS_PAYABLE_OUT", "Financial Ratios"},
	CapIQMnemonic{"Capex as % of Revenues", "IQ_CAPEX_PCT_REV", "Financial Ratios"},
	CapIQMnemonic{"Cash Conversion Cycle (Average Days)", "IQ_CASH_CONVERSION", "Financial Ratios"},
	CapIQMnemonic{"Current Ratio", "IQ_CURRENT_RATIO", "Financial Ratios"},
	CapIQMnemonic{"Days Outstanding Inventory (Average Inventory)", "IQ_DAYS_INVENTORY_OUT", "Financial Ratios"},
	CapIQMnemonic{"Days Sales Outstanding (Average Receivables)", "IQ_DAYS_SALES_OUT", "Financial Ratios"},
	CapIQMnemonic{"EBIT / Interest Expense", "IQ_EBIT_INT", "Financial Ratios"},
	CapIQMnemonic{"EBIT Margin %", "IQ_EBIT_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"EBITA Margin", "IQ_EBITA_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"EBITDA / Interest Expense", "IQ_EBITDA_INT", "Financial Ratios"},
	CapIQMnemonic{"EBITDA Margin %", "IQ_EBITDA_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Fixed Assets Turnover (Average Fixed Assets)", "IQ_FIXED_ASSET_TURNS", "Financial Ratios"},
	CapIQMnemonic{"Income From Continuing Operations Margin %", "IQ_EARNING_CO_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Inventory Turnover (Average Inventory)", "IQ_INVENTORY_TURNS", "Financial Ratios"},
	CapIQMnemonic{"Levered Free Cash Flow Margin", "IQ_LFCF_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Long-Term Debt / Total Capital", "IQ_LT_DEBT_CAPITAL", "Financial Ratios"},
	CapIQMnemonic{"LT Debt/Equity", "IQ_LT_DEBT_EQUITY", "Financial Ratios"},
	CapIQMnemonic{"Net Avail. For Common Margin %", "IQ_NI_AVAIL_EXCL_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Net Debt / (EBITDA - Capex)", "IQ_NET_DEBT_EBITDA_CAPEX", "Financial Ratios"},
	CapIQMnemonic{"Normalized Net Income Margin", "IQ_NI_NORM_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Operating Cash Flow to Current Liabilities", "IQ_CFO_CURRENT_LIAB", "Financial Ratios"},
	CapIQMnemonic{"Quick Ratio", "IQ_QUICK_RATIO", "Financial Ratios"},
	CapIQMnemonic{"Receivables Turnover (Average Receivables)", "IQ_AR_TURNS", "Financial Ratios"},
	CapIQMnemonic{"Return on Assets", "IQ_RETURN_ASSETS", "Financial Ratios"},
	CapIQMnemonic{"Return on Common Equity", "IQ_RETURN_COMMON_EQUITY", "Financial Ratios"},
	CapIQMnemonic{"Return On Equity %", "IQ_RETURN_EQUITY", "Financial Ratios"},
	CapIQMnemonic{"Return on Total Capital", "IQ_RETURN_CAPITAL", "Financial Ratios"},
	CapIQMnemonic{"Revenue Per Employee (In Thousands)", "IQ_TOTAL_REV_EMPLOYEE", "Financial Ratios"},
	CapIQMnemonic{"SG&A Margin", "IQ_SGA_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Total Debt / EBITDA", "IQ_TOTAL_DEBT_EBITDA", "Financial Ratios"},
	CapIQMnemonic{"Total Debt / Total Capital", "IQ_TOTAL_DEBT_CAPITAL", "Financial Ratios"},
	CapIQMnemonic{"Total Debt/Equity", "IQ_TOTAL_DEBT_EQUITY", "Financial Ratios"},
	CapIQMnemonic{"Total Liabilities / Total Assets", "IQ_TOTAL_LIAB_TOTAL_ASSETS", "Financial Ratios"},
	CapIQMnemonic{"Unlevered Free Cash Flow Margin", "IQ_UFCF_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Advertising Expense", "IQ_ADVERTISING", "Income Statement"},
	CapIQMnemonic{"Aircraft Rent", "IQ_AIRCRAFT_RENT", "Income Statement"},
	CapIQMnemonic{"Amortization of Goodwill and Intangible Assets - (IS)", "IQ_GW_INTAN_AMORT", "Income Statement"},
	CapIQMnemonic{"Amortization of Negative Goodwill", "IQ_NEGATIVE_GW_AMORT", "Income Statement"},
	CapIQMnemonic{"Amortization of Negative Goodwill - (After Tax)", "IQ_NEGATIVE_GW_AMORT_AT", "Income Statement"},
	CapIQMnemonic{"Asset Writedown", "IQ_ASSET_WRITEDOWN", "Income Statement"},
	CapIQMnemonic{"Basic EPS - Continuing Operations", "IQ_BASIC_EPS_EXCL", "Income Statement"},
	CapIQMnemonic{"Basic Weighted Average Shares Outstanding", "IQ_BASIC_WEIGHT", "Income Statement"},
	CapIQMnemonic{"Common Stock Dividend", "IQ_COMMON_DIV", "Income Statement"},
	CapIQMnemonic{"Cost of Goods Sold, Total", "IQ_COGS", "Income Statement"},
	CapIQMnemonic{"Cost Of Revenues", "IQ_COST_REV", "Income Statement"},
	CapIQMnemonic{"Currency Exchange Gains (Loss)", "IQ_CURRENCY_GAIN", "Income Statement"},
	CapIQMnemonic{"Debt Issuance Costs", "IQ_DIC", "Income Statement"},
	CapIQMnemonic{"Diluted Net Income", "IQ_DILUT_NI", "Income Statement"},
	CapIQMnemonic{"Diluted Weighted Average Shares Outstanding", "IQ_DILUT_WEIGHT", "Income Statement"},
	CapIQMnemonic{"Dividend Per Share", "IQ_DIV_SHARE", "Income Statement"},
	CapIQMnemonic{"Earnings From Continuing Operations", "IQ_EARNING_CO", "Income Statement"},
	CapIQMnemonic{"Earnings Of Discontinued Operations", "IQ_DO", "Income Statement"},
	CapIQMnemonic{"Earnings to Parent before Extraordinary Items", "IQ_EARNING_PARENT_EXCL_EXTRA", "Income Statement"},
	CapIQMnemonic{"EBIT (Excl. SBC)", "IQ_EBIT_EXCL_SBC", "Income Statement"},
	CapIQMnemonic{"EBITA", "IQ_EBITA", "Income Statement"},
	CapIQMnemonic{"EBITA (Excl. SBC)", "IQ_EBITA_EXCL_SBC", "Income Statement"},
	CapIQMnemonic{"EBITDA", "IQ_EBITDA", "Income Statement"},
	CapIQMnemonic{"EBITDA (Excl. SBC)", "IQ_EBITDA_EXCL_SBC", "Income Statement"},
	CapIQMnemonic{"EBITDAR", "IQ_EBITDAR", "Income Statement"},
	CapIQMnemonic{"EBITDAR (Excl. SBC)", "IQ_EBITDAR_EXCL_SBC", "Income Statement"},
	CapIQMnemonic{"EBT, Excl. Unusual Items", "IQ_EBT_EXCL", "Income Statement"},
	CapIQMnemonic{"EBT, Incl. Unusual Items", "IQ_EBT", "Income Statement"},
	CapIQMnemonic{"Effective Tax Rate - (Ratio)", "IQ_EFFECT_TAX_RATE", "Income Statement"},
	CapIQMnemonic{"Excise Taxes Excluded from Sales, Total", "IQ_EXCISE_TAXES_EXCL_SALES", "Income Statement"},
	CapIQMnemonic{"Excise Taxes Included in Sales, Total", "IQ_EXCISE_TAXES_INCL_SALES", "Income Statement"},
	CapIQMnemonic{"Exploration/Drilling Expenses", "IQ_EXPLORE_DRILL_EXP_TOTAL", "Income Statement"},
	CapIQMnemonic{"Fees and Other Income", "IQ_FEES_OTHER_INCOME", "Income Statement"},
	CapIQMnemonic{"Finance Div. Operating Exp.", "IQ_FIN_DIV_EXP", "Income Statement"},
	CapIQMnemonic{"Finance Div. Revenues", "IQ_FIN_DIV_REV", "Income Statement"},
	CapIQMnemonic{"Gain (Loss) On Sale Of Assets", "IQ_GAIN_ASSETS", "Income Statement"},
	CapIQMnemonic{"Gain (Loss) On Sale Of Investments", "IQ_GAIN_INVEST", "Income Statement"},
	CapIQMnemonic{"General and Administrative Expenses", "IQ_GA_EXP", "Income Statement"},
	CapIQMnemonic{"Impairment of Goodwill", "IQ_IMPAIRMENT_GW", "Income Statement"},
	CapIQMnemonic{"Impairment of Oil, Gas & Mineral Properties - (IS)", "IQ_IMPAIR_OIL", "Income Statement"},
	CapIQMnemonic{"In Process R&D Expenses", "IQ_IPRD", "Income Statement"},
	CapIQMnemonic{"Income (Loss) On Equity Invest.", "IQ_INC_EQUITY", "Income Statement"},
	CapIQMnemonic{"Income Tax Expense", "IQ_INC_TAX", "Income Statement"},
	CapIQMnemonic{"Insurance Division Operating Expenses, Total", "IQ_INS_DIV_EXP", "Income Statement"},
	CapIQMnemonic{"Insurance Division Revenues", "IQ_INS_DIV_REV", "Income Statement"},
	CapIQMnemonic{"Insurance Settlements", "IQ_INS_SETTLE", "Income Statement"},
	CapIQMnemonic{"Interest And Invest. Income (Rev)", "IQ_INT_INV_INC", "Income Statement"},
	CapIQMnemonic{"Interest And Investment Income", "IQ_INTEREST_INVEST_INC", "Income Statement"},
	CapIQMnemonic{"Interest Capitalized", "IQ_CAPITALIZED_INTEREST", "Income Statement"},
	CapIQMnemonic{"Interest Expense - Finance Division", "IQ_FIN_DIV_INT_EXP", "Income Statement"},
	CapIQMnemonic{"Interest Expense (incl. Cap. Interest)", "IQ_INT_EXP_INCL_CAP", "Income Statement"},
	CapIQMnemonic{"Interest Expense, Total", "IQ_INTEREST_EXP", "Income Statement"},
	CapIQMnemonic{"Labor and Related Expenses", "IQ_SALARIES_OTHER_BENEFITS", "Income Statement"},
	CapIQMnemonic{"Legal Settlements", "IQ_LEGAL_SETTLE", "Income Statement"},
	CapIQMnemonic{"Marketing Expenses", "IQ_MARKETING", "Income Statement"},
	CapIQMnemonic{"Merger & Related Restructuring Charges", "IQ_MERGER", "Income Statement"},
	CapIQMnemonic{"Merger & Restructuring Charges", "IQ_MERGER_RESTRUCTURE", "Income Statement"},
	CapIQMnemonic{"Minimum Rental Expenses, Total", "IQ_MINIMUM_RENTAL", "Income Statement"},
	CapIQMnemonic{"Minority Interest in Earnings - (IS)", "IQ_MINORITY_INTEREST_IS", "Income Statement"},
	CapIQMnemonic{"Net EPS - Basic", "IQ_BASIC_EPS_INCL", "Income Statement"},
	CapIQMnemonic{"Net EPS - Diluted", "IQ_DILUT_EPS_INCL", "Income Statement"},
	CapIQMnemonic{"Net Income - (IS)", "IQ_NI", "Income Statement"},
	CapIQMnemonic{"Net Income to Common Excl. Extra Items", "IQ_NI_AVAIL_EXCL", "Income Statement"},
	CapIQMnemonic{"Net Income to Common Incl Extra Items", "IQ_NI_AVAIL_INCL", "Income Statement"},
	CapIQMnemonic{"Net Income to Company", "IQ_NI_COMPANY", "Income Statement"},
	CapIQMnemonic{"Net Interest Expenses", "IQ_NET_INTEREST_EXP", "Income Statement"},
	CapIQMnemonic{"Net Rental Expense", "IQ_NET_RENTAL_EXP", "Income Statement"},
	CapIQMnemonic{"Net Rental Expense, Total", "IQ_NET_RENTAL_EXP_FN", "Income Statement"},
	CapIQMnemonic{"Non-Cash Pension Expense", "IQ_NONCASH_PENSION_EXP", "Income Statement"},
	CapIQMnemonic{"Normalized Basic EPS", "IQ_EPS_NORM", "Income Statement"},
	CapIQMnemonic{"Normalized Diluted EPS", "IQ_DILUT_EPS_NORM", "Income Statement"},
	CapIQMnemonic{"Normalized Net Income", "IQ_NI_NORM", "Income Statement"},
	CapIQMnemonic{"Operating Income", "IQ_OPER_INC", "Income Statement"},
	CapIQMnemonic{"Other Non Operating Expenses, Total", "IQ_OTHER_NON_OPER_EXP", "Income Statement"},
	CapIQMnemonic{"Other Non Operating Income (Expenses)", "IQ_OTHER_NON_OPER_EXP_SUPPL", "Income Statement"},
	CapIQMnemonic{"Other Operating Expenses, Total", "IQ_TOTAL_OTHER_OPER", "Income Statement"},
	CapIQMnemonic{"Other Operating Expenses/(Income)", "IQ_OTHER_OPER", "Income Statement"},
	CapIQMnemonic{"Other Rental Expense, Total", "IQ_OTHER_RENTAL", "Income Statement"},
	CapIQMnemonic{"Other Revenues, Total", "IQ_OTHER_REV_SUPPL", "Income Statement"},
	CapIQMnemonic{"Other Unusual Items", "IQ_OTHER_UNUSUAL_SUPPL", "Income Statement"},
	CapIQMnemonic{"Other Unusual Items, Total", "IQ_OTHER_UNUSUAL", "Income Statement"},
	CapIQMnemonic{"Payout Ratio", "IQ_PAYOUT_RATIO", "Income Statement"},
	CapIQMnemonic{"Preferred Dividend and Other Adjustments", "IQ_PREF_DIV_OTHER", "Income Statement"},
	CapIQMnemonic{"Profit", "IQ_GP", "Income Statement"},
	CapIQMnemonic{"Provision for Bad Debts", "IQ_PROV_BAD_DEBTS", "Income Statement"},
	CapIQMnemonic{"Provision for Doubtful Accounts - Patient Service Revenue", "IQ_PROV_BAD_DEBTS_PATIENT_SERVICE_REV", "Income Statement"},
	CapIQMnemonic{"R&D Expenses", "IQ_RD_EXP", "Income Statement"},
	CapIQMnemonic{"Restructuring Charges", "IQ_RESTRUCTURE", "Income Statement"},
	CapIQMnemonic{"Revenue Per Share", "IQ_TOTAL_REV_SHARE", "Income Statement"},
	CapIQMnemonic{"Revenues", "IQ_REV", "Income Statement"},
	CapIQMnemonic{"Same Store Sales Growth (%)", "IQ_SAME_STORE", "Income Statement"},
	CapIQMnemonic{"Selling and Marketing Expenses", "IQ_SALES_MARKETING", "Income Statement"},
	CapIQMnemonic{"Selling General & Admin Expenses, Total", "IQ_SGA", "Income Statement"},
	CapIQMnemonic{"Special Dividend Per Share", "IQ_SPECIAL_DIV_SHARE", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., COGS (Total)", "IQ_STOCK_BASED_COGS", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., Exploration/Drilling Exp.", "IQ_STOCK_BASED_EXPLORE_DRILL", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., G&A Exp. (Total)", "IQ_STOCK_BASED_GA", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., Net of Tax (Total)", "IQ_STOCK_BASED_AT", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., Other (Total)", "IQ_STOCK_BASED_OTHER", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., R&D Exp. (Total)", "IQ_STOCK_BASED_RD", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., S&M Exp. (Total)", "IQ_STOCK_BASED_SM", "Income Statement"},
	CapIQMnemonic{"Stock-Based Comp., SG&A Exp. (Total)", "IQ_STOCK_BASED_SGA", "Income Statement"},
	CapIQMnemonic{"Total Operating Expenses", "IQ_TOTAL_OPER_EXPEN", "Income Statement"},
	CapIQMnemonic{"Total Revenues", "IQ_TOTAL_REV", "Income Statement"},
	CapIQMnemonic{"Total Revenues (As Reported)", "IQ_TOTAL_REV_AS_REPORTED", "Income Statement"},
	CapIQMnemonic{"Total Stock-Based Compensation", "IQ_STOCK_BASED_TOTAL", "Income Statement"},
	CapIQMnemonic{"Unusual Items, Total", "IQ_TOTAL_UNUSUAL", "Income Statement"},
	CapIQMnemonic{"Accumulated Allowance for Doubtful Accounts", "IQ_ALLOW_DOUBT_ACCT", "Supplemental"},
	CapIQMnemonic{"Assets on Operating Lease - Accumulated Depreciation", "IQ_ASSETS_OPER_LEASE_DEPR", "Supplemental"},
	CapIQMnemonic{"Assets on Operating Lease - Gross", "IQ_ASSETS_OPER_LEASE_GROSS", "Supplemental"},
	CapIQMnemonic{"Assets under Capital Lease - Accumulated Depreciation", "IQ_ASSETS_CAP_LEASE_DEPR", "Supplemental"},
	CapIQMnemonic{"Assets under Capital Lease - Gross", "IQ_ASSETS_CAP_LEASE_GROSS", "Supplemental"},
	CapIQMnemonic{"AUM End of Period", "IQ_AUM_EOP", "Supplemental"},
	CapIQMnemonic{"AUM Market Appreciation/(Depreciation)", "IQ_AUM_MARKET_APPRECIATION_DEPRECIATION", "Supplemental"},
	CapIQMnemonic{"AUM Net Inflows/(Outflows)", "IQ_AUM_NET_INFLOWS_OUTFLOWS", "Supplemental"},
	CapIQMnemonic{"Current Domestic Taxes", "IQ_CURR_DOMESTIC_TAXES", "Supplemental"},
	CapIQMnemonic{"Current Foreign Taxes", "IQ_CURR_FOREIGN_TAXES", "Supplemental"},
	CapIQMnemonic{"Deferred Domestic Taxes", "IQ_DEFERRED_DOMESTIC_TAXES", "Supplemental"},
	CapIQMnemonic{"Deferred Foreign Taxes", "IQ_DEFERRED_FOREIGN_TAXES", "Supplemental"},
	CapIQMnemonic{"Gross Profit Margin %", "IQ_GROSS_MARGIN", "Financial Ratios"},
	CapIQMnemonic{"Maintenance & Repair Expenses, Total", "IQ_MAINT_REPAIR", "Supplemental"},
	CapIQMnemonic{"Order Backlog", "IQ_ORDER_BACKLOG", "Supplemental"},
	CapIQMnemonic{"Deferred Income Taxes, Total (CF)", "IQ_CHANGE_DEF_TAX_TOTAL", ""},
	CapIQMnemonic{"Total Capital Loss C/F", "IQ_CAP_LOSS_CF_TOTAL", "Supplemental"},
	CapIQMnemonic{"Total Current Taxes", "IQ_CURR_TAXES", "Supplemental"},
	CapIQMnemonic{"Total Deferred Taxes", "IQ_DEFERRED_TAXES_TOTAL", "Supplemental"},
	CapIQMnemonic{"Total NOL C/F", "IQ_NOL_CF_TOTAL", "Supplemental"},
}

func FilterTTM(arr []CapIQMnemonic) []CapIQMnemonic {
	arr2 := make([]CapIQMnemonic, 0, len(arr))

	for _, v := range arr {
		if v.AdditionalInfo == "Income Statement" || v.AdditionalInfo == "Cash Flow" || v.AdditionalInfo == "Financial Ratios" {
			arr2 = append(arr2, v)
		}
	}

	return arr2
}

func SplitIntoTwoLists(arr []CapIQMnemonic) ([]string, []string) {
	names := make([]string, len(arr), len(arr))
	ids := make([]string, len(arr), len(arr))

	for i, v := range arr {
		names[i] = v.Name
		ids[i] = v.Code
	}

	return names, ids
}

func GetData(ticker string, mnemonic string, periodType PeriodType) ([]string, []float64, string) {
	dateArr := make([]string, 0)
	dataArr := make([]float64, 0)
	multiplier, units := GetUnitsAndMultiplier(mnemonic)

	client := &http.Client{}
	reader := strings.NewReader(`inputRequests={inputRequests: [ {function:"GDSHE",identifier:"` + ticker + `:",mnemonic:"` + mnemonic + `",properties:{PeriodType: "` + string(periodType) + `-100", restatementTypeId: "LC",metaDataTag:"PeriodDate"}}  ] }`)
	req, err := http.NewRequest("POST", apiEndpoint, reader)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Basic <YOUR ENCODED PASSWORD HERE>")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Printf("body=%s\n", body)

		var v CapIQData
		json.Unmarshal(body, &v.RawData)

		tempArr := diveOneLevelArray("GDSSDKResponse", v.RawData)

		// fmt.Printf("tempArr=%s\n", tempArr)

		if len(tempArr) > 0 {
			tempArr = diveOneLevelArray("Rows", tempArr[0].(map[string]interface{}))

			for _, v := range tempArr {
				m := v.(map[string]interface{})
				// fmt.Printf("m=%s\n", m)

				row := diveOneLevelArray("Row", m)
				if len(row) == 2 {
					dataStr := row[0].(string)
					dateStr := row[1].(string)

					dataFloat, err := strconv.ParseFloat(dataStr, 64)

					if err == nil && dateStr != "" {
						dateArr = append(dateArr, dateStr)
						dataArr = append(dataArr, dataFloat*multiplier)
					}
				}
			}
		}
	}

	return dateArr, dataArr, units
}

func GetUnitsAndMultiplier(mnemonic string) (float64, string) {
	for _, v := range GetCapIQCodes {
		if v.Code == mnemonic {
			if strings.Contains(v.Name, "%") {
				return 0.01, "%"
			}
			if strings.Contains(v.Name, "Thousands") {
				return 1000, "$"
			}
			if strings.Contains(v.Name, "Employees") {
				return 1, "Employees"
			}
			if v.AdditionalInfo == "Financial Ratios" {
				return 1, "Ratio"
			}
			if strings.Contains(v.Name, "EPS") {
				return 1, "EPS"
			}
		}
	}

	return 1e6, "$"
}

func diveOneLevel(tag string, m map[string]interface{}) map[string]interface{} {
	temp, found := m[tag]

	if found {
		return temp.(map[string]interface{})
	}

	return nil
}

func diveOneLevelArray(tag string, m map[string]interface{}) []interface{} {
	temp, found := m[tag]

	if found {
		return temp.([]interface{})
	}

	return nil
}
