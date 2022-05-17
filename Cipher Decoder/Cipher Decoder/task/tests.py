from hstest import *
from random import randrange


class TestUserProgram(StageTest):
    @dynamic_test
    def test(self):
        pr = TestedProgram()
        pr.start()

        g, p = randrange(100, 1000), randrange(100, 1000)
        user_answer = f"g={g} and p={p}"

        raw_output = pr.execute(f"g is {g} and p is {p}")

        self.check_empty_or_none_output(raw_output)

        if raw_output.lower().strip() != user_answer:
            raise WrongAnswer(
                f"Your output should be equal to  \"{user_answer}\".\n"
                f"Your output: {raw_output}.")

        return CheckResult.correct()

    @staticmethod
    def check_empty_or_none_output(output):
        if not output:
            raise WrongAnswer("Your output is empty or None.")


if __name__ == '__main__':
    TestUserProgram().run_tests()