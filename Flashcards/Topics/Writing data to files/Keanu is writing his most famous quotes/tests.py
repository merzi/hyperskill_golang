import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

inputs = [
    "You're breathtaking!\nWake up Samurai! We have a city to burn.\nLose? I don't lose; I win! I'm a lawyer, "
    "that's my job, that's what I do!"
]

FILENAME = "keanu_quotes.txt"


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

        print(f"Well done! Keanu says thanks!\n\n{attach[0]}")
        return CheckResult.correct()


if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()
