package main

import (
	"context"
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

	text, err := ioutil.ReadFile(argsWithoutProg[0]) // just pass the file name
	if err != nil {
		fmt.Print(err)
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

	if sentiment.DocumentSentiment.Score >= 0 {
		fmt.Println("Sentiment: positive")
	} else {
		fmt.Println("Sentiment: negative")
	}

	if err := Save("./db/sentiment/migraine.json", sentiment); err != nil {
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

	firstEntity := entitySentiment.Entities[0]

	fmt.Printf("First Entity: %v\n", firstEntity.Name)
	if firstEntity.Sentiment.Score >= 0 {
		fmt.Println("Entity Sentiment: positive")
	} else {
		fmt.Println("Entity Sentiment: negative")
	}

}
