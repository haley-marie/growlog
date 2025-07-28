package migrate

import "embed"

//go:embed schema/*.sql
var EmbedMigrations embed.FS
