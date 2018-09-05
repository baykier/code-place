# include <stdio.h>

/**
 * 
 * 9以内乘法口诀
 * */
int main(void)
{
    for (int i = 1; i < 10;i++)
    {
        for (int j = 1; j <= i;j++)
        {
            printf("%d\t", i * j);
        }
        printf("\n");
    }
}