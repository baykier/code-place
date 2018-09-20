//位操作符
#include <stdio.h>

unsigned n, mask = 0x0000ff00;//mask 掩码

int main(int argc, char const *argv[])
{
    int a, b;
    a = 12;
    b = 29;

    a = a ^ b;
    b = b ^ a;
    a = a ^ b;

    printf("a :%d\n", a);
    printf("b:%d\n", b);
    return 0;
}
