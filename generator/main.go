package main

import (
	"os"
	"text/template"
)

func routingMethods() ([]string, string, string) {
	httpMethods := []string{
		"CONNECT", "DELETE", "GET", "HEAD",
		"OPTIONS", "PATCH", "POST", "PUT", "TRACE",
	}
	snip := `"goa design Routing {{.}}": {
    "prefix": "goa-Routing-{{.}}",
    "body": "{{.}}(\"${1:routing}\")"
},
`
	return httpMethods, "routings", snip
}

func attributeAndAliases() ([]string, string, string) {
	// https://github.com/goadesign/goa/blob/c9cbfd74c4eb09677fe0f899ae5988119bfb12d1/design/apidsl/attribute.go#L69
	// Here are all the valid usage of the Attribute function:
	//
	//	Attribute(name string, dataType DataType, description string, dsl func())
	//
	//	Attribute(name string, dataType DataType, description string)
	//
	//	Attribute(name string, dataType DataType, dsl func())
	//
	//	Attribute(name string, dataType DataType)
	//
	//	Attribute(name string, dsl func())	// dataType is String or Object (if DSL defines child attributes)
	//
	//	Attribute(name string)			// dataType is Strin

	snip := `
"goa design {{.}} declaration": {
    "prefix": "goa-{{.}}",
    "body": "{{.}}(\"${1:name}\")",
    "description": "Defines an {{.}} with name only"
},
"goa design {{.}} declaration with dsl": {
    "prefix": "goa-{{.}}-dsl",
    "body": [
        "{{.}}(\"${1:name}\", func() {",
        "\t$2",
        "\\})"
    ],
    "description": "Defines an {{.}} with func for DSL"
},
"goa design {{.}} declaration with dataType": {
    "prefix": "goa-{{.}}-dataType",
    "body": "{{.}}(\"${1:name}\", ${2:dataType})",
    "description": "Defines an {{.}} with dataType"
},
"goa design {{.}} declaration with dataType and dsl": {
    "prefix": "goa-{{.}}-dataType-dsl",
    "body": [
        "{{.}}(\"${1:name}\", ${2:dataType}, func() {",
        "\t$3",
        "\\})"
    ],
    "description": "Defines an {{.}} with dataType and func for DSL"
},
"goa design {{.}} declaration with dataType and description": {
    "prefix": "goa-{{.}}-dataType-description",
    "body": "{{.}}(\"${1:name}\", ${2:dataType}, \"${3:$1}\")",
    "description": "Defines an {{.}} with dataType, description and func for DSL"
},
"goa design {{.}} declaration with dataType, description and dsl": {
    "prefix": "goa-{{.}}-dataType-description-dsl",
    "body": [
        "{{.}}(\"${1:name}\", ${2:dataType}, \"${3:$1}\", func() {",
        "\t$4",
        "\\})"
    ],
    "description": "Defines an {{.}} with dataType, description and func for DSL"
},
`
	return []string{"Attribute", "Param", "Member", "Header"}, "attribute_and_alias", snip
}

func render(targets []string, name string, snippet string) {
	tmpl := template.Must(template.New("name").Parse(snippet))

	for _, t := range targets {
		tmpl.Execute(os.Stdout, t)
	}
}

func main() {
	render(routingMethods())
	render(attributeAndAliases())
}
