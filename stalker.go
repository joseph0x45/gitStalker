package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Stalk(code string){
	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)
	var content string
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		content = e.Text
		fmt.Printf(" %v ", content)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %v\n", r.URL.String())
	})

	c.Visit("https://github.com/TheWisePigeon/ssj3Store/blob/main/src/index.ts")
}