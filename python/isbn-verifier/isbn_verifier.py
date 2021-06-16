def is_valid(isbn: str) -> bool:
    try:
        check_digit = 10 if isbn[-1] == "X" else int(isbn[-1])
        isbn = [int(x) for x in isbn.replace("-", "")[:-1]]
        isbn = isbn + [check_digit]
        if len(isbn) != 10:
            return False
        return sum((num * (10 - i)) for i, num in enumerate(isbn)) % 11 == 0
    except:
        return False
