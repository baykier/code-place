//直方图
# include <stdio.h>
# include <stdlib.h>

int c[20];

void get_random(void)
{
    for (int i = 0; i < 20;i++)
    {
        c[i] = rand() % 10;
    }
}

int main(int argc, char const *argv[])
{
    get_random();
    int count[10] = {0};
    //直方图
    for (int i = 0; i < 20;i++)
    {
        count[c[i]]++;
    }
    //第一排
    for (int i = 0; i < 10;i++)
    {
        printf("%d \t", i, count[i]);     
    }
    printf("\n");
    //第二排    
    for (int i = 0; i < 10;i++)
    {
        for (int j = 0; j < 10;j++)
        {
            if(count[j] > i)
            {
                printf("*\t");
            }else
            {
                printf("\t");
            }
        }
        printf("\n");
    }

    return 0;
}
