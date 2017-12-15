package main

import (
	"flag"
	"fmt"
	"github.com/aliate/banner"
)

func main() {
	var origin string
	var adjoin bool

	flag.StringVar(&origin, "origin", "", "origin string, will convert it to banner show.")
	flag.BoolVar(&adjoin, "adjoin", false, "adjoin mode")
	flag.Parse()

	if len(origin) > 0 {
		banner.NewBanner(origin, adjoin).Show()
	} else {
		fmt.Println("Please give a origin string!")
	}
}
