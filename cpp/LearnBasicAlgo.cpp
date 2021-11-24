// LearnBasicAlgo.cpp
// 一些算法的实现
#include <iostream>
#include <iomanip>
using namespace std;

#include <string.h>
#include "./util.cpp"
using namespace learncode;

namespace other
{
    // 字符串匹配
    // KMP
    int* getnext(const char* prstr, int n)
    {
        int* next = new int[n + 1];
        int k=-1; next[0] = -1;
        for (int i=0; i< n;)
        {
            if (k == -1 || prstr[k] == prstr[i])
            {
                ++k;++i;
                next[i] = k;
            } else k = next[k];
        }
        show(next, n + 1);
        return next;
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
} // namespace learnSort

namespace learnSort
{

    template<class T>
    void insertSort(T* arr, int size)
    {
        int currval = 0;
        for (int i = 1; i < size; ++i)
        {
            currval = arr[i];
            int j = i-1;
            for (;j>=0 && arr[j] >= currval; --j)
            {
                arr[j+1] = arr[j];
            }
            arr[j+1] = currval;
        }
        show(arr, size, '\0');
    }
    
} // namespace learnSort


int main()
{
    char a[] = "1234671231212421212125121212128721212467";
}