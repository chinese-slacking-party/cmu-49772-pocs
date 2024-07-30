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


if __name__ == '__main__':
    # The more samples, the narrower (more precise) the frequency bins
    print(fft(np.array([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])))
