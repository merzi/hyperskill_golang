import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

GOPHER = '''|
         ,_---~~~~~----._
  _,,_,*^____      _____``*g*"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \/
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
            GOPHER
'''

inputs = [GOPHER]

FILENAME = "ascii_art.txt"


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [TestCase(stdin=GOPHER, attach=GOPHER)]

    def check(self, reply: str, clue: str):
        if not os.path.exists(FILENAME):
            raise WrongAnswer(f"Cannot find file {FILENAME}")

        if clue.rstrip() != reply.rstrip():
            raise WrongAnswer(
                f"Your program printed:\n{reply}\n"
                f"Expected:\n{clue}\n")

        print(clue)
        return CheckResult.correct()


if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()