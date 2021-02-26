# analyse new text
# find alternative in database/database.json


gcloud ml language analyze-entity-sentiment --content-file="myfile"
    select: entities[0].name > subject

gcloud ml language analyze-sentiment --content-file="myfile"
