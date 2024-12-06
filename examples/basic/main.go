package main

import (
	"fmt"
	"github.com/mikeee/altalabs-go"
	"os"
)

func main() {
	client, err := altalabs.NewAltaClient(os.Getenv("SDK_ALTA_USER"), os.Getenv("SDK_ALTA_PASS"))
	if err != nil {
		panic(err)
	}

	sites, err := client.ListSites()
	if err != nil {
		panic(err)
	}

	for _, site := range sites {
		fmt.Println("Site found:", site.Name)
	}

}
