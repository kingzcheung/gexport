package static

import "embed"

//go:embed dist/**
var WebDict embed.FS
