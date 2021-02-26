# Automatically select text arguing a different opinion to a given text

The BHAG is to show people ideas that directly contradict each other. An example would be Trumps tweet about the number of people at his inauguration being the highest ever, which would then find and show information about Obama’s inauguration.

Facebook could use this in stead of filtering: whenever someone sees a message claiming A, Facebook would show a message next to it, claiming not-A. And the other way around. This feature would make people think about what they consider to be the truth, without making claims about truthfulness. You may think of it as an automated, large scale version of the weekly feature in Dutch newspaper `NRC` (www.nrc.nl), called the ‘twist gesprek’ (https://www.nrc.nl/rubriek/twistgesprek/) where two people discuss a statement during the week by exchanging messages. The results of this discussion are presented as a back-and-forth of messages in the Saturday opinion pages, like a civilised version of a political debate.

I like reading that kind of thing because it always makes me question my beliefs.
Don’t hide opinions from me, like the current sentiment seems to be, but make me think. As usual, the Internet offers technology that has obvious benefits to how we used to do things in the past, but those benefits come at a price.
The challenge is to keep the benefits without paying the price. And that is what caused this research.

```
How do I extract the main idea out of piece of text in a way 
that would allow me to select another piece of text with the opposite opinion?
```

ideas

*   AutoML Natural Language: https://cloud.google.com/natural-language#natural-language-api-demo
*   Social Media Analytics research by Sandjai Bhulai: https://www.math.vu.nl/~sbhulai/
*   Overview on Wikipedia: https://en.wikipedia.org/wiki/Automatic_summarization
*   About the problems of filtering and moderation at Facebook: https://www.eff.org/deeplinks/2021/02/facebooks-latest-proposed-policy-change-exemplifies-trouble-moderating-speech-0


find a number of messages to use as a training set 
    
try out googles natural language api
    this is a generic model 
    but it's manageable for a days work.


Natural Language API - Sentiment analysis

    Understand the overall opinion, feeling, or attitude sentiment expressed in a block of text. 



./google-cloud-sdk/install.sh
export GOOGLE_APPLICATION_CREDENTIALS=/Users/jan/Downloads/albert-brand-speeltuin-bcfa65f1b1db.json
gcloud auth activate-service-account --key-file=/Users/jan/downloads/albert-brand-speeltuin-bcfa65f1b1db.json
gcloud ml language analyze-entities --content-file=./README.md > test.json


```
gcloud ml language classify-text --content-file=./README.md
{
  "categories": [
    {
      "confidence": 0.55,
      "name": "/News"
    }
  ]
}
```

```
 $ gcloud ml language classify-text --content-file=./text2.txt
{
  "categories": [
    {
      "confidence": 0.76,
      "name": "/Arts & Entertainment"
    }
  ]
}
```


```
gcloud ml language analyze-sentiment --content-file=./text2.txt

    {
      "sentiment": {
        "magnitude": 0.9,
        "score": -0.9
      },
      "text": {
        "beginOffset": 340,
        "content": "the world is broken.we are doomed!"
      }
    },
    {
      "sentiment": {
        "magnitude": 0.2,
        "score": 0.2
      },
      "text": {
        "beginOffset": 375,
        "content": "it wil all end in tears."
      }
    },
    {
      "sentiment": {
        "magnitude": 0.9,
        "score": -0.9
      },
      "text": {
        "beginOffset": 400,
        "content": "the world is broken."
      }
    }
```

```
 $ gcloud ml language analyze-sentiment --content-file=./hoeraText.txt
{
  "documentSentiment": {
    "magnitude": 43.8,
    "score": 0.5
  },
  "language": "en",
  "sentences": [
    {
      "sentiment": {
        "magnitude": 0.2,
        "score": 0.2
      },
      "text": {
        "beginOffset": 0,
        "content": "The sun is shining."
      }
    },
```

idea:

    this way we can find the sentiment of each sentence  in a time line.  if sentiment is 
    negative overall, we might insert a couple of positive sentences.

`gcloud ml language analyze-sentiment --content-file=optimist-migrane.txt > optimist-migrane.json`

gets a neutral overall score. how about taking the most positive sentence and using that as a score?

```
 $ gcloud ml language classify-text --content-file=optimist-migrane.txt
{
  "categories": [
    {
      "confidence": 0.95,
      "name": "/Health/Health Conditions/Pain Management"
    }
  ]
}
```

`gcloud ml language analyze-sentiment --content="Just recently, Joe Biden renewed his pledge to TAKE YOUR LAWFULLY OWNED FIREARMS with the help of his gun-hating friends in Congress."`
```
{
  "documentSentiment": {
    "magnitude": 0.4,
    "score": 0.4
  },
  "language": "en",
  "sentences": [
    {
      "sentiment": {
        "magnitude": 0.4,
        "score": 0.4
      },
      "text": {
        "beginOffset": 0,
        "content": "Just recently, Joe Biden renewed his pledge to TAKE YOUR LAWFULLY OWNED FIREARMS with the help of his gun-hating friends in Congress."
      }
    }
  ]
}
```

TODO: 
- find a number of sample documents
- classify-text 

- generate a list of texts on /Health/Health Conditions/Pain Management. These could be no more than fakes.
- give each text a sentiment, some positive, some negative
- analyze a new text (e.g. the optimist article)
- select another article in the same category that has the opposite sentiment (or a different sentiment?)

This article has a negative sentiment about migraine
```
      "metadata": {},
      "name": "migraine",
      "salience": 0.097846925,
      "sentiment": {
        "magnitude": 0.3,
        "score": -0.3
      },
```

use documentSentiment
```
{
  "documentSentiment": {
    "magnitude": 4.4,
    "score": 0.0
  },
```

or

- if someone sees an article from cnn, show them an article from fox

input: some text in a file 
output: alternative texts, so texts about the same subject, in our case the word with the highest salience. 
        we use this in stead of the classifier because these categories are too generic.

1. get a list of documents about 3 different subjects. Store their sentiment in a database (a json file?) 
1. analyze a new document, find its sentiment, select another document about the same sentiment.


gcloud ml language analyze-entity-sentiment --content-file="myfile" 
    select: entities[0].name > subject 

use `subject` to match docuents in database

gcloud ml language analyze-sentiment --content-file="myfile"
    select: documentSentiment.score


Find `alternative facts` given a `fact` is difficult, at least with Google ML. 

# Summary 

1. Used `gcloud ml language` to analyze text 
1. Find out what a text is about: `gcloud ml language classify-text --content-file=optimist-migrane.txt`
   (this is the text: https://www.optimistdaily.com/2021/02/just-2-5-hours-of-weekly-exercise-could-help-reduce-migraine-triggers/)
1. This is spot-on, super.
```
 $ gcloud ml language classify-text --content-file=optimist-migrane.txt
{
  "categories": [
    {
      "confidence": 0.95,
      "name": "/Health/Health Conditions/Pain Management"
    }
  ]
}
```
1. Analyze sentiment: `gcloud ml language analyze-sentiment --content='kill all software developers.'` 
``` $ gcloud ml language analyze-sentiment --content='kill all software developers.'
{
  "documentSentiment": {
    "magnitude": 0.8,
    "score": -0.8
  }, ...
```
1. Getting a positive outcome is not so easy. This is a text about handling migraine. This text also contains negative sentences
```
 $ gcloud ml language analyze-sentiment --content-file=optimist-migrane.txt
{
  "documentSentiment": {
    "magnitude": 4.4,
    "score": 0.0
  },
  "language": "en",
  "sentences": [
    {
      "sentiment": {
        "magnitude": 0.4,
        "score": -0.4
      },
      "text": {
        "beginOffset": 0,
        "content": "Anyone who has experienced a migraine knows how debilitating the pain can be."
      }
    },
    {
      "sentiment": {
        "magnitude": 0.8,
        "score": 0.8
      },
      "text": {
        "beginOffset": 78,
        "content": "And while there are several medicinal ways to help alleviate migraine symptoms, a recent study has found that getting enough exercise might just do the trick."
      }
    },
```
1. Sentiment analysis is not enough. 
`gcloud ml language analyze-sentiment --content="Just recently, Joe Biden renewed his pledge to TAKE YOUR LAWFULLY OWNED FIREARMS with the help of his gun-hating friends in Congress."`
```
    {
    "documentSentiment": {
    "magnitude": 0.4,
    "score": 0.4
    }, ...
```




