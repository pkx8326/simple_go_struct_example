# Simpe Go Struct Example

### Overview
This is a very simple program written in the Go Programming Language to demonstrate the use of the struct data type.

### Program manual
When run, the program will ask the user to input a student's first name, last name (both of string type) and the student's score (floating number). After that, the program will just display all the input information in another set of lines. There is no input validation for this very simple program since people's first names or last names can be anything. There is nither an input validation for the score. If the user inputs something which is not a floating number, the program will return 0 (the default value of the variable that stores the score).

### Code structure
The code of this program is divided into two main parts: the main program and the functions section which includes the function for getting data and storing it in a struct-type variable and a function which takes a struct as an argument and display the information stored in that struct. This is to demonstrate possible uses of struct-type variables and functions in the Go Programming Language.

### Program flow
The program flow is shown as the following numbered list:
1. The program calls the function that reads in user inputs, including a student's first name and last name (strings), and the student's score (floating number) then store those information in a struct variable. The function then returns the stored information as a struct to a variable declared in the main program.
2. The program then calls on the function to display the stored information. This function takes the struct variable returned by the first function as an argument. It reads each field of the struct and displays the information accordingly.
