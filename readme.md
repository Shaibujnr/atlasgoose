# Database Migration Using Atlasgo and Goose

This project demonstrates how to combine atlas and goose for 
database schema and data migration.

* We use Gorm to define database tables
* We use Atlas to automatically generate a schema migration plan
* We use Goose to apply our migration files 
* We use Goose to generate `Go` based migration files for data migrations
* We build a custom `Goose` binary to support `Go` based migrations

You can read the article on this project [here](https://volomn.com/blog/database-migration-using-atlas-and-goose)

## PreRequisites
* Golang 1.23.1
* Docker
* Docker Compose

## Setup
1. Clone this repository
2. Run `go mod tidy`
