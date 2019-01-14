# -*- coding:utf-8 -*-

def s(a,l = []):
    l.append(a)
    return l
print(s(1)) # [1]
print(s(4)) # [1, 4]
print(s(6)) # [1, 4, 6]

def sn(a, l = None):
    l = []
    l.append(a)
    return l

print(sn(1))
print(sn(2))
print(sn(6))