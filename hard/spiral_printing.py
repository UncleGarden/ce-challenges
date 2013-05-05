"""print a 2D array (n x m) in spiral order (clockwise)"""
from __future__ import print_function
import fileinput

for line in fileinput.input():
    st = line.split(';')
    s = st[2].split(' ')
    m = [x.strip() for x in s]

    a = int(st[0])
    b = int(st[1])
    c = 0
    la = 0
    lb = 0
    ha = a - 1
    hb = b - 1

    while 1:
        while c % b < hb:
            if c > 0:
                print(" ", end = '')
            print(m[c], end = '')
            m[c] = -1
            c += 1
        if (c + b >= a * b) or (m[c + b] == -1):
            break
        la += 1
        while c / b < ha:
            if c > 0:
                print(" ", end = '')
            print(m[c], end = '')
            m[c] = -1
            c += b
        if (c % b == 0) or (m[c - 1] == -1):
            break
        hb -= 1
        while c % b > lb:
            print('', m[c], end = '')
            m[c] = -1
            c -= 1
        if m[c - b] == -1:
            break
        ha -= 1
        while c / b > la:
            print('', m[c], end = '')
            m[c] = -1
            c -= b
        if m[c + 1] == -1:
            break
        lb += 1

    if c > 0:
        print(" ", end = '')
    print(m[c])
