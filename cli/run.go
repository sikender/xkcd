package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/sikender/xkcd/client"
)

func Run(args []string) int {
	comicNo := flag.Int("n", 0, "Comic number to fetch. Fetches the latest by default")
	save := flag.Bool("s", false, "Save the comic to current directory")
	flag.Parse()

	xkcd := client.NewXKCDClient()
	resp, err := xkcd.Fetch(*comicNo, *save)
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Println(resp.JSON())
	return 0
}
