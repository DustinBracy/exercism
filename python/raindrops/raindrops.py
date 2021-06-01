def convert(number):
    if type(number) != int:
        return str(number)
    text = ""
    if number % 3 == 0:
        text += "Pling"
    if number % 5 == 0:
        text += "Plang"
    if number % 7 == 0:
        text += "Plong"
    if text == "":
        return str(number)
    return text
