//测试字符串反转
#include <stdio.h>

int main(void)
{
    char str[6] = "hello";
    char r_str[6] = "";
    for (int i = 0; i <= 4;i++) {
        r_str[i] = str[4-i];
    }
    printf("原始字符串:%s \n", str);
    printf("反转字符串:%s\n", r_str);
}