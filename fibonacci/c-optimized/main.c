#include <stdio.h>
#include <sys/time.h>

unsigned long fibonacci(int n);

int main()
{
    for (int i = 0; i <= 50; i++)
    {
        unsigned long result = fibonacci(i);

        struct timeval currentTime;
        gettimeofday(&currentTime, NULL);

        printf("%d: %ld (%ld)\r\n", i, result, currentTime.tv_sec);
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