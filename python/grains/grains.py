"""An exercism solution for:
https://exercism.org/tracks/python/exercises/grains
"""


def square(number: int) -> int:
    """Calculate the number of grains on a chessboard square.

    :param number: the chess square for which grain amounts should be calculated.
    :return: int number of grains on a given square.
    """
    if number > 64 or number <= 0:
        raise ValueError("square must be between 1 and 64")

    return 2 ** (number - 1)


def total() -> int:
    """Calculate the total number of grains on all 64 chessboard squares.
    :return: int total number of grains the king owes.
    """
    return sum(square(x) for x in range(1, 65))
