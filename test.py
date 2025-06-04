import random
n = random.randint(3, 10)
print(n)
l = [i for i in range(n)]
random.shuffle(l)
for i in l: print(i)
