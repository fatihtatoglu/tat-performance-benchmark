import time


def fibonacci(n):
    if n == 0:
        return 0
    if n == 1:
        return 1

    return fibonacci(n - 1) + fibonacci(n - 2)


limit = range(51)

for i in limit:
    result = fibonacci(i)

    epoch = time.time()

    print(i, ":", result, "(", epoch, ")")
