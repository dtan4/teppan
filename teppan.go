package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load .env!")
		os.Exit(1)
	}

	fmt.Println(os.Getenv("TEPPAN_SAMPLE"))
}
