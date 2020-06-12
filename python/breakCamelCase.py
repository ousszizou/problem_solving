# Python episode N05
# challenge N10 => (Break camelCase) - 6Kyu

# By ALGORITHM ACADEMY

def solution(s):

  r = []
  old_i = 0

  for i,l in enumerate(s):
    if l.isupper():
      r.append(s[old_i:i])
      old_i = i
  if old_i < len(s):
    r.append(s[old_i:])
  
  return "".join(w if w.islower() else " "+w for w in r)


# another clever solution from codewars (not my code)
def solution1(s):
  return ''.join(' ' + c if c.isupper() else c for c in s)

print(solution("helloWorld"))
print(solution1("breakCamelCase"))
