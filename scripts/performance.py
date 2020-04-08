import random
import datetime
import time


def calPi(pointCount):
    inCircleCount = 0

    x = 0.0
    y = 0.0
    Pi = 0.0

    for i in range(pointCount):
        x = random.random()
        y = random.random()

        if (x * x + y * y) < 1:
            inCircleCount = inCircleCount + 1

    Pi = (4.0 * inCircleCount) / pointCount

    return Pi

def square(a):
    return a * a


random.seed()

startTime = datetime.datetime.now()

result = calPi(10000000)
# result = 0.0

# for i in range(1000000000):
#     result += square(i)

endTime = datetime.datetime.now()

print("Pi: ", result)
print("Duration: ", endTime - startTime)
