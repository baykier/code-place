# -*- coding: utf-8 -*-
#list
l = [1, 2, 9, 5]
print('list l is ')
print(l)

#索引
print('l index 0 is ')
print(l[0])
# 添加元素
print('l append 89')
l.append(89)
print(l)
# 长度
print('l lenth is')
print(len(l))

# 列表推导
lp = [(x, x * x) for x in range(5)]
print(lp)

for k, v in enumerate(lp):
    print(k,' => ', v)