object Main {
    def main(args: Array[String]): Unit = {
        for (i <- 0 to 50) {
            val result = fibonacci(i)

            val epoch = java.time.Instant.now.toEpochMilli

            printf("%d: %d (%d)\r\n", i, result, epoch)
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