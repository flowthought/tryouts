#include <iostream>
#include <cmath>
using namespace std;

int main()
{
    float x = nan("");
    cout << "x = " << x << " Comparison = " << ((x < 1.0f) || (x > 10.0f)) << endl;
    return 0;
}
