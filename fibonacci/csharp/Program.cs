namespace Main
{
    using System;

    public static class Program
    {
        public static void Main()
        {
            for (int i = 0; i < 98; i++)
            {
                ulong result = Fibonacci(i);
                Console.Write("{0}: {1}\r\n", i, result);
            }
        }

        private static ulong Fibonacci(int n)
        {
            if (n == 0)
            {
                return 0;
            }

            if (n == 1)
            {
                return 1;
            }

            return Fibonacci(n - 1) + Fibonacci(n - 2);
        }
    }
}