import re

with open("./README.md", "rb") as f:
    students = str(f.read())
    students = re.findall(r"Alice.*Larry", students)[0]
    students = re.sub(r"\\n-", "", students)
    students = sorted(re.sub(r"and ", "", students).split(", "))


class Garden:
    def __init__(self, diagram: str, students: str = students):
        self.students = sorted(students)
        self.plant_dict = {"V": "Violets", "C": "Clover", "R": "Radishes", "G": "Grass"}
        self.diagram = diagram

    def _student_plants(self) -> list:
        """Returns a list of nested lists of plants sorted by student"""
        row1, row2 = self.diagram.split("\n")
        student_plants = [
            row1[i : i + 2] + row2[i : i + 2] for i in range(0, len(row1), 2)
        ]
        return [[self.plant_dict.get(key) for key in plant] for plant in student_plants]

    def plants(self, student: str) -> list:
        """Zips sorted plant list and sorted student list, then returns plants for a given student."""
        student_plant_dict = dict(zip(self.students, self._student_plants()))
        return student_plant_dict.get(student)
