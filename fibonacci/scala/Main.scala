object Main {
    def main(args: Array[String]): Unit = {
        for (i <- 0 to 98) {
            val result = fibonacci(i)
            printf("%d: %d\r\n", i, result)
        }
    }

    def fibonacci(n: Int): Long = {
        if (n == 0) {
            return 0
        }

        if (n == 1) {
            return 1
        }

        return fibonacci(n - 1) + fibonacci(n - 2)
    }
}