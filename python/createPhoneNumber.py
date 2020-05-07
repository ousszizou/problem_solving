# Python episode N01
# challenge N03 => (Create Phone Number) - 6Kyu

# By ALGORITHM


def create_phone_number(n):
  n = ''.join(map(str,n))
  return "({}) {}-{}".format(n[:3],n[3:6],n[6:])

print(create_phone_number([1, 2, 3, 4, 5, 6, 7, 8, 9, 0]))
print(create_phone_number([1, 1, 1, 1, 1, 1, 1, 1, 1, 1]))
