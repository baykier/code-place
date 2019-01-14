import glob
import os

print(os.getcwd())
os.chdir('./')
fs = glob.glob('*.*')

print(fs)

