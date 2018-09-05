//判断莫一年是不是闰年
//编译时加上参数 -lm 
// gcc -lm -o xxx xxx

# include <stdio.h>
# include <math.h>

int is_lead_year(int y)
{
    if((y%4 == 0 && y%100 != 0) || y%400 == 0)
    {
        return 1;
    }else{
        return 0;
    }
}

int my_round(float x)
{
    int r;
    if (x > 0.0)
    {
        r =  ceil(x);
    }else{
        r =  floor(x);
    }
    return r;
}

int main(void)
{
    int iy = is_lead_year(2009);
    printf("year is %d \n",iy);
    int c = my_round(5.69);
    int d = my_round(-6.88);
    printf("%f round is : %d\n",5.69,c);
    printf("%f round is : %d\n",-6.88,c);
    return 0;
}
