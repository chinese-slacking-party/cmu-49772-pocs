#!/usr/bin/env python3
import csv

import numpy as np


def read_SAM40_coords(fname):
    results = []
    mapping = {}
    with open(fname) as csvfile:
        reader = csv.reader(csvfile, delimiter='\t')
        for row in reader:
            results.append(np.array(row[:-1], dtype=float))
            mapping[row[-1].strip()] = len(results) - 1
            #print(row)
    return results, mapping


if __name__ == '__main__':
    print(read_SAM40_coords('locs_32.tsv'))
