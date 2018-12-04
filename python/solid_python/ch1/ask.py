# -*- coding:utf-8 -*-

def ask(promt, n = 3, msg = 'yes or no ?'):
    while True:
        anwser = input(promt)
        if anwser in ('y','ye','yes','ok'):
            return True
        if anwser in ('n','no','not'):
            return False        
        n = n - 1
        if n < 1:
            raise  OSError('未知的输入')
        print(msg)


o = ask('你是傻瓜吗？')

if o :
    print('你是傻瓜')
else :
    print('你不是傻瓜')
