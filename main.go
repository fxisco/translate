package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func main() {
	dest := flag.String("dest", "es", "Destination language for translation.")

	flag.Parse()

	sentence := strings.Join(flag.Args(), " ")

	fmt.Println(*dest, sentence)

	ctx := context.Background()
	client, err := translate.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	target, err := language.Parse(*dest)

	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	translations, err := client.Translate(ctx, []string{sentence}, target, nil)

	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Translation: %v\n", translations[0].Text)
}
