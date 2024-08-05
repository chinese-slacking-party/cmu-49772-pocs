#!/usr/bin/env python3
import glob
import logging
import os.path
import pandas as pd
import random
import scipy.io as sio

from config import *

from frame03 import make_steps


class SEEDLoader:
    def __init__(self, fname):
        self.mat_data_seed = sio.loadmat(fname)
        self.data_dict = {}
        self.count = 0
        first = True
        for key in self.mat_data_seed.keys():
            if '_eeg' not in key:
                continue
            if first:
                self.subject_name = key[0:key.index('_eeg')]
                logging.debug('Subject name:', self.subject_name)
                first = False
            self.data_dict[int(key[key.index('_eeg')+4:]) - 1] = pd.DataFrame(self.mat_data_seed[key]).T
            self.count += 1
    
    def __len__(self):
        return self.count
    
    def __getitem__(self, key):
        return self.data_dict[key]
    
    def __iter__(self):
        return iter(self.data_dict)
    
    def raw_frames(self, key):
        data_frame = self.data_dict[key]
        steps = make_steps(len(data_frame), FRAME_DURATION, OVERLAP)
        ret = []
        for start, end in steps:
            ret.append(data_frame.iloc[start:end])
        return ret


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


SEED_DIR_FRED_LAPTOP = r'C:\Users\bspub\Desktop\Temp\SEED-III\Preprocessed_EEG'
SEED_DIR_CLOUD = r'/content/Preprocessed_EEG'
FILE_SUB01_EXPR1 = os.path.join(SEED_DIR_FRED_LAPTOP, '1_20131027.mat')
FILE_SUB15_EXPR3 = os.path.join(SEED_DIR_CLOUD, '15_20131105.mat')


def test_subject1():
    data_dict = load_SEED(FILE_SUB01_EXPR1)
    data_frame = data_dict['djc_eeg1']
    time_stamps = pd.date_range(start='2013-10-27', periods=data_frame.shape[0], freq='5ms')
    '''
    According to the Excel file, channel order is:
    (1)FP1 FPZ (17)FP2 (2)AF3 (18)AF4 (4)F7 F5 (3)F3
    F1 (19)FZ F2 (20)F4 F6 (21)F8 FT7 (5)FC5
    FC3 (6)FC1 FCZ (23)FC2 FC4 (22)FC6 FT8 (8)T7
    C5 (7)C3 C1 (24)CZ C2 (25)C4 C6 (26)T8
    TP7 (9)CP5 CP3 (10)CP1 CPZ (28)CP2 CP4 (27)CP6
    TP8 (12)P7 P5 (11)P3 P1 (16)PZ P2 (29)P4
    P6 (30)P8 PO7 PO5 (13)PO3 POZ (31)PO4 PO6
    PO8 CB1 (14)O1 (15)OZ (32)O2 CB2
    Numbers in parentheses are 1-based channel mappings to the Geneva device described in the DEAP
    dataset brief.
    '''
    print(data_frame)


def test_subject15():
    ldr = SEEDLoader(FILE_SUB15_EXPR3)
    for i in ldr:
        print()
        print(i)
        rf = ldr.raw_frames(i)
        print(rf, len(rf))


def test_list_like_loader():
    for file in glob.glob(os.path.join(SEED_DIR_FRED_LAPTOP, '*_*.mat')):
        print()
        print(file)
        ldr = SEEDLoader(file)
        testcase = random.randint(0, len(ldr) - 1)
        print(ldr.subject_name, testcase, '/', len(ldr))
        print(testcase, ldr[testcase])


if __name__ == '__main__':
    test_subject15()
