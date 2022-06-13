import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

inputs = [
    "Krabby Pattie $2.00\n"
    "Krusty Combo $3.99\n"
    "Krusty Deluxe $3.00\n"
    "Seaweed Salad $1.50\n"
    "Coral Bits $1.95\n"
    "Kelp Shake $2.00"
]

FILENAME = "galley_grub.txt"


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [TestCase(stdin=[test], attach=[test]) for test in inputs]

    def check(self, reply: str, attach: list):
        if not os.path.exists(FILENAME):
            raise WrongAnswer(f"Cannot find file {FILENAME}")

        with open(FILENAME, "r") as f:
            content = f.read().strip()
            if content != attach[0]:
                raise WrongAnswer(
                    f'Invalid content of {FILENAME} file, got:\n{content}\n\nExpected:\n{attach[0]}'
                )

        print("""Welcome to the Krusty Krab! 🍔🥤🍟

Krabby Pattie $2.00
Krusty Combo $3.99
Krusty Deluxe $3.00
Seaweed Salad $1.50
Coral Bits $1.95
Kelp Shake $2.00""")
        return CheckResult.correct()


if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()