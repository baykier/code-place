// 递归倒序打印

#include <stdio.h>
#include <string.h>

char str[] = "abcde";
char arr[] = {'1', '2', '3', '4'};

void back_print(int point)
{
    int size = strlen(str);
    if (point == size) {
        return;
    }
    back_print(point + 1);
    putchar(str[point]);
}

int main(int argc, char const *argv[])
{
    printf("str is : %s \n", str);
    printf("arr is : %s \n", arr);
    back_print(0);
    putchar('\n');
    return 0;
}
