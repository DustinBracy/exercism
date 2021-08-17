class School:
    def __init__(self):
        self.students = {}

    def add_student(self, name, grade):
        try:
            self.students[grade] = sorted(self.students[grade] + [name])
        except KeyError:
            self.students[grade] = [name]

    def roster(self):
        student_list = []
        [student_list.extend(self.grade(k)) for k in sorted(self.students.keys())]
        return student_list

    def grade(self, grade_number):
        return self.students.get(grade_number, [])
