import re


def abbreviate(words):
    words = re.sub(r"-", " ", words)
    words = re.sub(r"[^A-Z ]", "", words.upper())
    return "".join(word[0] for word in words.split())
