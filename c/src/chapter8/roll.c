//石头剪刀布
//计算是否胜利，算法
//剪刀石头布，分别用0，1，2标识，前后相减后加上模3 除以3取余，等于1则前面的胜，等于0则平局，等于2则后面的胜利
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char const *argv[])
{
    //c语言中字符串是数组
    char re[][10] = {"剪刀", "石头", "布"};
    int man, computer, ret, con;
    srand(time(NULL));    
    while(1)
    {
        printf("请输入数字: 0 => %s \t 1 => %s \t 2 => %s \n", re[0], re[1], re[2]);
        scanf("%d", &man);
        if(man < 0 || man > 2 )
        {            
            continue;
        }
        computer = rand() % 3;
        con = (man - computer + 3) % 3 ;
        printf("你出的是%d (%s) 电脑出的是%d (%s)\n", man, re[man], computer, re[computer]);
        if(con == 1) {
            printf("你赢了！\n");
        }else if (con == 0) {
            printf("平局!\n");
        }else {
            printf("电脑赢了\n");
        }
    }
    return 0;
}
