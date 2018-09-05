#include <stdio.h>

int main(void)
{
    char char1[6] = "hello\0";
    printf("char1 只读 value is %s\n", char1);
    char char2[] = "hhhhahahahhaa";
    printf("chat2 可读写 value is %s \n", char2);
}