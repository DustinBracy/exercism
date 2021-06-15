def slices(series: str, length: int) -> list:
    if length < 1:
        raise ValueError("Length must be greater than 0")
    if len(series) < length:
        raise ValueError("Series does not contain enough characters")
    slices = []

    for i in range(len(series)):
        window = series[i : i + length]
        if len(window) < length:
            break
        slices.append(window)
    return slices
