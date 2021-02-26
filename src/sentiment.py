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


entity_sentiment = read_json_index("./db/entity-sentiment")
sentiment = read_json_index("./db/sentiment")
text = read_text_index("./db/text")

print(len(entity_sentiment), len(sentiment), len(text))