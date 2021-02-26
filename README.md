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

```
 $ gcloud ml language analyze-sentiment --content='kill all homeless people.'
{
  "documentSentiment": {
    "magnitude": 0.6,
    "score": -0.6
  },
  "language": "en",
  "sentences": [
    {
      "sentiment": {
        "magnitude": 0.6,
        "score": -0.6
      },
      "text": {
        "beginOffset": 0,
        "content": "kill all homeless people."
      }
    }
  ]
}
[ 10:22AM ]  [ jan@jans-mbp:~/dev/inno/twistgesprek(master✗) ]
 $ gcloud ml language analyze-sentiment --content='kill all software developers.'
{
  "documentSentiment": {
    "magnitude": 0.8,
    "score": -0.8
  },
  "language": "en",
  "sentences": [
    {
      "sentiment": {
        "magnitude": 0.8,
        "score": -0.8
      },
      "text": {
        "beginOffset": 0,
        "content": "kill all software developers."
      }
    }
  ]
}
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

TODO: 
- find a number of sample documents
- classify-text 

