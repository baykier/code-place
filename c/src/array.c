# include <stdio.h>




int main(int argc, char const *argv[])
{
    int array[7];
    int count[5] = {1, 5, 7};

    for (int i = 0; i < 5;i++)
    {
        printf("k is %d v is %d \n", i, count[i]);
    }
    return 0;
}
