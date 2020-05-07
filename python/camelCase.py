# Python episode N01
# challenge N02 => (Camel Case Method) - 6Kyu

# By ALGORITHM


def camel_case(string):
  return ''.join(w.capitalize() for w in string.split())

print(camel_case('hello world'))
print(camel_case('camel case method'))

