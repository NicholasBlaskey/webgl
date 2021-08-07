package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"

	"strings"

	"golang.org/x/net/html"
)

const constantsUrl = "https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Constants"

func main() {
	// Get the documentation on constants from mozilla
	res, err := http.Get(constantsUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Parse
	doc, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}

	constants := extractConstants(doc)
	fmt.Println(constants)
}

type constantInfo struct {
	name string
	val  string
	desc string
}

func extractConstants(n *html.Node) []constantInfo {
	if n.Type == html.ElementNode && n.Data == "tbody" {
		parseTable(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractConstants(c)
	}

	return nil
}

func parseTable(n *html.Node) []constantInfo {
	fmt.Println("Parsing table")
	vals := []constantInfo{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "tr" || c.Data == "th") {
			val := getRowValues(c)

			val = append(val, "") // Add empty string to handle cases where there are no comments in the description
			for i := 3; i < len(val); i++ {
				val[2] += val[i]
			}

			vals = append(vals, constantInfo{val[0], val[1], val[2]})
		}
	}
	fmt.Println(vals)
	return vals
}

func getRowValues(n *html.Node) []string {
	if n.Type == html.TextNode {
		if strings.Trim(n.Data, " \n\t\r") == "" {
			return nil
		}
		return []string{n.Data}
	}

	vals := []string{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		vals = append(vals, getRowValues(c)...)
	}
	return vals
}
