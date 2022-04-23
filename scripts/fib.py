import datetime
import time

def fib(n):
  if n < 2:
    return n

  return fib(n-1) + fib(n-2) 

startTime = datetime.datetime.now()

print(fib(35))

endTime = datetime.datetime.now()

print("Duration: ", endTime - startTime)

