//石头剪刀布
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char const *argv[])
{
    char re[][11] = {"石头", "剪刀", "布"};
    int man, computer, ret, con;
    while(1)
    {
        printf("请输入数字: 0 => %s \t 1 => %s \t 2 => %s \n", re[0], re[1], re[2]);
        scanf("%d", &man);
        if(man < 0 || man > 2 )
        {
            printf("非法输入，请输入0，1，2 \n");
            continue;
        }
    }
    return 0;
}
