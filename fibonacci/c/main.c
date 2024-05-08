#include <stdio.h>

unsigned long fibonacci(int n);

int main()
{
    for (int i = 0; i <= 100; i++)
    {
        unsigned long result = fibonacci(i);
        printf("%d: %ld\r\n", i, result);
    }

    return 0;
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