def is_pangram(sentence):
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    return set(sentence.lower()).issuperset(alphabet)
