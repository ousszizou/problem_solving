# Python episode N01
# challenge N01 => (Century From Year) - 8Kyu

# By ALGORITHM

def century(year):
  if year % 100 == 0:
    cent = year // 100
    return cent
  else:
    cent = (year // 100) + 1
    return cent


print(century(1705))
print(century(1900))
