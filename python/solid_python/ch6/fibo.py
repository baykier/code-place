# -*- coding: utf-8 -*-

class fibo:
    name = 'dog'
    def __init__(self, name):
        self.name = name

    def fib(self, n):
        a, b = 0, 1
        while b < n:
            print(b, end=' ')
            a, b = b, a+b
        print()

class ffibo (fibo):

    def fib2(self, n):
        a, b = 0, 1
        while b < n:
            print(b, end=',')
            a, b = b, a+b
        print()

f = ffibo('dog')

f.fib2(8)