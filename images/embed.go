package images

import "embed"

//go:embed regular/* shiny/*
var PokemonImages embed.FS
