# include <stdio.h>
//获取给定数的个位数和十位数

void get_num(int num)
{
    int p1 = num % 10;
    int p2 = (num / 10) % 10;
    printf("num %d p1:%d p2:%d \n", num, p1, p2);
}

int main(void)
{
    get_num(865);
    return 0;
}
