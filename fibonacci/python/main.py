def fibonacci(n):
    if n == 0:
        return 0
    if n == 1:
        return 1

    return fibonacci(n - 1) + fibonacci(n - 2)


limit = range(101)

for i in limit:
    result = fibonacci(i)
    print(i, ":", result)
