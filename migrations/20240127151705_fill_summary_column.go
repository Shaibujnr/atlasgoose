package migrations

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pressly/goose/v3"
	"log/slog"
)

func init() {
	goose.AddMigrationContext(upFillSummaryColumn, downFillSummaryColumn)
}

func upFillSummaryColumn(ctx context.Context, tx *sql.Tx) error {
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
		summary := fmt.Sprintf("%s-Summary", content[:2])
		_, err = tx.ExecContext(ctx, "UPDATE blog_posts SET summary=$1 WHERE id=$2;", summary, id)
		if err != nil {
			slog.Error("Unable to update row", "err", err.Error())
			panic(err)
		}
	}
	return nil
}

func downFillSummaryColumn(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, "UPDATE blog_posts SET summary=NULL")
	if err != nil {
		slog.Error("Unable to update row", "err", err.Error())
		panic(err)
	}
	return nil
}
