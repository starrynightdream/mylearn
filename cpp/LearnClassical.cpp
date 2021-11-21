// LearnClassical.cpp
// 经典问题的解
#include <iostream>
using namespace std;

// 约瑟夫环 （报数退出序）
void Jonseph(int all, int count)
{
    if (all <=0 || count <= 0)
        return; 
    int* people = new int[all];
    for (int i=0; i < all; ++i)
        people[i] = i + 1;

    int outter = 0;
    int idx = 0;
    int speak = 0;
    while (outter < all)
    {
        while (people[idx] == 0) idx++;
        if (++speak == count)
        {
            speak = 0;
            cout<< people[idx] << ' ';
            people[idx] = 0;
            ++outter;
        }

        idx++;

        if (idx >= all) idx = 0;
    }
    cout<<endl;
}

// 模拟积分
double Int(double (*func)(double), double from, double to)
{
    double peer = 2e-10;
    double idx = from;
    double ans = 0;
    while (idx < to)
    {
        ans += peer * func(idx);
        idx += peer;
    }
    return ans;
}
double y(double x)
{
    return 2 * x;
}

// 数组移动
void reverse(int* a, int start, int end)
{
    while(start < end)
    {
        int temp = a[end];
        a[end] = a[start];
        a[start] = temp;
        ++start; --end;
    }
}
void moveback(int* a, int n, int m)
{
    int offset = 0;
    if (m < 0) offset = n;
    m = m%n + offset;
    if (!m) return;
    reverse(a, 0, m-1);
    reverse(a, m, n-1);
    reverse(a, 0, n-1);
}
void show(int *a, int n)
{
    for (int i=0; i < n; ++i)
        cout<< a[i] <<' ';
    cout<<endl;
}

void othermoveback(int* a, int n, int m)
{
    if (n == 0 || m%n==0)
        return;
    int temp = a[n-1];
    for (int i = n-1; i>0;--i)
        a[i] = a[i-1];
    a[0] = temp;
    othermoveback(a, n, m-1);
}

int main()
{
    int a[10] = {0, 1, 2, 3, 4, 5, 6, 7 ,8 ,9};
    show(a, 10);
    othermoveback(a, 10, 3);
    show(a, 10);
}