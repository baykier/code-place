//变量
// http://akaedu.github.io/book/expr.variable.html
// 全局变量，只能是常量表达式，全局变量未初始化为0
//局部变量，未初始化是随机值

# include <stdio.h>

char gloabl_var = 'g';
char test_char;

int gloabl_num = 5;
int test_null;

int main()
{
    //声明，先定义再赋值
    char funny;
    char lost;
    int age, num;
    int fc_num;
    char fc_char;

    funny = 'f';
    lost = 'l';
    age = 16;
    num = 25;
    //直接初始化
    int total = 520;
    printf("this is variable chapter!\n");
    printf("char funny:%c\n", funny);
    printf("char lost:%c\n", lost);
    printf("integer age:%d\n", age);
    printf("gloabl variable :%c \n", gloabl_var);
    printf("global variable num:%d \n", gloabl_num);
    //函数
    printf("test un init var test_null :%d \n func_var:%d \n", test_null,fc_num);
    printf("test un init var test_char :%d \n func_char:%d \n", test_char,fc_char);
}