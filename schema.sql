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
  , source TEXT REFERENCES sources (name)
  , license TEXT
  , duplicate_of TEXT REFERENCES dumptruck.data_sets (name)
  , last_updated TIMESTAMP
  , description TEXT
);

-- each series should look like this
CREATE TABLE dumptruck.series_1 (
  id INT REFERENCES dumptruck.data_master (id)
  , ts TIMESTAMP WITHOUT TIME ZONE -- store all in UTC
  , val NUMERIC -- should be appropriate type for series
  , UNIQUE(id, ts)
);
