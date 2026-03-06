import datetime


def todo_cli():
    tasks = []

    print("----- Simple Todo CLI -----")

    while True:
        print("\n1. Add Task")
        print("2. View Tasks")
        print("3. Exit")

        choice = input("Enter choice: ")

        if choice == "1":
            task = input("Enter task: ")
            tasks.append(task)
            print("Task added successfully!")

        elif choice == "2":
            print("\nYour Tasks:")
            if len(tasks) == 0:
                print("No tasks available")
            else:
                for i, t in enumerate(tasks, start=1):
                    print(f"{i}. {t}")

        elif choice == "3":
            print("Exiting Todo App")
            break

        else:
            print("Invalid choice")


def module_import_example():
    print("\n----- Modules & Imports Example -----")

    now = datetime.datetime.now()

    print("Current Date:", now.date())
    print("Current Time:", now.time())


def error_handling_example():
    print("\n----- Error Handling Example -----")

    try:
        num = int(input("Enter a number: "))
        result = 10 / num
        print("Result:", result)

    except ValueError:
        print("Error: Please enter a valid number")

    except ZeroDivisionError:
        print("Error: Cannot divide by zero")


def day5():
    todo_cli()
    module_import_example()
    error_handling_example()


if __name__ == "__main__":
    day5()