/*
Sample factset schema derived from factset CSVs. 
Includes example ingest process using psql.
*/

DROP TABLE factset_prices;
CREATE TABLE factset_prices (
	FS_PERM_SEC_ID TEXT 
	, day DATE
	, ADJDATE DATE
	, CURRENCY TEXT
	, P_PRICE NUMERIC
	, P_PRICE_OPEN NUMERIC
	, P_PRICE_HIGH NUMERIC
	, P_PRICE_LOW NUMERIC
	, P_VOLUME NUMERIC
	, UNIQUE(FS_PERM_SEC_ID, day)
);
GRANT SELECT ON factset_prices TO gui;

-- initial ingest from file
-- for i in `ls ~/Downloads/factset/`
-- do
-- 	cat $i | psql -h master-db.clrssu6hwyib.us-east-1.rds.amazonaws.com invest dbadmin -c "copy factset_prices from stdin with delimiter '|' header csv"
-- 	cat $i | psql -h elsen-db-master invest dbadmin -c "copy factset_prices from stdin with delimiter '|' header csv"
-- done

CREATE TABLE factset_security_master (
	FS_PERM_SEC_ID TEXT
	, ISIN CHAR(12) PRIMARY KEY
	, PROPER_NAME TEXT
	, FACTSET_ENTITY_ID TEXT
	, FS_PRIMARY_EQUITY_ID TEXT
	, FS_PRIMARY_LISTING_ID TEXT
	, INACTIVE_FLAG BIT(1)
	, FREF_SECURITY_TYPE TEXT
);

-- initial ingest from file
-- cat h_security_isin.txt | psql -h master-db.clrssu6hwyib.us-east-1.rds.amazonaws.com invest dbadmin -c "copy factset_security_master from stdin with delimiter '|' csv"
