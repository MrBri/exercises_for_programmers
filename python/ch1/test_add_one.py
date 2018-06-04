import unittest
from add_one import fun

class MyTest(unittest.TestCase):
    def test(self):
        self.assertEqual(fun(3), 4)
