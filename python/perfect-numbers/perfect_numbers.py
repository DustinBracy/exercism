def classify(number: int) -> str:
    if number < 1:
        raise ValueError("Number must be > 0")

    max_factor = int(number / 2) + 1

    aliquot_sum = sum(x for x in range(1, max_factor) if number % x == 0)

    if aliquot_sum == number:
        return "perfect"

    if aliquot_sum > number:
        return "abundant"

    return "deficient"
