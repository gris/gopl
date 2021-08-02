// Issues exibe tabela de problemas do GitHub que correspondem aos termos
// de pesquisa
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const ONE_MONTH_IN_HOURS = float64(24 * 30)
const ONE_YEAR_IN_HOURS = float64(12 * ONE_MONTH_IN_HOURS)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %17.20s\n", item.Number, item.User.Login, item.Title, formatDate(item.CreatedAt))
	}
}

func formatDate(d time.Time) string {
	if time.Since(d).Hours() < ONE_MONTH_IN_HOURS {
		return "less than a month"
	} else if time.Since(d).Hours() < ONE_YEAR_IN_HOURS {
		return "less than a year"
	} else {
		return "more than a year"
	}
}
