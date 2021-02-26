import json
from pathlib import Path


def read_json_index(folder):
    path = Path(folder)
    index = {}
    for f in path.glob("*.json"):
        with open(f) as of:
            index[f.name.split(".")[0]] = json.load(of)
    return index


def read_text_index(folder):
    path = Path(folder)
    index = {}
    for f in path.glob("*.text"):
        with open(f) as of:
            index[f.name.split(".")[0]] = of.read()
    return index


def load_db():
    entity_sentiment = read_json_index("./db/entity-sentiment")
    sentiment = read_json_index("./db/sentiment")
    text = read_text_index("./db/text")

    return entity_sentiment, sentiment, text


def find_matching_article(keyword, score):
    entity_sentiment, sentiment, text = load_db()

    for id, doc in entity_sentiment.items():
        entity = doc["entities"][0]
        if entity["name"] == keyword:
            sentiment_doc = sentiment[id]
            sentiment_score = sentiment_doc["document_sentiment"].get("score", 0.0)
            if score > 0 and sentiment_score < 0:
                return text[id], sentiment_score
            elif score < 0 and sentiment_score > 0:
                return text[id], sentiment_score

if __name__ == "__main__":
    main()