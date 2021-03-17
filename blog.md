Social media has been blamed for locking people in a bubble, showing them news that is in line with their beliefs. 
This divides society into different groups that have almost nothing in common anymore. People read what they think they want to read, never seeing a different opinion. 
At the same time governments and influencers have started to call for filtering. Facebook would have to filter out lies and fake news, so we all see the truth only. 
The problem with the filter approach is that opinions will drift toward some bottom line truisms we can all agree on. 
If we start fining social media for violations, the companies will get more and more conservative, and we'll end up in a boring world. Like having a perpetually overcast sky and an eternal drizzle. Grey goo everywhere.

This is not what we need. What we need is to be confronted with opinions that differ from what we think is right. 
This project was inspired by a feature of my favorite Dutch newspaper, NRC (nrc.nl). The feature is called 'Twistgesprek'. The format is that two 
people discuss a statement during the week. Their conversation is summarized and published in the Saturday paper as a back-and-forth of messages. 
Quite often I start with a strong opinion about the subject being discussed, but end up with a more thorough understanding of its nuances because of the discussion.
Having your convictions challenged and modified is a wonderful gift.

So, the idea was to
```
Show people ideas that directly contradict each other
```

With Google ML we can find the sentiment and the subject of a text and then select another text about the same subject but with a different sentiment.
So if someone sees a post that is really positive about X, we want to show a post about X that is really negative.
And vice versa.

The idea we came up with is to use Google ML to summarize a post and have it derive two things:
- the main subject of the document
- the document 'sentiment', which is a number between -1 and 1, where -1 is very negative and +1 is very positive

The main subject of the document can be found with Google ML's `qualify-text` command. Using a message from TheOptimist (https://www.optimistdaily.com/2021/02/just-2-5-hours-of-weekly-exercise-could-help-reduce-migraine-triggers/)
about curing migraines we get this:

```
$ gcloud ml language classify-text --content-file=optimist-migraine.txt

{
 "categories": [
   {
     "confidence": 0.95,
     "name": "/Health/Health Conditions/Pain Management"
   }
 ]
}
```

`Health/Health Conditions/Pain Management`: spot on, awesome. Though maybe this qualification might be too limited for our purposes. The number of categories would probably be too small. 
So we tried an alternative: find the salience (defined as `a striking point or feature` by m-w.com) of a text.

To find the salience, Google offers `analyze-entity-sentiment`. 

```
 $ gcloud ml language analyze-entity-sentiment --content-file=optimist-migraine.txt
{
  "entities": [
    {
      "mentions": [
         ...
      "metadata": {},
      "name": "migraine",
      "salience": 0.097846925,
      "sentiment": {
        "magnitude": 0.3,
        "score": -0.3

```

This shows the concepts the document is about, ordered by salience. We just used the name of the first row in the result, `migraine` in this case.

To get the sentiment, Google offers a tool called `analyze-sentiment`, like this:

```
gcloud ml language analyze-sentiment --content='we love all software developers.'

{
  "documentSentiment": {
    "magnitude": 0.9,
    "score": 0.9
  },...
}
```

The score is 0.9 so this is a positive statement. 

Using the same message from TheOptimist about curing migraines we get this: 

```
 $ gcloud ml language analyze-sentiment --content-file=optimist-migraine.txt
{
  "documentSentiment": {
    "magnitude": 4.4,
    "score": 0.0
  },
```

Which is not very positive, more neutral. This is probably caused by the first sentences of this article, talking about pain and suffering.

Next we tried a message from NRA on Facebook. We expected this to qualify as negative:

```
$ gcloud ml language analyze-sentiment --content="Just recently, Joe Biden renewed his pledge to TAKE YOUR LAWFULLY OWNED FIREARMS with the help of his gun-hating friends in Congress. Help us FIGHT FOR FREEDOM"

{
"documentSentiment": {
"magnitude": 0.5,
"score": 0.2
...
```
I still think this is counter-intuitive, though I can imagine that text like 'Help us fight for freedom' might be considered positive?

So we now have a document about `migraine` with a neutral sentiment score. To make our idea work we would need more documents, some positive and some with a negative sentiment score.

The general idea is this:
1. Suppose a database with text, characterized with
  - classification to find out general area of interest
  - salience/subject to get detailed qualification
  - sentiment
2. Then, given a new text we could use its classification/subject/sentiment to find documents with a mirror sentiment score

This would allow us (or Facebook, to improve the impact of this idea) to whenever a person sees a message about X, calculate its sentiment, and present the reader a different document that has an opposite sentiment value. 
So, Mark, go ahead and reuse our code. You can find it here: https://github.com/jvermeir/twistgesprek