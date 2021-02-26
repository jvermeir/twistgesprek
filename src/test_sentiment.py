import sentiment

def test_find_matching_article():
    keyword = "migraine"
    score = 0.3

    matching_text, score = sentiment.find_matching_article(keyword, score)

    # expect a file with a negative score
    assert score == -0.3
    # assert matching_text == ""