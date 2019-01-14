# -*- coding: utf-8 -*-

class fibo:
    def fib(self, n):
        a, b = 0, 1
        while b < n:
            a, b = b, a + b
            print(b, end=' ')
        print()

        