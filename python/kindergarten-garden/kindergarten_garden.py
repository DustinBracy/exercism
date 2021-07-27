import re

with open("./README.md", "rb") as f:
    STUDENTS = str(f.read())
    STUDENTS = re.findall(r"Alice.*Larry", STUDENTS)[0]
    STUDENTS = re.sub(r"\\n-", "", STUDENTS)
    STUDENTS = sorted(re.sub(r"and ", "", STUDENTS).split(", "))

PLANT_DICT = {"V": "Violets", "C": "Clover", "R": "Radishes", "G": "Grass"}


class Garden:
    def __init__(self, diagram: str, students: str = STUDENTS):
        self.students = sorted(students)
        self.diagram = diagram
        self.student_plant_dict = dict(zip(self.students, self._student_plants()))

    def _student_plants(self) -> list:
        """Returns a list of nested lists of plants sorted by student"""
        row1, row2 = self.diagram.split("\n")
        student_plants = [
            row1[i : i + 2] + row2[i : i + 2] for i in range(0, len(row1), 2)
        ]
        return [[PLANT_DICT.get(key) for key in plant] for plant in student_plants]

    def plants(self, student: str) -> list:
        """Returns plants for a given student."""
        return self.student_plant_dict.get(student)
