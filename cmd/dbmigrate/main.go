package main

import (
	"context"
	"flag"
	_ "github.com/Shaibujnr/atlasgoose/migrations"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	err := flags.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}

	args := flags.Args()
	slog.Info("Args are", "args", args, "dir", *dir)
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	dsn := "host=db user=tdb password=tdb dbname=tdb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}()

	arguments := make([]string, 0)
	if len(args) > 1 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.RunContext(context.Background(), command, sqlDB, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
