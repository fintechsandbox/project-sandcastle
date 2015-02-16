CREATE TABLE object_master (
  id UUID PRIMARY KEY
  , external_id_1
);

CREATE TABLE providers (
  name TEXT PRIMARY KEY
  , contact_email TEXT
);

CREATE TABLE data_sets (
  data_series_name TEXT PRIMARY KEY
  , provider_name TEXT REFERENCES providers (name)
  , license TEXT
  , ingest_script TEXT
);

-- create by some ingest script
CREATE TABLE data_series_1 (
  object_id UUID REFERENCES object_master (id)
  , ts TIMESTAMP 
  , value0 NUMERIC
  , value1 TEXT
  , ..
);
