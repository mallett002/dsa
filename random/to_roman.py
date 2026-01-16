def to_roman_improved(num):
    lookup = [
        ('M', 1000), ('CM', 900), ('D', 500), ('CD', 400), ('C', 100),
        ('XC', 90), ('L', 50), ('XL', 40), ('X', 10), ('IX', 9),
        ('V', 5), ('IV', 4), ('I', 1)
    ]

    romans = []

    for symbol, val in lookup:
        times, num = divmod(num, val)
        romans.append(symbol * times)  # put quo many symbols (multiply strings in python)

    return "".join(romans)


print(to_roman_improved(395))
# print(divmod(25, 2))  # (12, 1)

#     ______
# 25 | 2

# is 12 with remainder of 1

print(divmod(395, 300))  # (12, 1)

# 395 / 1000


# This one was the first attempt:
def to_roman(num):
    lookup = {
        'M': 1000,
        'CM': 900,
        'D': 500,
        'CD': 400,
        'C': 100,
        'XC': 90,
        'L': 50,
        'XL': 40,
        'X': 10,
        'IX': 9,
        'V': 5,
        'IV': 4,
        'I': 1
    }

    romans = []

    for symbol, value in lookup.items():
        while num >= value:
            romans.append(symbol)
            num -= value

    return "".join(romans)


res = to_roman(395)
