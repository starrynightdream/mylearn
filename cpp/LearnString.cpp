// LearnString.cpp
// 关于c++ string 类的使用

#include <iostream>
#include <sstream>
#include <string>
using namespace std;
void init_way()
{
    string s1 = "abcdef";
    string str1("abcdef"); // "abcdef"
    string str2(6, 'a');   // "aaaaaa"
    string str3("0123456789", 3); // "012"
    string str4("0123456789", 3, 6); // "345678"

    cout<<str3.length()<<endl; // size() same
}

void change_way()
{
    string s1 = "abcdef";
    string s2 = "123456";
    // add
    s1 + s2;

    // compare
    (s1 > s2); 
    s1.compare(s2);

    // find
    auto itr = s1.find("cde");

    // insert
    s1.insert(3, s2); // abc123456def
    cout<<s1<<endl;
    s1.insert(6, 5,'X'); // abc123XXXXX456def
    cout<<s1<<endl;

    // swap
    s1.swap(s2);

    // to char*
    const char* cp = s1.c_str();
}

// use string as stream
void ss()
{
    stringstream ss;
    string s("123 456 789");
    ss<<s;
    int a,b,c;
    ss>>a>>b>>c;
    cout<<a<<b<<c<<endl;
}

int main()
{
    ss();
    return 0;
}