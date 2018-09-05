//生成随机数
// gcc -E rand.c 可以查看预编译之后的代码
# include <stdio.h>
# include <stdlib.h>
# define N 10 //预编译

int c[N];

void get_rand(void)
{
    for (int i = 0; i < N;i++)
    {
        c[i] = rand() % 10;
    }
}


void show_rand(void)
{
    for (int i = 0; i < N;i++)
    {
        printf("c[%d] is %d \n", i, c[i]);
    }
}

void get_10_20_rand(int n)
{
    int a[n];
    for (int i = 0; i < n; i++)
    {
        a[i] = 10 + rand() % 10;
        printf("a[%d] is %d \n", i, a[i]);
    }
}

int main(int argc, char const *argv[])
{
    get_rand();
    show_rand();
    get_10_20_rand(20);
    return 0;
}
