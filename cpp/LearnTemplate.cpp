// LearnTemplate.cpp
// 关于 template 使用的总结
#include <iostream>
using namespace std;

// 模板元编程
// 可以将运行时的计算在编译时完成。

template<int N>
struct Fbnq{
    static long get_val()
    {
        return Fbnq<N-1>::get_val() + Fbnq<N-2>::get_val();
    }
};

template<>
struct Fbnq<2>
{
    static long get_val()
    {
        return 1;
    }
};

template<>
struct Fbnq<1>
{
    static long get_val()
    {
        return 1;
    }
};

template<int N, int X>
struct Pow
{
    enum{ VALUE = N * Pow<N, X - 1>::VALUE };
};

template<int N>
struct Pow<N, 0>
{
    enum{ VALUE = 1 };
};

void test01() {
    cout<< Fbnq<3>::get_val() <<endl;
    cout<<Pow<3, 10>::VALUE<<endl;
}
// 模板元部分

// 特化

// 原模板
template<class T, class N>
class TestingClass
{
    T a;
    N b;
    T doubleA()
    {
        return 2 * a;
    }
public:
    static void test()
    {
        cout<<"other"<<endl;
    }
};

// 偏特化
template<class T>
class TestingClass<T, int>
{
    T a;
    int b;
    T more_it()
    {
        return b * a;
    }
public:
    static void test()
    {
        cout<<"int"<<endl;
    }
};

// 全特化
template<>
class TestingClass<int, int>
{
public:
    static void test()
    {
        cout<<"int , int "<<endl;
    }
};

// 严格来说，指针类型也是类型，也符合普通模板。但是下面的偏特化更严格
// 可以尝试注释查看效果

// 指针偏特化
template<class T, class P>
class TestingClass<T, P*>
{
public:
    static void test()
    {
        cout<<"with pointer"<<endl;
    }
};

// 常量指针偏特化
template<class T, class P>
class TestingClass<T, const P*>
{
public:
    static void test()
    {
        cout<<"with const pointer"<<endl;
    }
};

void test02()
{
    TestingClass<double, char>::test();
    TestingClass<double, int>::test();
    TestingClass<int, int>::test();
    TestingClass<int, int*>::test();
    TestingClass<int, const int *>::test();
}

// 特化部分

int main()
{
    test01();
    return 0;
}