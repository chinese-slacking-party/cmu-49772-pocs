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


GENEVA_SEQ = [
    'Fp1',    'AF3',    'F3',    'F7',    'FC5',    'FC1',    'C3',    'T7',
    'CP5',    'CP1',    'P3',    'P7',    'PO3',    'O1',    'Oz',    'Pz',
    'Fp2',    'AF4',    'Fz',    'F4',    'F8',    'FC6',    'FC2',    'Cz',
    'C4',    'T8',    'CP6',    'CP2',    'P4',    'P8',    'PO4',    'O2',
]


GENEVA_TO_SEED = [
    0,    3,    7,    5,    15,    17,    25,    23,
    33,    35,    43,    41,    52,    58,    59,    45,
    2,    4,    9,    11,    13,    21,    19,    27,
    29,    31,    39,    37,    47,    49,    54,    60,
]


if __name__ == '__main__':
    geneva = read_DEAP_coords('loc2d.csv')
    seed_coords = {}
    for i in range(len(GENEVA_SEQ)):
        seed_coords[GENEVA_TO_SEED[i]] = geneva[i]
    for i in sorted(seed_coords):
        print(i, seed_coords[i])
