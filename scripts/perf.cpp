#include <stdio.h>
#include <stdlib.h>
#include <time.h>

double calPi(int pointCount)
{
    int inCircleCount = 0;
    double x, y, Pi;

    for (int i = 0; i < pointCount; i++)
    {
        x = rand() / (double)RAND_MAX; // x 将介于[0, 1]之间
        y = rand() / (double)RAND_MAX; // x 将介于[0, 1]之间

        if (x * x + y * y < 1) // 判断随机产生的点是否落在四分之一圆内
        {
            inCircleCount++; //
        }
    }

    Pi = (4.0 * inCircleCount) / pointCount;

    return Pi;
}

double square(double a)
{
    return a * a;
}

long long fibonacci(long long c)
{
    if (c < 2)
    {
        return c;
    }

    return fibonacci(c - 2) + fibonacci(c - 1);
}

int main()
{
    srand((int)time(NULL)); // 初始化随机数

    time_t startTime = time(NULL);

    double result = 0.0;

    for (double i = 0; i < 1000000000; i = i + 1)
    {
        result += i * i;
    }

    time_t endTime = time(NULL);

    printf("Result: %f\n", result);

    printf("耗时: %ld秒\n", (endTime - startTime));

    printf("Test 2\n");

    startTime = time(NULL);

    long long result2 = 0;

    result2 = fibonacci(50);

    endTime = time(NULL);

    printf("Result: %lld\n", result2);

    printf("耗时: %ld秒\n", (endTime - startTime));

    return 0;
}