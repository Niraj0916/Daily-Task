class BankAccount:

    def __init__(self, owner, balance):
        self.owner = owner
        self.__balance = balance   # private variable

    def deposit(self, amount):
        if amount > 0:
            self.__balance += amount
            print("Deposit successful")
        else:
            print("Invalid deposit")

    def withdraw(self, amount):
        if amount <= self.__balance:
            self.__balance -= amount
            print("Withdrawal successful")
        else:
            print("Insufficient funds")

    def get_balance(self):
        return self.__balance


def run_bank_example():

    account = BankAccount("Niraj", 1000)

    account.deposit(500)
    account.withdraw(200)

    print("Current Balance:", account.get_balance())
run_bank_example()

#Abstraction
from abc import ABC, abstractmethod


class Shape(ABC):

    @abstractmethod
    def area(self):
        pass


class Circle(Shape):

    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return 3.14 * self.radius * self.radius


class Rectangle(Shape):

    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height


def run_shape_example():

    circle = Circle(5)
    rectangle = Rectangle(4, 6)

    print("Circle Area:", circle.area())
    print("Rectangle Area:", rectangle.area())


run_shape_example()
