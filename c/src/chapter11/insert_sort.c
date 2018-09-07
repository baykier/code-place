//插入排序算法
# include <stdio.h>

int arr[10]; //最大10个
int len = 0;
int insert_sort(int n)
{
    arr[len] = n;
    len++;
    int j;
    int tmp;
    for (int i = 1; i < len; i++)  {        
        tmp = arr[i];
        j = i - 1;//前一个元素
        while (j >= 0 && arr[j] > tmp) {
            arr[j + 1] = arr[j];
            j--;
        }
        arr[j + 1] = tmp;
    }
    
    //
    for (int k = 0; k < len;k++) {
        printf("%d ", arr[k]);
    }
    printf("\n");
    
    return 0;
}

int main(int argc, char const *argv[])
{
    int num;
    while(len <= 10) {        
        printf("请输入要插入的数字:\n");
        scanf("%d", &num);
        printf("输入:%d\n", num);
        insert_sort(num);
    }
    return 0;
}
//算法时间复杂度和空间复杂度
