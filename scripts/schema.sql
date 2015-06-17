/*
Sample timeseries database schema that provides
basic relations between a data master, data sources, 
data sets, and any number of data series. Includes
an example of how to directly ingest time series 
from CSV into the data tables using 
psql (postgres' command prompt).
*/

-- recommend always using a schema definition
-- as the 'public' schema is subject to change by other apps
DROP SCHEMA IF EXISTS dumptruck CASCADE;
CREATE SCHEMA dumptruck;

-- track all objects across all time series using this id
CREATE TABLE dumptruck.data_master (
  id SERIAL PRIMARY KEY
  , display_name TEXT
  , ric VARCHAR(7)
  , sedol CHAR(7)
  , cusip CHAR(9)
  , isin CHAR(12)
  , quantopian_id INT
);

-- basic contact information for each data source 
CREATE TABLE dumptruck.sources (
  name TEXT PRIMARY KEY
  , contact_name TEXT
  , contact_email TEXT
  , contact_phone TEXT
  , description TEXT
);

-- basic information for each data series
CREATE TABLE dumptruck.data_sets (
  name TEXT PRIMARY KEY
  , description TEXT DEFAULT NULL
  , val_type TEXT DEFAULT NULL -- should be one of types listed on http://www.postgresql.org/docs/9.4/static/datatype.html
  , source TEXT DEFAULT NULL REFERENCES dumptruck.sources (name)
  , license TEXT DEFAULT NULL
  , duplicate_of TEXT DEFAULT NULL REFERENCES dumptruck.data_sets (name)
  , last_updated TIMESTAMP DEFAULT NULL
);

-- each series should look like this
CREATE TABLE dumptruck.series_1 (
  id INT REFERENCES dumptruck.data_master (id)
  , ts TIMESTAMP WITHOUT TIME ZONE -- store all in UTC
  , val NUMERIC -- should be appropriate type for series
  , UNIQUE(id, ts)
);

-- how to ingest directly to a data table from a csv 
-- psql -c "COPY (name, description, val_type) dumptruck.data_sets FROM STDIN CSV DELIMITER ','" < data_sets.csv
