CREATE SCHEMA dumptruck;

-- track all objects across all time series using this id
CREATE TABLE dumptruck.data_master (
  id SERIAL PRIMARY KEY
  , external_id_0 TEXT
  , external_id_1 TEXT
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
  , duplicate_of TEXT REFERENCES data_sets (data_series_name)
  , last_updated TIMESTAMP
);

-- each series should look like this
CREATE TABLE dumptruck.series_1 (
  id INT REFERENCES data_master (id)
  , ts TIMESTAMP WITHOUT TIME ZONE -- store all in UTC
  , val NUMERIC -- should be appropriate type for series
);
