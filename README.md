# dumptruck
Git tracked security master and ingest scripts for a variety of financial data providers designed to facilitate financial data retrieval.

## Goals
- Create a high quality security master in an easily parseable (csv?) format for ALL securities worldwide
- Create/compile a set of easy-to-use (python?) scripts that will fetch data from the target providers and store them in a time series format (described in schema.sql)
- Document core features of each data source (distribution license, contact info, overarching description of available data sets)

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

For maintainability and ease of understanding, all data sets should be decomposed into single value time series. For example, a table such as 
```
uuid, ts, alexandira.sentiment, alexandria.confidence, currentprice.val
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:00, 1, 0.81, 14.32
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:01, 1, 0.68, 14.35
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:00, -1, 0.31, 23.78
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:01, 0, 0.52, 24.29
```
should be broken into three time series like
```
uuid, ts, alexandira.sentiment
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:00, 1
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:01, 1
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:00, -1
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:01, 0

uuid, ts, alexandria.confidence
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:00, 0.81
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:01, 0.68
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:00, 0.31
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:01, 0.52

uuid, ts, currentprice.val
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:00, 14.32
f792f39e-75f9-41b9-afab-bf0a7f67183f, 2008-01-01 09:01, 14.35
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:00, 23.78
b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4, 2008-01-01 09:01, 24.29
```
where the id (uuid in this example) is converted to a listed identifier in the data_master table. If there is not yet an identifier for this object, a new one should be generated with the relevant details added. Each of the three time series should be "checked in" to the data_sets table, allowing for universal discoverability. The final result would be a data layout like:
```
# select id, external_id_0 from dumptruck.data_master;

 id |            external_id_0             
----+--------------------------------------
  1 | f792f39e-75f9-41b9-afab-bf0a7f67183f
  2 | b5aa4745-1c3c-418b-b2f7-ad7bf3aa23f4

# select name,description from dumptruck.data_sets;

         name          |                       description                        
-----------------------+----------------------------------------------------------
 alexandria_confidence | Confidence predictor derived from the Alexandria dataset
 alexandria_sentiment  | Sentiment score derived from the Alexandria dataset
 currentprice          | Closing price of day before derived from EODDATA dataset

# select * from dumptruck.alexandria_sentiment ;

 id |         ts          | val 
----+---------------------+-----
  1 | 2008-01-01 09:00:00 |   1
  1 | 2008-01-01 09:01:00 |   1
  2 | 2008-01-01 09:00:00 |  -1
  2 | 2008-01-01 09:01:00 |   0

# select * from dumptruck.alexandria_confidence ;

 id |         ts          | val  
----+---------------------+------
  1 | 2008-01-01 09:00:00 | 0.81
  1 | 2008-01-01 09:01:00 | 0.68
  2 | 2008-01-01 09:00:00 | 0.31
  2 | 2008-01-01 09:01:00 | 0.52
  
# select * from dumptruck.currentprice ;

 id |         ts          |  val  
----+---------------------+-------
  1 | 2008-01-01 09:00:00 | 14.32
  1 | 2008-01-01 09:01:00 | 14.35
  2 | 2008-01-01 09:00:00 | 23.78
  2 | 2008-01-01 09:01:00 | 24.29
```

## Requirements
- psycopg2 > 2.5 

## Testing 
```
python test.py
```
