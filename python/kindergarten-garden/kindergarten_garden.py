#%%
import re

with open("./README.md", "rb") as f:
    student_list = re.findall(r"Alice.*Larry", str(f.read()))[0]
    student_list = re.sub(r"\\n-", "", student_list)
    student_list = re.sub(r" and", "", student_list).split(", ").sort()


class Garden:
    def __init__(self, diagram, students=student_list):
        self.students = sorted(students)
        self.plant_dict = {"V": "Violets", "C": "Clover", "R": "Radishes", "G": "Grass"}
        self.diagram = diagram

    def _student_plants(self):
        row1, row2 = self.diagram.split("\n")
        student_plants = [
            row1[i : i + 2] + row2[i : i + 2] for i in range(0, len(row1), 2)
        ]
        return [[self.plant_dict.get(c) for c in g] for g in student_plants]

    def plants(self, student):
        response = dict(zip(self.students, self._student_plants()))
        return response.get(student)


# %%
garden = Garden("VCRRGVRG\nRVGCCGCV")

garden.plants("Xander")
# test = garden._student_plants()

# %%
