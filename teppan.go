package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/joho/godotenv"
)

const Usage = "usage: teppan [--base64] <template file>"

func envmap(base64Encode bool) map[string]string {
	items := make(map[string]string)

	for _, env := range os.Environ() {
		keyValue := strings.SplitN(env, "=", 2)

		if base64Encode {
			items[keyValue[0]] = base64.StdEncoding.EncodeToString([]byte(keyValue[1]))
		} else {
			items[keyValue[0]] = keyValue[1]
		}
	}

	return items
}

func main() {
	var (
		base64Encode bool
	)

	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}

	flag.BoolVar(&base64Encode, "base64", false, "Embed Base64-encoded variables")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load .env!")
		os.Exit(1)
	}

	templateFile := os.Args[len(os.Args)-1]
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse template file: %s\n", templateFile)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, envmap(base64Encode))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to apply from template file: %s\n", templateFile)
		os.Exit(1)
	}
}
