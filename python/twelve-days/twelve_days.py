# fmt: off
gifts = "twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, \
nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, \
six Geese-a-Laying, five Gold Rings, four Calling Birds, \
three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."\
    .split(", ")[::-1] 

days = "first, second, third, fourth, fifth, sixth, \
seventh, eighth, ninth, tenth, eleventh, twelfth"\
    .split(", ") 
# fmt: on


def recite(start_verse: int, end_verse: int) -> list:
    recital = []
    for i in range(start_verse, end_verse + 1):
        gift_slice = gifts[:i] if i != 1 else ["a Partridge in a Pear Tree."]
        recital.append(
            f"On the {days[i-1]} day of Christmas my true love gave to me: {', '.join(gift_slice[::-1])}"
        )
    return recital
