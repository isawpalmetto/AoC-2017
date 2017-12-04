#!/usr/bin/env python3
import csv

def main():
    chksum = 0
    with open('input.txt', newline='') as f:
        reader = csv.reader(f, delimiter='\t', quoting=csv.QUOTE_NONE)
        for row in reader:
            mini = None
            maxi = None
            for k in row:
                mini, maxi = getMinMax(mini, maxi, int(k))
            chksum += (maxi-mini)
    print(chksum)


def getMinMax(mini, maxi, current):
    if mini is None:
        mini = current
    if maxi is None:
        maxi = current
    if current < mini:
        mini = current
    if current > maxi:
        maxi = current
    return mini, maxi



if __name__ == "__main__":
    main()


