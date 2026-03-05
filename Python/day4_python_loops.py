# Multiplication Table

num = int(input("Enter a number: "))

print("Multiplication Table:")
for i in range(1, 11):
    print(num, "x", i, "=", num * i)


# Factorial Program

n = int(input("Enter a number to find factorial: "))
factorial = 1

for i in range(1, n + 1):
    factorial *= i

print("Factorial of", n, "is", factorial)