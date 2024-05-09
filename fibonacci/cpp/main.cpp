#include <iostream>
#include <chrono>

using namespace std;

unsigned long fibonacci(int n);

int main()
{
    for (int i = 0; i <= 50; i++)
    {
        unsigned long result = fibonacci(i);

        auto now = std::chrono::system_clock::now();
        auto epoch = std::chrono::duration_cast<std::chrono::seconds>(now.time_since_epoch()).count();

        printf("%d: %ld (%ld)\r\n", i, result, epoch);
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