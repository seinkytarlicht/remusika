package frontend

import "embed"

//go:embed .output/public/*
var UI embed.FS
