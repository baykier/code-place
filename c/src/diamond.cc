# include <stdio.h>
# include <math.h>
//http://akaedu.github.io/book/ch06s05.html
//打印菱形
void diamond(int n,char s)
{
    if(n %2 == 0 || n ==1 )
    {
        printf("n is :%d must be odd,enven given and larger than 1 \n", n);
        return;
    }
    int mid = (n + 1) / 2;//中间数
    int tag_num;
    for (int i = 1; i <= n; i++)
    {
        tag_num = abs(i - mid);
        for (int m = 0; m < tag_num; m++)
        {
            printf("\t");            
        }
        for (int j = 0; j < ((mid - tag_num) * 2 - 1);j++)
        {
            printf("%c\t", s);
        }
        printf("\n");
    }
}
int main(void)
{
    diamond(7, '*');
}