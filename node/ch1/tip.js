tip = (amount, rate) => {
	let tip, total;
	tip = amount * (rate / 100); // calculate
	tip = parseFloat(tip.toFixed(2)); // round
	total = amount + tip;

	return [tip, total];
};

module.exports = tip
