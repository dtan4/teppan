package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/higanworks/envmap"
	"github.com/joho/godotenv"
)

const Usage = "usage: teppan [--base64] <template file>"

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

	fmt.Println(base64Encode)

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load .env!")
		os.Exit(1)
	}

	envs := envmap.All()
	templateFile := os.Args[len(os.Args)-1]
	tmpl, err := template.ParseFiles(templateFile)
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
