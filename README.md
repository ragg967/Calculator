# Calculator

Calculator is a Command Line Interface (CLI) calculator written in Go.

## Features

- Supports basic arithmetic operations: addition, subtraction, multiplication, and division.
- Supports exponentiation and square root operations.
- Handles parenthesized expressions for proper order of operations.

## Installation

To install the calculator, clone the repository and build the project using Go:

```sh
git clone https://github.com/ragg967/Calculator.git
cd Calculator
go build -o calculator main.go
```

## Usage

To run the calculator, execute the built binary:

```sh
./calculator
```

You will be prompted to enter an expression. Type your expression and press Enter to see the result. Type `exit` to quit the calculator or `operators` to see the list of supported operators.

### Example

```sh
GoCalculator - Enter an expression ('exit' to quit and 'operators' to show all operators and what they do):
> 5 + 5
Result: 10
> 10 / 2
Result: 5
> 2 ^ 3
Result: 8
> 25 % 2
Result: 5
> exit
Goodbye!
```

## Supported Operators

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Exponentiation: `^`
- Square Root: `%`

## Contribution

If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
