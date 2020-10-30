CREATE TABLE greetings(
	id char(36) PRIMARY KEY,
	greeting varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	used boolean NOT NULL DEFAULT false,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
