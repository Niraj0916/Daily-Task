## Day 1
- Installed Python
- Installed Go
- Ran Hello World programs

## Day 2

Python:
Learned variables and data types such as string, integer, float, and boolean.

Go:
Learned variable declaration and basic data types.

GitHub:
Added and pushed Day 2 programs to the repository.

## Day 3

Python:
Implemented if-else programs to check even/odd numbers and find the maximum number.

Go:
Implemented the same logic using Go if-else statements.

GitHub:
Used git add, commit, and push to upload Day 3 programs.

## Day 4 

Python:
Learned how loops work using the `for` loop and practiced working with lists, dictionaries, and functions.

Programs implemented:
- Multiplication table
- Factorial calculation
- List operations (adding and printing elements)
- Dictionary example (storing and accessing key-value data)
- Calculator program using functions

Go (Golang):
Practiced loops using Go's `for` statement and learned arrays, slices, and functions.

Programs implemented:
- Multiplication table
- Factorial program
- Fibonacci series
- Array traversal program
- Slice example
- Calculator program using functions

Concepts Learned:
- Looping concepts
- Repetition in programming
- Basic algorithm practice
- Working with arrays and slices in Go
- Using functions to organize code

## Day 5

Python: Built a simple CLI todo application and learned how modules, imports, and error handling work in Python.

Programs implemented:

- Simple command line todo application
- Modules and imports example using datetime
- Error handling using try and except

Go (Golang): Learned how structs organize data and how Go handles errors.

Programs implemented:

- Student struct example
- Book struct practice program
- Error handling using strconv conversion

Concepts Learned:

- Structs for organizing related data
- Error handling in Python using try/except
- Error handling in Go using error values
- Building simple command line tools

## Day 7

Python: Practiced object-oriented programming by creating classes and methods to represent real-world objects.

Programs implemented:

* Calculator class with add, subtract, multiply, and divide methods
* Book class with title, author, and pages attributes
* Display function to print book information

Go (Golang): Learned how Go uses structs and receiver methods to implement object-oriented style programming.

Programs implemented:

* Calculator struct with methods for addition, subtraction, multiplication, and division
* Book struct with fields for title, author, and pages
* DisplayInfo() method to print formatted book information

Concepts Learned:

* Classes and objects in Python
* Structs as Go’s alternative to classes
* Methods attached to structs using receivers
* Organizing related data and behavior together
* Basic error handling when dividing by zero

## Day 8

Python: Practiced object-oriented programming using inheritance and polymorphism to model animal behavior.

Programs implemented:

* Animal parent class with name attribute and speak() method
* Dog child class that inherits from Animal and overrides speak() to return "Woof"
* Cat child class that inherits from Animal and overrides speak() to return "Meow"
* Created objects of Dog and Cat to demonstrate different behaviors
* Used a list of animals and looped through them to demonstrate polymorphism

Go (Golang): Learned how Go implements object-oriented behavior using struct composition and interfaces instead of inheritance.

Programs implemented:

* Animal struct with name field
* Dog struct embedding Animal to demonstrate composition
* Cat struct embedding Animal
* Speak() methods implemented for Dog and Cat
* Speaker interface with Speak() method
* Created a slice of Speaker interface type and looped through it to demonstrate polymorphism

Concepts Learned:

* Inheritance in Python using parent and child classes
* Method overriding in child classes
* Polymorphism using the same method name with different outputs
* Composition in Go as an alternative to inheritance
* Interfaces in Go to define shared behavior
* Using slices of interfaces to achieve polymorphism

## Day 9

Python: Practiced OOP concepts using encapsulation and abstraction to model secure data handling and enforce method implementation.

Programs implemented:

* BankAccount class with private balance using encapsulation
* Deposit and withdraw methods with validation logic
* Shape abstract base class with abstract area() method
* Circle and Rectangle classes implementing area calculation


Go (Golang): Implemented struct methods and encapsulation concepts using receiver functions and validation logic.

Programs implemented:

* BankAccount struct with owner and balance fields
* Deposit and withdraw methods using pointer receivers
* Validation logic to prevent invalid transactions
* Getter method to retrieve account balance


Concepts Learned:

* Encapsulation for protecting internal data in Python
* Private variables and controlled access using methods
* Abstract base classes and enforced method implementation
* Struct methods and pointer receivers in Go
* Data validation in application logic

## Day 10

Python: Implemented a simple web service using Flask and containerized the application using Docker to run inside an isolated environment.

Programs implemented:

* Flask web server with a route returning a simple message
* Application configured to run on port 5000 and accessible through Docker
* Dockerfile created to package the Python application and install dependencies using requirements.txt

Go (Golang): Implemented a lightweight HTTP server using Go’s net/http package and containerized the application using Docker.

Programs implemented:

* HTTP web server responding with a simple message using the net/http package
* Handler function for the root route ("/") returning a response to the browser
* Dockerfile created to compile the Go program into a binary and run it inside a container

Concepts Learned:

* Containerization using Docker to package applications with their dependencies
* Building Docker images using Dockerfile instructions
* Running containerized services and exposing ports for browser access
* Using Docker Compose to run multiple services (Python and Go) together
* Basic microservice architecture with separate containers for different services
