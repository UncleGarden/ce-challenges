import fileinput
import math

buckets = 7759

sq = [[] for i in range(buckets)]
for i in xrange(46341):
    i2 = i*i
    sq[i % buckets].append(i2)

number = True
for line in fileinput.input():
    if number:
        number = False
        continue

    x = int(line)

    top = int(math.sqrt(x)) + 1
    bot = int(math.sqrt(x/2))

    l = []
    n = 0
    for i in xrange(bot, top):
        t = (x - i*i)
        if t in sq[int(math.sqrt(t)) % buckets] and t not in l:
            l.append(i*i)
            n += 1
    print n
