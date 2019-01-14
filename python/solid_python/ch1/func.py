# -*- coding:utf-8 -*-

# def a function
def h(name, age=18):
    print(name, ' hello ', age)

h(name = 'dawang')

# 可变参数
def p(*arg):
    print(type(arg))
    for i in arg:
        print(i, end=' ')
    else:
        print('没有参数输入')

p()

p(1, 5, '324232', '5')

# 关键字参数

def k(**keys):
    print(type(keys))
    for i in keys:
        print(i, '=>', keys[i], end=' ')

k(l = 'm', m = 'o')