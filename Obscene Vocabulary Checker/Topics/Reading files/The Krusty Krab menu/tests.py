import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

inputs = [
    "Krabby Pattie $2.00\nKrusty Combo $3.99\nKrusty Deluxe $3.00\nSeaweed Salad $1.50\nCoral Bits $1.95"
]

FILENAME = "galley_grub.txt"


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [TestCase(stdin=[test], attach=[test]) for test in inputs]

    def check(self, reply: str, attach: list):
        if not os.path.exists(FILENAME):
            raise WrongAnswer(f"Cannot find file {FILENAME}")

        with open(FILENAME, "r", encoding="utf8") as f:
            content = f.read().rstrip()
            if reply.rstrip() != content:
                raise WrongAnswer(
                    f"Incorrect! 😵❌ Wrong answer!\n"
                )

        print(f"\n{reply}")
        print("Welcome to the Krusty Krab! 🍔🥤🍟")
        return CheckResult.correct()


if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()
