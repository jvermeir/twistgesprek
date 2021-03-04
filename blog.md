Social media has been blamed for locking people in a bubble, showing them news that is in line with their believes. 
This divides society into different groups that have almost nothing in common anymore. People read what they think they want to read, never seeing a different opinion. 
At the same time governments and influencers have started to call for filtering. Facebook would have to filter out lies and fake news so we all see the truth only. 
The problem with the filter approach is that opinions will drift toward some bottom line believes we can all agree on. 
If we start fining social media for violations, the companies will get more and more conservative and we'll end up in a boring world. Like having a perpetually overcast sky and an eternal drizzle.

This is not what we need. What we need is to be confronted with opinions that differ from what we think is right. 
This project was inspired by a feature of my favorite Dutch newspaper, NRC (nrc.nl). The feature is called 'Twistgesprek'. The format is that two 
people discuss a statement during the week. Their conversation is summarized and published in the Saterday paper as a back-and-forth of messages. 
Quite often I start with a strong opinion about the subject being discussed, but end up with a more thorough understanding of the subject and its nuances because of the statements made.
Having your convictions challenged and modified is a wonderful gift.

So we started out with the ambition to present people on their Facebook timeline with the posts they would normally see, interspersed with posts that would claim the contrary. 
A post claiming statement-A would be followed by a post claiming statement-not-A. This turned out to be difficult. But we found a strategy that might work too. 
The idea we came up with is to use Google ML to summarize a post and have it derive two things:
- the document 'sentiment', which is a number between -1 and 1, where -1 is very negative and +1 is very positive
- the main subject of the document

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

Using a message from TheOptimist (https://www.optimistdaily.com/2021/02/just-2-5-hours-of-weekly-exercise-could-help-reduce-migraine-triggers/) 
about curing migraines we get this: 

```
 $ gcloud ml language analyze-sentiment --content-file=optimist-migraine.txt
{
  "documentSentiment": {
    "magnitude": 4.4,
    "score": 0.0
  },
```
Which is not very positive, more neutral. This is probably caused by the first sentences of this article, talking about pain and suffering.

To find out what a document is about, Google offers `analyze-entity-sentiment`, like this:

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

This shows the concepts the document is about, ordered by salience (defined as `a striking point or feature` by m-w.com). We just used the name of the first row in the result, `migraine` in this case. 
So we now have a document about `migraine` with a neutral sentiment score. To make our idea work we would need more documents, some positive and some with a negative sentiment score.
These documents and their subject would go into a database and then whenever a person sees a message about migraine, the sentiment is calculated and the reader is 
presented with a different document that has an opposite sentiment value. 

This is still work in progress, but we had a useful day learning a lot about the power of automated document analysis.  