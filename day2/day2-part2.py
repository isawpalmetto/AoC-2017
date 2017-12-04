#!/usr/bin/env python3
import csv

def main():
    chksum = 0
    with open('input.txt', newline='') as f:
        reader = csv.reader(f, delimiter='\t', quoting=csv.QUOTE_NONE)
        for row in reader:
            found = False
            for i in range(0, len(row)):
                if found:
                    break
                current = int(row[i])
                for j, k in enumerate(row):
                    if i != j and divisible(current, int(k)):
                        found = True
                        if current > int(k):
                            div = current/int(k)
                        else:
                            div = int(k)/current
                        chksum += div
                        break


    print(chksum)

def divisible(a, b):
    if a > b:
        return a % b == 0
    else:
        return b % a == 0


if __name__ == "__main__":
    main()


