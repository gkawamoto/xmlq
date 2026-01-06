package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
	xml2json "github.com/basgys/goxml2json"
	"github.com/go-xmlfmt/xmlfmt"
)

func main() {
	prettyPrint := flag.Bool("p", false, "pretty print")
	innerText := flag.Bool("t", false, "inner text")
	jsonOutput := flag.Bool("j", false, "json output")

	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	doc, err := xmlquery.Parse(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	args := flag.Args()

	if len(args) == 0 {
		output(doc, *jsonOutput, *prettyPrint)
		return
	}

	for _, q := range args {
		for _, n := range xmlquery.Find(doc, q) {
			if *innerText {
				fmt.Println(n.InnerText())
				continue
			}

			if n.Type == xmlquery.AttributeNode {
				fmt.Println(n.InnerText())
				continue
			}

			output(n, *jsonOutput, *prettyPrint)
		}
	}
}

func output(n *xmlquery.Node, jsonOutput bool, prettyPrint bool) {
	if jsonOutput {
		r, err := xml2json.Convert(strings.NewReader(n.OutputXML(true)))
		if err != nil {
			panic(err)
		}
		fmt.Fprint(os.Stdout, r.String())
		return
	}

	if prettyPrint {
		fmt.Fprint(os.Stdout, xmlfmt.FormatXML(n.OutputXML(true), "", "  "))
		return
	}

	fmt.Fprint(os.Stdout, n.OutputXML(true))
}
