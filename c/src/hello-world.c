#include <stdio.h>

/**
 * 
 * 这是第一个 C 程序
 * 输出hello world
 *  
 * */

int main()
{
    // 单行注释
    //字符串用双引号引起来 'Hello world' 错误的写法 单引号只能用于字符(character)
    //调试用 gcc - Wall hello-world.c -Wall 显示所有警告信息

    printf("Hello world\n");

    // 返回值 用 return 
    return 0;
}