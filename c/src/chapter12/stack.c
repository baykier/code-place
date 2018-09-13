/**
*
* 堆栈
* LIFO 先进后出
* 主要应用，倒序打印
**/

# include <stdio.h>

char str[512];
int size = 0;

void push (char c) 
{
    str[size++] = c;
}

char pop(void)
{
    return str[--size];
}

int is_empty(void)
{
    return size == 0;
}

int main(int argc, char const *argv[])
{
    push('a');
    push('b');
    push('c');

    while(! is_empty()) {
        putchar(pop());
    }
    putchar('\n');
    return 0;
}
