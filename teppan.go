package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/higanworks/envmap"
	"github.com/joho/godotenv"
)

const Usage = "usage: teppan <template file>"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load .env!")
		os.Exit(1)
	}

	envs := envmap.All()

	templateFile := os.Args[1]
	rawData, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load template file: %s\n", templateFile)
		os.Exit(1)
	}

	tmpl, err := template.New("teppan").Parse(string(rawData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse template file: %s\n", templateFile)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, envs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to apply from template file: %s\n", templateFile)
		os.Exit(1)
	}
}
