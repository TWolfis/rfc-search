package main

import (
	"flag"
	"fmt"
	"os"
	"rfc-search/internal/fetch"

	"rfc-search/internal/search"
)

// CMD set for searching documents

func main() {

	// CMD set for fetching documents

	fetchCmd := flag.NewFlagSet("fetch", flag.ExitOnError)
	fetchDstDir := fetchCmd.String("dst", "docs", "destination directory")
	fetchMax := fetchCmd.Int("max", 50, "maximum number of documents to fetch")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchSrcDir := searchCmd.String("src", "docs", "source directory")
	searchStr := searchCmd.String("search", "", "search string")

	if len(os.Args) < 2 {
		fmt.Println("expected 'fetch' or 'search' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "fetch":
		fetchCmd.Parse(os.Args[2:])
		err := fetch.FetcRfc(*fetchDstDir, *fetchMax)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "search":
		searchCmd.Parse(os.Args[2:])
		rl := search.NewRfcList(*searchSrcDir)
		rfc, err := rl.SearchRfc(*searchStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println(rfc.Title)

	default:
		fmt.Println("expected 'fetch' or 'search' subcommands")
	}

	fmt.Println("Done")
}
