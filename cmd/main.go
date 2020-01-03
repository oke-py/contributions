package main

import (
	"github.com/oke-py/contributions/pkg/exporter"
	"github.com/oke-py/contributions/pkg/github"
)

func main() {
	exporter.WriteMarkdown(github.GetContributions().Convert())
}
