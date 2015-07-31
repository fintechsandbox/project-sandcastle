Thompson Reuters provides access to thousands of fundamental data points from the three main financial statements and beyond. Information includes typical standardized financial data such as earnings, dividends, and margins, but also contains sector specific data such as business and geographic segments, operating information (business intelligence), and major customers. TR prefers to deliver their data through an online SOAP-based API or XML structured files via FTP.

TR has become Elsen's primary provider, particularly for fundamental data and price information from the DataStream, WorldScope, and Reuters Knowledge Direct system. 

## tr_RefInfo_AEA8B.xml
Report number `AEA8B`. Contains high level report information for a specific company which we've used for mapping identifiers, such as `CUSIP` to `ISIN`. One file per company.

## tr_STDANN_B0905
Full fundamental report number `B0905`. Contains standardized annual financial statements going back up to 10 years. One file per company. Note the `COA` fields can be mapped to more detailed documentation as defined in `Fundamentals_Glossary_Parsed.xlsx`. 

## tr_from_ftp.sh
Example script to download files from the RKD ftp site. 

## tr_parse_file.py
Example script to parse one of the funadmental reports downloaded with `tr_from_ftp.sh`

## Fundamentals_Glossary.pdf
Official fundamental documentation.

## Fundamentals_Glossary_Parsed.xlsx
Parsed version of the official fundamental documentation. Helpful if you want to programmatically generate documentation for a subset of all of the fundmentals. 
