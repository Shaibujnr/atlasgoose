package migrations

import (
	"context"
	"database/sql"
	"errors"
	"github.com/pressly/goose/v3"
	"log/slog"
	"strings"
)

var (
	SUPPORTED_TAGS = []string{"golang", "food", "science", "animal", "dog", "cat", "tech", "life", "fun", "book"}
)

func init() {
	goose.AddMigrationContext(upFillTagsColumn, downFillTagsColumn)
}

func getTags(content string) string {
	result := "basic"
	words := strings.Split(content, " ")
	for _, word := range words {
		for _, tag := range SUPPORTED_TAGS {
			if strings.Contains(word, tag) {
				result += " " + tag
			}
		}
	}
	return result
}

func upFillTagsColumn(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.

	// create a cursor for fetching the blog posts
	_, err := tx.ExecContext(ctx, `DECLARE posts_cursor CURSOR FOR SELECT id, content FROM blog_posts;`)
	if err != nil {
		panic(err)
	}

	// defer closing the cursor
	defer func() {
		_, err := tx.ExecContext(ctx, "CLOSE posts_cursor;")
		if err != nil {
			panic(err)
		}
	}()

	var id uint
	var content string

	for {
		err := tx.QueryRowContext(ctx, "FETCH NEXT FROM posts_cursor").Scan(&id, &content)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				break
			}
			slog.Error("Fetching next row failed", "err", err.Error())
			panic(err)
		}
		tags := getTags(content)
		_, err = tx.ExecContext(ctx, "UPDATE blog_posts SET tags=$1 WHERE id=$2;", tags, id)
		if err != nil {
			slog.Error("Unable to update row", "err", err.Error())
			panic(err)
		}
	}
	return nil
}

func downFillTagsColumn(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
