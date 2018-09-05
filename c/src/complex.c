//复数计算
# include <stdio.h>

// 定义一个长方形
struct cube {
    double x, y;
};
//定义一个函数，用来生成长方形结构体
struct cube make_cube(double x,double y)
{
    struct cube z = {x, y};
    return z;
};
//计算长方形结构体
double calculate_cube(struct cube m)
{
    return m.x * m.y;
};

int main(void)
{
    struct cube m;
    m = make_cube(9, 8);
    double n = calculate_cube(m);
    printf("cube calculate_cube is %f ,cube.x is %f cube.y is %f \n", n, 9, 8);
    return 0;
}

