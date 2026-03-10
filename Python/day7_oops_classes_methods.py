class Calculator:
    def __init__(self):
        pass
    
    def add(self, a, b):
        return a + b
    
    def subtract(self, a, b):
        return a - b
    
    def multiply(self, a, b):
        return a * b
    
    def divide(self, a, b):
        if b == 0:
            return "Error: Cannot divide by zero"
        return a / b

# Test the Calculator
calc = Calculator()
print(f"10 + 5 = {calc.add(10, 5)}")
print(f"10 - 5 = {calc.subtract(10, 5)}")
print(f"10 * 5 = {calc.multiply(10, 5)}")
print(f"10 / 5 = {calc.divide(10, 5)}")
print(f"10 / 0 = {calc.divide(10, 0)}")

#Create a Book class
class Book:
    def __init__(self, title, author, pages):
        self.title = title
        self.author = author
        self.pages = pages
    
    def display_info(self):
        return f"'{self.title}' by {self.author}, {self.pages} pages"

# Test the Book class
book1 = Book("Python Crash Course", "Eric Matthes", 544)
book2 = Book("Clean Code", "Robert Martin", 464)

print(book1.display_info())
print(book2.display_info())