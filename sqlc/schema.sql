CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "item" (
  "id" bigserial PRIMARY KEY,
  "description" varchar,
  "amount" bigint NOT NULL,
  "likes" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "account" ("username");
