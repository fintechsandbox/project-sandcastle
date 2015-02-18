# dumptruck
Git tracked security master and ingest scripts for a variety of financial data providers designed to facilitate financial data retrieval.

## Goals
- Create a high quality security master in an easily parseable (csv?) format for ALL securities worldwide
- Create/compile a set of easy-to-use (python?) scripts that will fetch data from the target providers
- Document core features of each target provider (distribution license, contact info, overarching description of available data sets)

## Target Supported Data Sets / Providers
- Fintech Sandbox partners (http://www.fintechsandbox.org/data-partners)
- Thomson Reuters
- Factset
- Six Financials
- Benzinga
- Stocktwits
- S&P
- Interactive Data
- Alexandria

## Data Series Design

Utilizing the schema described in https://github.com/elsen-trading/dumptruck/blob/master/schema.sql and a specification of data series like:
```
{ 
      "granularity": "day"
      , "start_date": "2008-01-01"
      , "stop_date": "2009-01-01"
      , "data_series": {

            { "name": "groupings", "fields": [ 
                  { "name": "val", "equals": "SP900" }
            ]}
            , { "name": "alexandria", "fields": [ 
                { "name": "sentiment", "min": 0, "max": 1 }
              , { "name": "confidence", "min": 0.75, "max": 1 }
            ]}
            , { "name": "currentprice", "fields": [ 
                { "name": "val", "min": 0, "max": 1 }
            ]}
      }
}
```
we want to be able to produce a query like:
```
SELECT groupings.did, DATE_TRUNC('day', groupings.ts), alexandria.sentiment, alexandria.confidence, currentprice.val
FROM groupings
INNER JOIN alexandria ON groupings.did = alexandria.did 
  AND DATE_TRUNC('day', groupings.ts) = DATE_TRUNC('day', alexandria.ts) 
INNER JOIN currentprice ON groupings.did = currentprice.did 
  AND DATE_TRUNC('day', groupings.ts) = DATE_TRUNC('day', currentprice.ts) 
WHERE groupings.val = 'SP900'
AND groupings.ts BETWEEN '2008-01-01'::TIMESTAMP AND '2009-01-01'::TIMESTAMP
AND alexandria.ts BETWEEN '2008-01-01'::TIMESTAMP AND '2009-01-01'::TIMESTAMP
AND currentprice.ts BETWEEN '2008-01-01'::TIMESTAMP AND '2009-01-01'::TIMESTAMP
AND alexandria.sentiment BETWEEN 0 AND 1
AND alexandria.confidence BETWEEN 0.75 AND 1
AND currentprice.val BETWEEN 0 AND 1;
```
which is equivalent to the sentiment, confidence, and currentprice values for each did (data id) in the S&P900 from 2008-01-01 through 2009-01-01. 
