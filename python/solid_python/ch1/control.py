# -*- coding: utf-8 -*-
#while
a, b = 0, 1
print(a)
print(b)

print('fab is ')
while a < 10:
    print(b)
    a, b = b, a + b
    
# if
x = int(input('输入一个数字:'))
print('输入数字为:', x)
if x > 100 :
    print('输入的数字大于100')
elif  x == 100 :
        print('输入的数字等于100')
else:
    print('输入的数字小于100')

# for

l = [1, 9, 39, 45, 33]

for m in l:
    print(m, m ** 2)