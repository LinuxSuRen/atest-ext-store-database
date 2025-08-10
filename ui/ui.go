package ui

import (
	_ "embed"
)

//go:embed dist/atest-ext-store-database.umd.js
var js string

//go:embed dist/atest-ext-store-database.css
var css string

func GetJS() string {
	return js
}

func GetCSS() string {
	return css
}
