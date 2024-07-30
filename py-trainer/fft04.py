#!/usr/bin/env python3
import numpy as np

from config import *


def fft(nparray):
    y = nparray
    n = len(y)  # Number of samples
    freqs = np.fft.fftfreq(n, 1 / SAMPLE_RATE)
    freqs = freqs[:n // 2]  # Keep only first half (positive frequencies)

    Y = abs(np.fft.fft(y))  # Only the magnitude, not the phase
    Y = Y[:n // 2]
    return freqs, Y


def get_waves(f, Y, ranges):
    ret = []
    for low, high in ranges:
        low_idx = np.argmax(f >= low)
        high_idx = np.argmax(f >= high)
        if low_idx == 0 or high_idx == 0:
            ret.append(0)
            continue
        ret.append(np.mean(Y[low_idx:high_idx]))
    return ret


if __name__ == '__main__':
    # The more samples, the narrower (more precise) the frequency bins
    ret = fft(np.array([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]))
    print(ret)
    #print(get_waves(*ret, [(4, 8), (8, 13), (13, 30)]))
    print(get_waves(*ret, [(13, 30), (30, 75)]))
