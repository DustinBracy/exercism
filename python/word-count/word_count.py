import re


def count_words(sentence: str) -> dict:
    word_count = {}
    words = [
        word.strip("'") for word in re.sub(r"[^a-z0-9']", " ", sentence.lower()).split()
    ]
    for word in words:
        word_count[word] = word_count.get(word, 0) + 1
    return word_count
