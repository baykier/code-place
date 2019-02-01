# coding:utf-8
'''
实现adb命令
'''
import os
import platform

def run(command=None):
    """
    执行命令
    libs\\win\\ windows 版本adb
    libs/unix/linux 版本adb
    实现下载好的adb    
    """
    adb_path = os.path.dirname(os.path.abspath(__file__))
    if platform.system() == 'Windows':
        _adb = '{0}\\{1}' . format(adb_path, '\\' . join(('libs', 'win', 'adb.exe')))
    else:
        _adb = '{0}/{1}' . format(adb_path, '/' . join(('libs', 'win', 'adb.exe')))
    print(_adb)
    _command = '{} {}'.format(_adb, command)
    process = os.popen(_command)
    output = process.read()
    return output