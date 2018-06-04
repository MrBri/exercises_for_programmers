import unittest
from tip import tip

class TestTip(unittest.TestCase):
    def test_tip_10(self):
        self.assertEqual(tip(10, 15), [1.5, 11.5])

    def test_tip_10(self):
        self.assertEqual(tip(11.25, 15), [1.69, 12.94])
