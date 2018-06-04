const tip = require('./tip');

test('15% tip to $10 to total $11.50', () => {
  expect(tip(10, 15)).toEqual([1.5, 11.5]);
});

test('15% tip to $11.25 to total $12.95', () => {
  expect(tip(11.25, 15)).toEqual([1.69, 12.94]);
});

