// LearnStrF.cpp
// 字符串算法的实现示例
#include <iostream>
using namespace std;

int m_strlen(const char * ptr)
{
    if (ptr == nullptr)
    {
        return 0;
    }
    int l = 0;
    while (*ptr++ != '\0')
    {
        l++;
    }
    return l;
}

char* m_substr(const char * ptr, int start, int len)
{
    if (len <= 0)
    {
        return nullptr;
    }

    int length = 0;
    while (*(ptr + length) != '\0')
    {
        length++;
    }
    if (start >= length)
    {
        return nullptr;
    }

    int _str_len = (length - start + 1) < len ? (length - start + 1) : len;
    char* cp = new char[_str_len];
    for(int i=0;i < _str_len; ++i)
    {
        cp[i] = ptr[start + i];
    }
    return cp;
}

void m_strcopy_void(char* to, const char* from)
{
    while(*to++ = *from++){}
}

char* m_strcopy (char* to, const char* from)
{
    char* p = to;
    while(*to++ = *from++){}
    return p;
}

char* m_strcat(char* org, char* add)
{
    char* p = org;
    while (*org != '\0')
    {
        org++;
    }
    
    while(*org++ = *add++){}
    return p;
}

int m_strcmp(char* a, char* b)
{
    while(*a && *b && *a == *b)
    {
        a++;b++;
    }
    return *a - *b;
}

int main()
{
    char a[20] = "12345678900";
    char b[20] = "12344567890";

    cout<<m_strcmp(a,b);
}