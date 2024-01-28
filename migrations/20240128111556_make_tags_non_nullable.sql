-- +goose Up
-- modify "blog_posts" table
ALTER TABLE "blog_posts" ALTER COLUMN "tags" SET NOT NULL;

-- +goose Down
-- reverse: modify "blog_posts" table
ALTER TABLE "blog_posts" ALTER COLUMN "tags" DROP NOT NULL;
