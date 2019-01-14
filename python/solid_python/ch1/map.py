# -*- coding: utf-8 -*-

#map

m = {'me': 9, 'you': 9}

print('map:', m)

m['he'] = 78

print(m)

# dict

m1 = dict(n=8,m=9)
m2 = dict([('you', 4), ('me', 6)])

print(m1,m2)

# 字段推导式
m3 = {x: x * x for x in range(6)}

print(m3)

for k, v in m3.items():
    print('k:', k,'=>', 'v:',v)
