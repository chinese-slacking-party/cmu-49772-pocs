#!/usr/bin/env python3
import pandas as pd
import scipy.io as sio


def loac_SAM40():
    pass


def load_SEED(fname):
    mat_data_seed = sio.loadmat(fname)
    ret = {}
    first = True
    for key in mat_data_seed.keys():
        if '_eeg' not in key:
            continue
        if first:
            print('Subject name:', key[0:key.index('_eeg')])
            first = False
        ret[key] = pd.DataFrame(mat_data_seed[key]).T
    return ret


def load_SEED_labels(fname):
    mat_labels = sio.loadmat(fname)
    return mat_labels['label']


def test_subject1():
    data_dict = load_SEED(r'U:\SEED_EEG\Preprocessed_EEG\1_20131027.mat')
    data_frame = data_dict['djc_eeg1']
    time_stamps = pd.date_range(start='2013-10-27', periods=data_frame.shape[0], freq='5ms')
    print(data_frame)
    print(time_stamps)


def test_subject15():
    data_dict = load_SEED(r'U:\SEED_EEG\Preprocessed_EEG\15_20131105.mat')
    print(data_dict.keys())


if __name__ == '__main__':
    test_subject15()
