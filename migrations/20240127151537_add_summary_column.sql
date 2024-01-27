-- +goose Up
-- modify "blog_posts" table
ALTER TABLE "blog_posts" ADD COLUMN "summary" text;

-- +goose Down
-- reverse: modify "blog_posts" table
ALTER TABLE "blog_posts" DROP COLUMN "summary";
