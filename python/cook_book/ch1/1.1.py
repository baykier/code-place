# coding: utf-8

# 解压序列给单个变量

p = (2,3)

x, y = p

print(x, y); # out 2, 3

m, _ = p

print(m) # out 2

# 解压可迭代

o = [4,6,2,23,2131,3,23,7]

a, *b, c = o

print(a)
print(b) # out list [6, 2, 23, 2131, 3, 23]
print(c)