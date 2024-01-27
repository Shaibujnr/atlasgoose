package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/Shaibujnr/atlasgoose/models"
	"io"
	"log/slog"
	"os"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&models.User{}, &models.BlogPost{})
	if err != nil {
		slog.Error("Failed to load gorm schema", "err", err.Error())
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		slog.Error("Failed to write gorm schema", "err", err.Error())
		os.Exit(1)
	}
}
