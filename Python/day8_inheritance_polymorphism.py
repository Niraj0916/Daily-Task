# Parent Class
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        return "Animal makes a sound"


# Child Class - Dog
class Dog(Animal):
    def speak(self):
        return "Woof"


# Child Class - Cat
class Cat(Animal):
    def speak(self):
        return "Meow"


# Creating Objects
dog = Dog("Buddy")
cat = Cat("Whiskers")

print(dog.name, "says", dog.speak())
print(cat.name, "says", cat.speak())


# Polymorphism Example
animals = [Dog("Buddy"), Cat("Whiskers")]

for animal in animals:
    print(animal.name, "says", animal.speak())