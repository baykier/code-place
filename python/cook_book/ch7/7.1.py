# coding: utf-8

#接受任意数量的参数

# *参数后面只接受关键字参数即 m=v 形式，否则报错
def get_arg(*arg, m):
    print(arg, m)

get_arg(1, 23, 4, 55, 34434, 't', m='hhh')

# 一个函数只能有一个*参数 下面会报错
# def get_arg2(*arg, *arg2):
#     print(arg, arg2)

# get_arg2(3,2,2,2,7)

# 任意参数在位置参数之后,在最后一个位置参数后面
def get_arg3(m,n, *arg2):
    print(m, n,arg2)

get_arg3('hhh', 54, 3, 3, 2322, 'lll')


# 位置参数一关键字传参也可以接受
def get_arg4(*,a, b):
    print(a, b)

get_arg4(a=1, b=2)

# get_arg4( b = 4)
