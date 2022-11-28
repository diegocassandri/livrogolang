package main

import (
	"fmt"
	"livrogolang/ch4/github"
	"log"
	"os"
)

// Issues prints a table of GitHub issues matching the search terms.
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
