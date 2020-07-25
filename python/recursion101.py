# JavaScript episode N09
# challenge N14 = > (Recursion 101) - 7Kyu

# By ALGORITHM ACADEMY

# URL: https://www.codewars.com/kata/5b752a42b11814b09c00005d

def solve(a, b):
  if a == 0 or b == 0:
    return [a,b]
  else:
    if (a >= 2 * b):
      a -= 2 * b
      return solve(a, b)
    elif (b >= 2 * a):
      b -= 2 * a
      return solve(a, b)
    else:
      return [a, b]
    

# better performance (not my solution)
def solve_better_performance(a, b):
  if not (a and b): return [a, b]
  if a >= 2*b: return solve_better_performance(a % (2*b), b)
  if b >= 2*a: return solve_better_performance(a, b % (2*a))
  return [a, b]

print(solve(6,19)) # [6,7]
print(solve(2,1)) # [0,1]
print(solve(22, 5))  # [0,1]
print(solve(2,10)) # [2,2]
