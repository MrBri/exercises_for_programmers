def tip(amount, rate):
    tip = amount * (rate / 100)
    tip = round(tip, 2)
    total = amount + tip
    return [tip, total]
