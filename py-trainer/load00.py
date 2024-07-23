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


def examine_SEED(fname):
    # What, no metadata of any interest is there.
    # ----------------------------------------
    # PS C:\Users\bspub\go\src\github.com\chinese-slacking-party\cmu-49772-pocs\py-trainer> python .\load00.py
    # __header__
    # b'MATLAB 5.0 MAT-file, Platform: PCWIN64, Created on: Fri Nov 08 10:01:03 2013'
    # __version__
    # 1.0
    # __globals__
    # []
    # Subject name: zjy
    mat_data_seed = sio.loadmat(fname)
    first = True
    for key in mat_data_seed.keys():
        if '_eeg' in key:
            if first:
                print('Subject name:', key[0:key.index('_eeg')])
                first = False
            continue  # This time, we want to see the other keys (metadata)
        print(key)
        print(mat_data_seed[key])


def examine_SEED_raw(fname):
    # ValueError: Unknown mat file type, version 0, 0
    # TODO: How to read .cnt files
    cnt_data_seed = sio.loadmat(fname)
    print(cnt_data_seed.keys())


def test_subject1():
    data_dict = load_SEED(r'U:\SEED_EEG\Preprocessed_EEG\1_20131027.mat')
    data_frame = data_dict['djc_eeg1']
    time_stamps = pd.date_range(start='2013-10-27', periods=data_frame.shape[0], freq='5ms')
    '''
    According to the Excel file, channel order is:
    FP1 FPZ FP2 AF3 AF4 F7 F5 F3
    F1 FZ F2 F4 F6 F8 FT7 FC5
    FC3 FC1 FCZ FC2 FC4 FC6 FT8 T7
    C5 C3 C1 CZ C2 C4 C6 T8
    TP7 CP5 CP3 CP1 CPZ CP2 CP4 CP6
    TP8 P7 P5 P3 P1 PZ P2 P4
    P6 P8 PO7 PO5 PO3 POZ PO4 PO6
    PO8 CB1 O1 OZ O2 CB2
    '''
    print(data_frame)


def test_subject15():
    data_dict = load_SEED(r'U:\SEED_EEG\Preprocessed_EEG\15_20131105.mat')
    print(data_dict.keys())


if __name__ == '__main__':
    pass
