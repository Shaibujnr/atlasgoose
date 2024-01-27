# Database Migration Using Atlasgo and Goose


# PreRequisites
* Atlasgo (brew install ariga/tap/atlas on Macos version atlas version v0.18.1-c64cfd1-canary)
* Docker
* Docker Compose

# Install Gorm
`go get -u gorm.io/gorm`

NOTE: We need to run `atlas migrate hash` whenever we edit a migration file (we removed the not null when we first 
added the summary column)