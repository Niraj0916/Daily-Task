class BankAccount:

    def __init__(self, owner, balance, pin):
        self.owner = owner
        self.__balance = balance
        self.__pin = pin

    def deposit(self, amount):
        if amount > 0:
            self.__balance += amount
            print("Deposit successful")

    def withdraw(self, amount, entered_pin):

        if entered_pin != self.__pin:
            print("Incorrect PIN")
            return

        if amount <= self.__balance:
            self.__balance -= amount
            print("Withdrawal successful")
        else:
            print("Insufficient balance")

    def check_balance(self, entered_pin):

        if entered_pin != self.__pin:
            print("Incorrect PIN")
            return

        print("Current Balance:", self.__balance)


def run_bank_account():

    owner = input("Enter account holder name: ")
    balance = float(input("Enter initial balance: "))
    pin = input("Set a 4-digit PIN: ")

    account = BankAccount(owner, balance, pin)

    while True:

        print("\n1. Deposit")
        print("2. Withdraw")
        print("3. Check Balance")
        print("4. Exit")

        choice = input("Enter choice: ")

        if choice == "1":
            amount = float(input("Enter deposit amount: "))
            account.deposit(amount)

        elif choice == "2":
            amount = float(input("Enter withdrawal amount: "))
            entered_pin = input("Enter PIN: ")
            account.withdraw(amount, entered_pin)

        elif choice == "3":
            entered_pin = input("Enter PIN: ")
            account.check_balance(entered_pin)

        elif choice == "4":
            print("Exiting...")
            break

        else:
            print("Invalid choice")


run_bank_account()

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
