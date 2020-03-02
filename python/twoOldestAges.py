# Python episode N03
# challenge N07 => (Two Oldest Ages) - 7Kyu

# By ALGORITHM

# method 1
def two_oldest_ages_1(ages):
  ages.sort()
  return ages[-2:]

# method 2
def two_oldest_ages_2(ages):
  for i in range(len(ages)):
    min_val_index = i
    for j in range(i+1, len(ages)):
      if ages[j] < ages[min_val_index]:
        min_val_index = j
    ages[i], ages[min_val_index] = ages[min_val_index], ages[i]
  return ages[-2:]


print(two_oldest_ages_1([1, 3, 10, 0]))
print(two_oldest_ages_2([1, 3, 10, 0]))
