package main

import (
	"fmt"

	"github.com/oke-py/contributions/pkg/github"
)

func main() {
	fmt.Println(github.GetContributions().Convert())
}
