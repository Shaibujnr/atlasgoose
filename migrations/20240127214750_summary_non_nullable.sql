-- +goose Up
-- modify "blog_posts" table
ALTER TABLE "blog_posts" ALTER COLUMN "summary" SET NOT NULL;

-- +goose Down
-- reverse: modify "blog_posts" table
ALTER TABLE "blog_posts" ALTER COLUMN "summary" DROP NOT NULL;
