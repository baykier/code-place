# include <stdio.h>

//全局结构体
struct golbal_struct {
    int x, y;
};
//返回结构体
struct golbal_struct add_struct(struct golbal_struct z,struct golbal_struct y)
{
    z.x = z.x + y.x;
    z.y = z.y + y.y;
    return z;
}
int main(void)
{
    //定义结构体
    struct complex_struct1 {
        int age;
        char n;
    };
    struct complex_struct1 a1, a2;
     //结构体赋值
    struct golbal_struct n;
    n.x = 5;
    n.y = 6;
    struct golbal_struct m = { 1,6 };
    struct golbal_struct t = { 1,0 };
    //计算
    struct golbal_struct o;
    o = add_struct(m, t);
    printf("o.x is %d \t o.y is %d \n", o.x, o.y);
    struct golbal_struct k = add_struct(o, t);
    printf("k.x is %d \t k.y is %d \n", k.x, k.y);

}