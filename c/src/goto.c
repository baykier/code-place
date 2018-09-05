# include <stdio.h>
//尽量不要使用goto，使用循环替代

int main(void)
{
    goto war;
    war : printf("this is a waring\n");//必须定义在函数内部调用 
}