# Elsen's Experience
Thompson Reuters provides access to thousands of fundamental data points from the three main financial statements and beyond. Information includes typical standardized financial data such as earnings, dividends, and margins, but also contains sector specific data such as business and geographic segments, operating information (business intelligence), and major customers. TR prefers to deliver their data through an online SOAP-based API or XML structured files via FTP.

[Hopefully] helpful files:

1. #####[TR Quant Data sources - TRQA_DataSources042014.pdf](TRQA_DataSources042014.pdf)  
High level document of QA sources and corresponding licenses. Super helpful for quick checks into TR holdings. Note this is only for the QA product line. The six data sets listed below should help point you in the right direction of what you need and then the pdf mentioned will give you the exact detail of sources:
  - Datastream (pricing data)
  - Starmine (factor modeling data)
  - Worldscope (semi-cleaned fundamental data)
  - I/B/E/S (analysts reports)
  - Reuters Knowledge Direct (more fundamentals)
  - Thomson Reuters Business Classification (sector information)

1. #####[tr_RefInfo_AEA8B.xml](tr_RefInfo_AEA8B.xml)
Report number `AEA8B`. Contains high level report information for a specific company which we've used for mapping identifiers, such as `CUSIP` to `ISIN`. One file per company.

1. #####[tr_STDANN_B0905](tr_STDANN_B0905)
Full fundamental report number `B0905`. Contains standardized annual financial statements going back up to 10 years. One file per company. Note the `COA` fields can be mapped to more detailed documentation as defined in `Fundamentals_Glossary_Parsed.xlsx`. TR does a good job of keeping these reports up to date, but you will need to watch for any lagging reports. We had one case where there were over 600 companies whose latest reports were behind all the other companies and the TR support team was pretty quick to track down all 673 missing reports (which they pushed to our ftp shortly after our request). Bottom line, if something seems off, ask. They seem to spend a lot of resources on their support team. 

1. #####[Example RKD download script - tr_from_ftp.sh](tr_from_ftp.sh)  
Example script to download files from the RKD ftp site. 

1. #####[Example fundamental parsing script - tr_parse_file.py](tr_parse_file.py)  
Example script to parse one of the funadmental reports downloaded with `tr_from_ftp.sh`

1. #####[Official Fundamental Glossary - Fundamentals_Glossary.pdf](Fundamentals_Glossary.pdf)  
Official fundamental documentation. Helpful for finding definitions of more obscure fundamentals. Doesn't show how they're all calculated, but if you can find the fundamental code, TR can help you find more information about that specific entity. 

1. #####[(Useful) Official Fundamentals_Glossary_Parsed.xlsx](Fundamentals_Glossary_Parsed.xlsx)  
Parsed version of `Fundamentals_Glossary.pdf`. Helpful if you want to programmatically generate documentation for a subset of all of the fundmentals. Note the descriptions do have some artifacts from the pdf that may need to be removed, depending on which parts you want. 

1. #####[Official TR API Docs - TRKD_API_Developer_Guide.pdf](TRKD_API_Developer_Guide.pdf)  
Official TR api documentation.

# Kyper Data's Experience
Thomson Reuters offers a historical market data service called Thomson Reuters Tick History (TRTH). The data can be requested via a SOAP API or a web-based GUI. Files are delivered via FTP.

CAUTION: If a request is too large (say...more than 5GB), the request may be ignored. So, err on the small side with the number of securities and/or duration of history in each request.
