import os
import re

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

inputs = [
    "Uptown Funk\nLocked out of Heaven\nTalking to the Moon",
]

FILENAME = "songs.txt"


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [TestCase(stdin=[test], attach=[test]) for test in inputs]

    def check(self, reply: str, clue: str):
        if not os.path.exists(FILENAME):
            raise WrongAnswer(f"Cannot find file {FILENAME}")

        if clue[0].rstrip() != reply.rstrip():
            raise WrongAnswer(
                f"Incorrect! 😵❌ Wrong answer!\n"
                f"Your program printed:\n{reply.rstrip()}\n"
                f"\nAre you sure you properly opened {FILENAME}?")

        print(f"Well done! Bruno says thanks!\n\n{inputs[0]}")
        return CheckResult.correct()

if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()
