CREATE TYPE "device_type" AS ENUM (
  'IOS',
  'ANDROID',
  'WEB_APP'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "device_type" device_type NOT NULL,
  "device_id" varchar(64),
  "device_model" varchar(64),
  "device_token" varchar,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");