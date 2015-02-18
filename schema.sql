CREATE TABLE data_master (
  did UUID PRIMARY KEY
  , external_id_0 TEXT
  , external_id_1 TEXT
  , .. 
);

CREATE TABLE providers (
  name TEXT PRIMARY KEY
  , contact_email TEXT
);

CREATE TABLE data_sets (
  data_series_name TEXT PRIMARY KEY
  , provider TEXT REFERENCES providers (name)
  , source TEXT -- 2nd level description of where data came from
  , license TEXT
  , ingest_script TEXT
  , duplicate_of TEXT REFERENCES data_sets (data_series_name)
);

-- create by some ingest script
CREATE TABLE data_series_1 (
  did UUID REFERENCES data_master (did)
  , ts TIMESTAMP 
  , value0 NUMERIC
  , value1 TEXT
  , ..
);
