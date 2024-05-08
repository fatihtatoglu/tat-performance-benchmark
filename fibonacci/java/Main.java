import java.math.BigInteger;

public class Main {

    public static void main(String[] args) {
        for (int i = 0; i < 100; i++) {
            BigInteger result = Fibonacci(i);
            System.out.printf("%d: %d%n", i, result);
        }
    }

    public static BigInteger Fibonacci(int n) {
        if (n == 0) {
            return BigInteger.valueOf(0);
        }

        if (n == 1) {
            return BigInteger.valueOf(1);
        }

        return Fibonacci(n - 1).add(Fibonacci(n - 2));
    }
}