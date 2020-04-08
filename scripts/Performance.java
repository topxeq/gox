public class Performance {

    public static double calPi(int pointCount) {
        int inCircleCount = 0;

        double x, y;
        double Pi;

        for (int i = 0; i < pointCount; i++) {
            x = Math.random();
            y = Math.random();

            if (x * x + y * y < 1) {
                inCircleCount++;
            }
        }

        Pi = (4.0 * inCircleCount) / pointCount;

        return Pi;
    }

    public static double cal(double a, double b, double c) {
        return a * b / (c + 1) + a;
    }

    public static long fibonacci(long c) {
        if (c < 2)
            return c;
        return fibonacci(c - 2) + fibonacci(c - 1);
    }

    public static long fibonacciFlat(long c) {
        if (c < 2) {
            return c;
        }

        long fibo = 1;
        long fiboPrev = 1;
        for (long i = 2; i < c; i++) {
            long temp = fibo;
            fibo += fiboPrev;
            fiboPrev = temp;
        }

        return fibo;
    }

    public static void main(String[] args) {

        System.out.println("Test 1");

        long startTime = System.currentTimeMillis();

        double result = 0.0;

        for (double i = 0.0; i < 1000000000; i = i + 1) {
            result += i * i;
        }

        long endTime = System.currentTimeMillis();

        float duration = (float) (endTime - startTime) / 1000;

        System.out.println("Result: " + result);
        System.out.println("Duration: " + duration + " s");

        System.out.println("Test 2");

        startTime = System.currentTimeMillis();

        long resultInt = fibonacciFlat(10000000000L);

        endTime = System.currentTimeMillis();

        duration = (float) (endTime - startTime) / 1000;

        System.out.println("Result: " + resultInt);
        System.out.println("Duration: " + duration + " s");

        System.out.println("Test 2r");

        startTime = System.currentTimeMillis();

        resultInt = fibonacci(50);

        endTime = System.currentTimeMillis();

        duration = (float) (endTime - startTime) / 1000;

        System.out.println("Result: " + resultInt);
        System.out.println("Duration: " + duration + " s");

        System.out.println("Test 3");

        startTime = System.currentTimeMillis();

        result = 0.0;

        for (double i = 0; i < 100000000; i = i + 1) {
            result += Math.random();
        }

        endTime = System.currentTimeMillis();

        duration = (float) (endTime - startTime) / 1000;

        System.out.println("Result: " + result);
        System.out.println("Duration: " + duration + " s");

        System.out.println("Test 4");

        startTime = System.currentTimeMillis();

        result = calPi(100000000);

        endTime = System.currentTimeMillis();

        duration = (float) (endTime - startTime) / 1000;

        System.out.println("Result: " + result);
        System.out.println("Duration: " + duration + " s");

    }
}