#include <iostream>

using namespace std;

unsigned long fibonacci(int n);

int main()
{
    for (int i = 0; i < 98; i++)
    {
        unsigned long result = fibonacci(i);
        printf("%d: %ld\r\n", i, result);
    }
}

unsigned long fibonacci(int n)
{
    if (n == 0)
    {
        return 0;
    }

    if (n == 1)
    {
        return 1;
    }

    return fibonacci(n - 1) + fibonacci(n - 2);
}