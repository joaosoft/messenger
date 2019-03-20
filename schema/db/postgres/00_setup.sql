
-- migrate up
CREATE SCHEMA IF NOT EXISTS "messenger";

CREATE TABLE "messenger"."message" (
	id_message 		    TEXT PRIMARY KEY,
	"from"		        TEXT NOT NULL,
	"to"		        TEXT NOT NULL,
	message 		    TEXT NOT NULL,
	created_at			TIMESTAMP NOT NULL DEFAULT NOW()
);



-- migrate down
DROP TABLE "messenger"."message";

DROP SCHEMA "messenger";
