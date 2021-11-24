// util.cpp
// useful func

namespace learncode
{
    #include <iostream>
    using namespace std;
    template<class T>
    void show(T* arr, int n, char space = ' ')
    {
        if (space == '\0')
            for (int i =0; i< n;++i)
                cout<< arr[i];
        else 
            for (int i =0; i< n;++i)
                cout<< arr[i]<< space;
        cout<<endl;
    }
} // namespace learncode

