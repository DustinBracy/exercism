score_dict = {
    **{x: 1 for x in ["A", "E", "I", "O", "U", "L", "N", "R", "S", "T"]},
    **{x: 2 for x in ["D", "G"]},
    **{x: 3 for x in ["B", "C", "M", "P"]},
    **{x: 4 for x in ["F", "H", "V", "W", "Y"]},
    **{"K": 5},
    **{x: 8 for x in ["J", "X"]},
    **{x: 10 for x in ["Q", "Z"]},
}


def score(word: str) -> int:
    return sum(score_dict.get(char, 0) for char in word.upper())
