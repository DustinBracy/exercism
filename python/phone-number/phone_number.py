class PhoneNumber:
    def __init__(self, number: str):
        number = "".join([num for num in number if num.isnumeric()])
        self.area_code = number[:3] if number[0] != "1" else number[1:4]
        self.phone = number[3:] if number[0] != "1" else number[4:]
        if len(self.phone) != 7:
            raise ValueError("Invalid phone number: not enough or too many digits")
        if self.area_code[0] in ["0", "1"] or self.phone[0] in ["0", "1"]:
            raise ValueError("Invalid area or exchange code")
        self.number = self.area_code + self.phone

    def pretty(self) -> str:
        return f"({self.area_code})-{self.phone[:3]}-{self.phone[3:]}"
