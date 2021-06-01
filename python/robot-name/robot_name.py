import random


class Robot:
    def __init__(self):
        self.name = self._get_name()
        self.used_names = [self.name]

    def _get_name(self):
        def rand_char():
            return chr(random.randrange(65, 91))

        rand_three_digits = str(random.randrange(1000)).zfill(3)
        return rand_char() + rand_char() + rand_three_digits

    def reset(self):
        new_name = self._get_name()
        self.name = new_name if not new_name in self.used_names else self._get_name()
        self.used_names.append(self.name)
