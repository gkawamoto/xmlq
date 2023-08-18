package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/antchfx/xmlquery"
	"github.com/go-xmlfmt/xmlfmt"
)

func main() {
	prettyPrint := flag.Bool("p", false, "pretty print")
	innerText := flag.Bool("t", false, "inner text")

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

	if len(args) == 0 && *prettyPrint {
		fmt.Println(xmlfmt.FormatXML(doc.OutputXML(*prettyPrint), "", "  "))
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

			if *prettyPrint {
				fmt.Println(xmlfmt.FormatXML(n.OutputXML(*prettyPrint), "", "  "))
				continue
			}

			fmt.Println(n.OutputXML(true))
		}
	}
}
