package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var POTENTIAL_SENSITIVE_INFO = [...]string {
	"API_KEY",
	"GOOGLE_CREDENTIALS",
	"API_TOKEN",
	"DB_URL",
	"FIREBASE_CREDENTIALS",
	"PROJECT_ID",
	"STORAGE_BUCKET",
	"listen",
}

func checkSanity( code string ) ( clean bool ){
	for _, v :=range POTENTIAL_SENSITIVE_INFO{
		if strings.Contains(code, v){
			return false
		}
	}
	return true
}

func Stalk(repo string) ( clean bool) {
	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)
	var content string
	defer func ()  {
		clean = checkSanity(content)
	}()
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		content += e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %v\n", r.URL.String())
	})

	c.Visit(repo)
	return
}
