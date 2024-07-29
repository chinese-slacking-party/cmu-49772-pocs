#!/usr/bin/env python3
import numpy as np

from config import *


# Adapted from https://github.com/phamhoai366/EEG-signal-analysis-for-stress-studies/blob/main/SEED_IV.ipynb
# Will become useful if we implement overlapping windows
def make_steps(samples, frame_duration, overlap):
    '''
    in:
    samples - number of samples in the session
    frame_duration - frame duration in seconds
    overlap - float fraction of frame to overlap in range (0,1)

    out: list of tuple ranges
    '''

    Fs = SAMPLE_RATE
    i = 0
    intervals = []
    samples_per_frame = Fs * frame_duration
    while i+samples_per_frame <= samples:
        intervals.append((i,i+samples_per_frame))
        i = i + samples_per_frame - int(samples_per_frame*overlap)
    return intervals


if __name__ == '__main__':
    print(make_steps(1000, 1, 0.5))  # 1000 samples, 1 second frame, 50% overlap
