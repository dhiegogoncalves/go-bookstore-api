DROP KEYSPACE IF EXISTS oauth;

CREATE KEYSPACE IF NOT EXISTS oauth WITH replication={'class' : 'SimpleStrategy', 'replication_factor':1};

USE oauth;

CREATE TABLE IF NOT EXISTS access_tokens(
	access_token VARCHAR PRIMARY KEY,
	user_id BIGINT,
	client_id BIGINT,
	expires BIGINT
);