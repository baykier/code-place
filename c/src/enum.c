//枚举类型
# include <stdio.h>

//枚举类型 默认为int ，从0 开始累加
enum global_e
{
    FIRTT,
    SECOND
};

int main(int argc, char const *argv[])
{
    printf("global_e is enum type and FIRST is %d \t SECOND is %d \n", FIRTT, SECOND);
    printf("参数个数 argc is %d \n", argc);
    return 0;
}
