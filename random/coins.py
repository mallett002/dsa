# coins problem, same problem basically as the to_roman
# [qs,ds,ns,ps]: [number,number,number,number]
def coins(num):
    # symbol, value, index
    lookup = [
        (25, 0),
        (10, 1),
        (5, 2),
        (1, 3),
    ]

    res = [0, 0, 0, 0]

    for value, index in lookup:
        times, num = divmod(num, value)  # how many quarters fit into the num
        res[index] = times

        if num <= 0:
            return res

    return res


# [2,2,0,2]
print(coins(72))
