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


# Adapted from https://github.com/mkfzdmr/Deep-Learning-based-Emotion-Recognition
# See https://www.eecs.qmul.ac.uk/mmv/datasets/deap/readme.html "Geneva" for channel mapping
# We should use this because its values are adapted to the fact that our Azimuth projection assumes
# a pi/2 (1.571) radius
def read_DEAP_coords(fname):
    results = []
    with open("loc2d.csv") as csvfile:
        reader = csv.reader(csvfile, quoting=csv.QUOTE_NONNUMERIC)  # change contents to floats
        for row in reader:  # each row is a list
            results.append(np.array(row))
            # print(row)
    return np.array(results)


if __name__ == '__main__':
    print(read_DEAP_coords('loc2d.csv'))
