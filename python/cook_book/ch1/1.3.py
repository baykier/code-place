# coding: utf-8



def test_yield(i):
    for m in range(i):
        yield m,1

p = test_yield(5)

print(p)

for x in p:
    print(x)

for m,n in p:
    print(m,n)

a = 1, 2

print(a)


