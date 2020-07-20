# Python episode N08
# challenge N14 = > (Array element parity) - 7Kyu

# By ALGORITHM ACADEMY
# Url: https: // www.codewars.com/kata/5a092d9e46d843b9db000064/python

def solve(arr):
  for n in arr:
    c = arr.count(-n)
    if (c > 0):
      continue
    else:
      return n
  
print(solve([1, -1, 2, -2, 3]))
print(solve([-3, 1, 2, 3, -1, -4, -2]))
