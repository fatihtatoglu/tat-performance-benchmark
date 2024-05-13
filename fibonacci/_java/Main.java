import java.math.BigInteger;

public class Main {

    public static void main(String[] args) {
        for (int i = 0; i <= 50; i++) {
            BigInteger result = Fibonacci(i);

            long epoch = System.currentTimeMillis()/1000;

            System.out.printf("%d: %d (%d)\r\n", i, result, epoch);
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