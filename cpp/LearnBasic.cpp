#include<iostream>
using namespace std;

// 关于如何传递二维数组
// plan A
int funcList2A(int **a)
{
    // do somethind
}

int funcList2B(int *a)
{
    // i * 3 + j 不用二次导航
}

int main() {
    // test
    int *testlist[3];
    for (int i=0; i< 3; ++i)
        testlist[i] = new int[3];
    funcList2A(testlist);
    int* testhead = &testlist[0][0];
    funcList2B(testhead);
}