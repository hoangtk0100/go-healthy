CREATE TYPE "user_status" AS ENUM (
  'AVAILABLE',
  'LOCKED',
  'REMOVED'
);

CREATE TYPE "user_type" AS ENUM (
  'SYSTEM',
  'ADMIN',
  'USER'
);

CREATE TYPE "gender" AS ENUM (
  'MALE',
  'FEMALE',
  'OTHER'
);

CREATE TYPE "meal_type" AS ENUM (
  'MORNING',
  'LUNCH',
  'DINNER',
  'SNACK'
);

CREATE TABLE "users" (
  "username" varchar(255) PRIMARY KEY,
  "email" varchar(255) UNIQUE NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "phone_number" varchar(16),
  "gender" gender,
  "type" user_type NOT NULL DEFAULT 'USER',
  "status" user_status NOT NULL DEFAULT 'AVAILABLE',
  "old_status" user_status NOT NULL DEFAULT 'AVAILABLE',
  "avatar_url" varchar,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "body_records" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "weight" bigint,
  "body_fat" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "meals" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" varchar,
  "calories" int,
  "type" meal_type NOT NULL DEFAULT 'MORNING',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "exercises" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" varchar,
  "calories_burned" int,
  "duration" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "diaries" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "blog_posts" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "body_records" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "meals" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "exercises" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "diaries" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "blog_posts" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
