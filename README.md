# xmlq

`xmlq` is a tool to query a XML document using [XPath](https://www.w3schools.com/xml/xpath_intro.asp). It is written in Go and uses [github.com/antchfx/xmlquery](github.com/antchfx/xmlquery) to parse the XML and query it using XPath.

For information on XPath support, see the documentation of the underlying [github.com/antchfx/xpath lib](https://pkg.go.dev/github.com/antchfx/xpath).

## Installation

```bash
$ go install github.com/gkawamoto/xmlq@main
```

## Usage

```bash
$ xmlq -h
Usage of xmlq:
  -p	pretty print
  -t	inner text
```

## Examples

### Fetching a value from a path

```bash
$ echo "<payload><key>value</key></payload>" | xmlq -t '/payload/key'
value
```

### Formatting a XML

```bash
$ echo "<payload><key>value</key></payload>" | xmlq -p

<?xml version="1.0"?>
<payload>
  <key>value</key>
</payload>
```

## Contributing

Push your changes to a branch and open a pull request.

