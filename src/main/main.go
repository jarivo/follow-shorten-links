package main

import (
	"fmt"
	"flag"
	"os"
	"net/http"
	"net/url"
	"strings"
)

var Schemes = []string{"http://","https://"}

var (
	ShortenUrl string
	FullOutput bool
	RemoveQuery bool
	Redirects []*http.Request
)

func init() {
	flag.StringVar(&ShortenUrl, "u", "", "Shorten URL to follow")
	flag.BoolVar(&FullOutput, "f", false, "Show all redirects in output")
	flag.BoolVar(&RemoveQuery, "r", true, "Remove query strings from URL")
}

func main() {
	flag.Parse()
	
	argsOk := true

	if ShortenUrl == "" {
		argsOk = printErr("Error: URL is not defined, use -u to specify a URL")
	}

	if !contains(ShortenUrl, Schemes) {
		ShortenUrl = strings.Join([]string{Schemes[1], ShortenUrl}, "")
	}

	u, err := url.ParseRequestURI(ShortenUrl)
	if err != nil {
		argsOk = printErr(err.Error())
	}

	if !argsOk {
		printUsage()
	}

	followShortenUrl(u)
}

func followShortenUrl(u *url.URL) {
	client := &http.Client{
		CheckRedirect: logRedirect,
	}

	res, err := client.Head(u.String())
	
	if err != nil {
		printErr(err.Error())
		printUsage()
	}
	
	if FullOutput {
		for index, element := range Redirects {
			fmt.Printf("%v: %v\n", index, element.URL.String())
		}
	}

	fmt.Printf("--> %v\n", res.Request.URL.String())
}

func logRedirect(req *http.Request, via []*http.Request) error {
	if RemoveQuery {
		req.URL.RawQuery = ""
	}

	Redirects = via
	return nil
}

func contains(s string, substrings []string) bool {
	for _, i := range substrings {
		if strings.Contains(s, i) {
			return true
		}
	}

	return false
}

func printUsage() {
	fmt.Println("USAGE: follow-shorten-url [OPTIONS] -u url")
	flag.PrintDefaults()
	os.Exit(1)
}

func printErr(err string) bool {
	fmt.Fprintln(os.Stderr, err)
	return false
}