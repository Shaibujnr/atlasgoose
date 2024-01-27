data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/atlasloader",
  ]
}

env "dev" {
  src = data.external_schema.gorm.url
  dev = "postgres://tdb:tdb@devdb:5432/tdb?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations?format=goose"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}