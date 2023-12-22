def results(item):
    r = 1
    for i in item:
        r *= i
    return r


def part_one(time, dist):
    out = []
    for sp, d in zip(time, dist):
        record = []
        tmp = sp
        for v in range(sp):
            if v * tmp > d:
                record.append(v*tmp)
            tmp -= 1
        print(record)
        out.append(len(record))
    return results(out)

def part_two(time, dist):
    out = 0
    tmp = time
    for v in range(time):
        if v * tmp > dist:
            out += 1
        tmp -= 1
    return out


"""
Time:        62     64     91     90
Distance:   553   1010   1473   1074
"""
time1 = [62, 64, 91, 90]
dist1 = [553, 1010, 1473, 1074]

time2 = 62_649_190
dist2 = 553_101_014_731_074


# print(part_one(time1, dist1))
print(part_two(time2, dist2))

# 8,828,542
# 53,820,647
# 41,382,569