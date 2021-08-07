package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"

	"strconv"
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
	generateCode(constants)
}

func generateCode(constants []constantInfo) {
	fmt.Println(`package webgl

const (
`)
	for _, c := range constants {
		if c.desc != "" {
			c.desc = "//" + c.desc
		}
		fmt.Println(c.name + " = " + c.val + " " + c.desc)
	}
	fmt.Println(")")
}

type constantInfo struct {
	name string
	val  string
	desc string
}

func extractConstants(n *html.Node) []constantInfo {
	vals := []constantInfo{}
	if n.Type == html.ElementNode && n.Data == "tbody" {
		vals = append(vals, parseTable(n)...)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		vals = append(vals, extractConstants(c)...)
	}

	return vals
}

func parseTable(n *html.Node) []constantInfo {
	vals := []constantInfo{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "tr" || c.Data == "th") {
			val := getRowValues(c)

			// Add empty string to handle cases where there are no comments in the description
			// Add more rows to the third string value, to handle cases with <code>tag</code>.
			val = append(val, "")
			for i := 3; i < len(val); i++ {
				val[2] += val[i]
			}

			// Only add the value if we have a hex number in our second cell.
			if _, err := strconv.ParseInt(val[1], 0, 64); err == nil {
				vals = append(vals, constantInfo{val[0], val[1], val[2]})
			}

		}
	}
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
