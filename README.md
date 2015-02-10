# dumptruck
Git tracked security master and ingest scripts for a variety of financial data providers designed to facilitate financial data retrieval.

## Goals
- One click [Docker?] image deployment automatically exposes data retrieval API
- Image automates fetching, parsing, and repackaging data via API calls
- Community versioned security master is used to link disparate data sets and allows for easy addition of new data sets

## Target Supported Data Sets / Providers
- Start with Fintech Sandbox partners (http://www.fintechsandbox.org/data-partners)
- Thomson Reuters
- Factset
- Six Financials
- Benzinga
- Stocktwits
- S&P
- Interactive Data
- Alexandria

Each data source should provide:
- provider
- is it proprietary or redistributed? If redistributed, by whom?
- descriptions
- granularity
- update frequency
- methodology for creation, if applicable
- historical coverage
