package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/oke-py/contributions/pkg/exporter"
	"github.com/oke-py/contributions/pkg/github"
)

func main() {
	user := flag.String("u", "oke-py", "GitHub account to aggregate contributions")
	year := flag.Int("y", 0, "year to aggregate contributions")
	month := flag.String("m", "", "month to aggregate contributions (ex 2020/1)")
	flag.Parse()

	var from, to time.Time
	loc, _ := time.LoadLocation("Asia/Tokyo")

	if *year != 0 {
		from = time.Date(*year, 1, 1, 0, 0, 0, 0, loc)
		to = time.Date(*year+1, 1, 1, 0, 0, 0, 0, loc)
	}

	if *month != "" {
		a := strings.Split(*month, "/")
		if len(a) != 2 {
			fmt.Println("invalid argument")
			os.Exit(1)
		}

		y, err := strconv.Atoi(a[0])
		if err != nil {
			fmt.Println("invalid argument")
			os.Exit(1)
		}

		m, err := strconv.Atoi(a[1])
		if err != nil {
			fmt.Println("invalid argument")
			os.Exit(1)
		}

		from = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, loc)
		if m == 12 {
			to = time.Date(y+1, 1, 1, 0, 0, 0, 0, loc)
		} else {
			to = time.Date(y, time.Month(m+1), 1, 0, 0, 0, 0, loc)
		}
	}

	fmt.Println(exporter.WriteMarkdown(github.GetContributions(*user, from, to).Convert()))
}
