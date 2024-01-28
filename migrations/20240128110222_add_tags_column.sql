-- +goose Up
-- modify "blog_posts" table
ALTER TABLE "blog_posts" ADD COLUMN "tags" text;

-- +goose Down
-- reverse: modify "blog_posts" table
ALTER TABLE "blog_posts" DROP COLUMN "tags";
