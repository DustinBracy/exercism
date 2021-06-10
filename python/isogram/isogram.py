alphabet = 'abcdefghijklmnopqrstuvwxyz'

def is_isogram(string:str) -> bool:
    used_letters = []
    for letter in string.lower():
        if letter not in alphabet:
            continue
        if letter in used_letters:
            return False
        used_letters.append(letter)
    return True
    