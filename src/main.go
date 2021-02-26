package main

import (
	"context"
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	text := `Anyone who has experienced a migraine knows how debilitating the pain can be. And while there are several medicinal ways to help alleviate migraine symptoms, a recent study has found that getting enough exercise might just do the trick.

While exercise has long been known to help reduce migraine triggers, the new research from the University of Washington has found that just two-and-a-half hours of exercise a week can be especially effective with triggers such as stress, depression, and trouble sleeping.

“It’s a complex relationship, but we know that exercise, generally speaking, helps increase levels of good neurotransmitters, like dopamine, norepinephrine, serotonin, which contribute to not only fewer headaches but also better mood and overall well-being,” said study author Dr. Mason Dyess.

As part of the study, the researchers surveyed more than 4,600 diagnosed with migraines. About 75 percent had 15 or more migraines a month while the other 25 percent had 14 or fewer. Each participant completed a questionnaire about their migraine characteristics, sleep routines, depression, stress, and anxiety. And they also reported how much exercise they got each week.

Researchers then divided participants into five groups by frequency of exercise ranging from none to more than 150 minutes (two-and-a-half-hours) per week — the minimum recommended by the World Health Organization (WHO).

The results showed that only 27 percent of participants exercised more than 150 minutes per week, while those who exercised less than that reported increased rates of depression, anxiety, and sleep problems.

What’s more, almost half of the participants who did not exercise had depression, compared to a quarter of people in the group that exercised the most. Four in 10 people who did not exercise also had anxiety, compared to only 28 percent of the group exercising the most.`

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	if sentiment.DocumentSentiment.Score >= 0 {
		fmt.Println("Sentiment: positive")
	} else {
		fmt.Println("Sentiment: negative")
	}


	// Detects the subject of the text.
	entitySentiment, err := client.AnalyzeEntitySentiment(ctx, &languagepb.AnalyzeEntitySentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	if entitySentiment.Entities[0].Sentiment.Score >= 0 {
		fmt.Println("Sentiment: positive")
	} else {
		fmt.Println("Sentiment: negative")
	}


}
