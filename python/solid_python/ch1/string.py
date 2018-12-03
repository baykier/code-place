# -*- coding: utf-8 -*-

# 单引号
print('this is a single quto string')

# 双引号
print("this is double quto string")

# 多行
print("""
    usage: lalal [-hopq]
    -h                    显示帮助信息
    -H                    hostname                

""")
# 字符串拼接

h = 'hello'
w = ' world'
print(h + w)

# 字符串重复

m = "hello word "
print(m * 3)

# 字符串切片
print(m[:])
print('hey ' + m[:])

# 字符串不可修改
# 返回字符串长度
print('str m lenth is ')
print(len(m))