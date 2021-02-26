package main

import (
	"context"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatalln("Must supply file argument")
	}

	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Read 'database' on disk
	// ..TODO

	// Read text from file provided as argument
	text, err := ioutil.ReadFile(argsWithoutProg[0]) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	h := sha1.New()
	h.Write(text)
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x", bs)

	if err := ioutil.WriteFile(fmt.Sprintf("./db/text/%s.text", hash), text, 0644); err != nil {
		log.Fatalln(err)
	}

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: string(text),
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	if err := Save(fmt.Sprintf("./db/sentiment/%s.json", hash), sentiment); err != nil {
		log.Fatalln(err)
	}

	// Detects the subject of the text.
	entitySentiment, err := client.AnalyzeEntitySentiment(ctx, &languagepb.AnalyzeEntitySentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: string(text),
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	if err := Save(fmt.Sprintf("./db/entity-sentiment/%s.json", hash), entitySentiment); err != nil {
		log.Fatalln(err)
	}
}
