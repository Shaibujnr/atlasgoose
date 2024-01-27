-- +goose Up
-- create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "email" text NOT NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create "blog_posts" table
CREATE TABLE "blog_posts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_blog_posts" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "blog_posts_title_key" to table: "blog_posts"
CREATE UNIQUE INDEX "blog_posts_title_key" ON "blog_posts" ("title");
-- create index "idx_blog_posts_deleted_at" to table: "blog_posts"
CREATE INDEX "idx_blog_posts_deleted_at" ON "blog_posts" ("deleted_at");

-- +goose Down
-- reverse: create index "idx_blog_posts_deleted_at" to table: "blog_posts"
DROP INDEX "idx_blog_posts_deleted_at";
-- reverse: create index "blog_posts_title_key" to table: "blog_posts"
DROP INDEX "blog_posts_title_key";
-- reverse: create "blog_posts" table
DROP TABLE "blog_posts";
-- reverse: create index "users_email_key" to table: "users"
DROP INDEX "users_email_key";
-- reverse: create index "idx_users_deleted_at" to table: "users"
DROP INDEX "idx_users_deleted_at";
-- reverse: create "users" table
DROP TABLE "users";
