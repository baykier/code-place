# coding:utf-8
# package disbanker
# file setup.py
from setuptools import setup, find_packages
 
'''
打包

'''

setup(
    name = 'superadb',         # 应用名
    version = '1.0',        # 版本号
    keywords = ["adb", "windows",'linux'],  
    description = "project", 
    license = "MIT Licence", 
    url = "https://github.com/baykier/code-place/python/superadb",  
    author = "baykier",  
    author_email = "1035666345@qq.com",
    packages =  find_packages(exclude=["*.tests", "*.tests.*", "tests.*", "tests"]), # 包括在安装包内的Python包
    Platform = 'linux,windows'
)