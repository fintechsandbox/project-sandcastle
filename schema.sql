CREATE TABLE data_master (
  did UUID PRIMARY KEY
  , external_id_0 TEXT
  , external_id_1 TEXT
  , .. 
);

CREATE TABLE sources (
  , id SERIAL
  , name TEXT 
  , contact_email TEXT
  , url TEXT 
);

CREATE TABLE data_sets (
  data_series_name TEXT PRIMARY KEY
  , source INT REFERENCES sources (id)
  , license TEXT
  , ingest_script TEXT
  , duplicate_of TEXT -- points to a data_series_name
);

-- create by some ingest script
CREATE TABLE data_series_1 (
  did UUID REFERENCES data_master (did)
  , ts TIMESTAMP 
  , value0 NUMERIC
  , value1 TEXT
  , ..
);
