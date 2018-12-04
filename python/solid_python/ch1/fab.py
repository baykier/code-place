# -*- coding:utf-8 -*-
# function

def fab(n):
    """打印一个fab数列"""
    a, b = 0, 1
    while a < n:
        print(a, end=' ')
        a, b = b, a+b
    print()

def fab2(n):
    """返回给定的fab数列"""
    result = []
    a, b  =0, 1
    while a < n:
        result.append(a)
        a, b = b, a+b
    return result

fab(10)
f = fab2(9)
print(f)