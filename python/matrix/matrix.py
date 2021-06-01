class Matrix:
    def __init__(self, matrix_string: str) -> list:
        self.matrix = [
            [int(x) for x in row.split()] for row in matrix_string.splitlines()
        ]

    def row(self, index: int) -> list:
        return self.matrix[index - 1]

    def column(self, index: int) -> list:
        return [row[index - 1] for row in self.matrix]
