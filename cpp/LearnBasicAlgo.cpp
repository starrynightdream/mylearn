// LearnBasicAlgo.cpp
// 一些算法的实现
#include <iostream>
using namespace std;

// 字符串匹配
// KMP
int* getnext(const char* prstr, int n)
{
    int* next = new int[n];
    int k=-1; next[0] = -1;
    for (int i=0; i<n; ++i)
    {
        if (k == -1 || prstr[k] == prstr[i])
        {
            ++k;++i;
            next[i] = k;
        } else k = next[k];
    }
}
int KMP(const char* str1, int m, const char* str2, int n)
{
    if (n > m)
        return -1;

    int* next = getnext(str2, n);
    int i =0, j =0;
    while (i < m && j < n)
    {
        if (j == -1 || str1[i] == str2[j])
        {
            ++i;++j;
        }
        else
        {
            j = next[j];
        }
    }
    if (j == n)
    {
        return i - j;
    }
    return -1;
}

int main()
{

}