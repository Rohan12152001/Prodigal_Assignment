-- // Mutual Fund Data Table // --
CREATE TABLE MF_data (
	schemeCode INT NOT NULL,
	schemeName TEXT,
	parentScheme TEXT NOT NULL,
	isinDivPayoutGrowth TEXT,
	isinDivReinvestment TEXT,
	nav FLOAT,
	repurchasePrice FLOAT,
	salePrice FLOAT,
	date TIMESTAMP NOT NULL,
    createdAt TIMESTAMP NOT NULL
);

-- // Indexed on schemeName // --
CREATE INDEX mf_sn ON MF_data USING btree (schemeName);

-- // TimeStampDB for keeping a track of the last record fetched // --
CREATE TABLE timeStampDB (
	fetchFrom TIMESTAMP
)