-- track all objects across all time series using this id
CREATE TABLE data_master (
  id SERIAL PRIMARY KEY
  , external_id_0 TEXT
  , external_id_1 TEXT
  , .. 
);

-- basic contact information for each data source provider
CREATE TABLE providers (
  name TEXT PRIMARY KEY
  , contact_email TEXT
  , description TEXT
  , ..
);

-- basic information for each data series
CREATE TABLE data_sets (
  data_series_name TEXT PRIMARY KEY
  , provider TEXT REFERENCES providers (name)
  , source TEXT -- 2nd level description of where data came from
  , license TEXT
  , ingest_script TEXT
  , duplicate_of TEXT REFERENCES data_sets (data_series_name)
  , update_interval INTERVAL
  , ..
);

-- create by some ingest script
CREATE TABLE data_series_1 (
  id INT REFERENCES data_master (id)
  , ts TIMESTAMP WITHOUT TIME ZONE -- store all in UTC
  , val ANY -- actual val will be typed
);
