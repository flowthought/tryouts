#include <iostream>
using namespace std;
void main()
{
    int x = 0.0f / 0.0f;
    cout << "x = " << x << " Comparison = " << ((x < 1.0f) || (x > 10.0f)) << endl;
}