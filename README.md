# Monkey Language

This is a Basic Go implementation of Thorsten Ball's "Writing an Interpreter in Go" book. It involves creating an interpreted language in Go that does basic functions.

## Prerequisites

Before you continue, ensure you meet the following requirements:

- You have installed the latest version of Golang.
- You are using a Linux or Mac OS machine. Windows is not currently supported.
- You have git installed on your machine.

## How to install

- change directory into a folder you want to use the monkey language:
- Clone the repository into your machine

```bash
$ git clone https://github.com/wesleymutwiri/monkey-language.git
```

- change directory into the newly cloned application in your machine by using the following command

```bash
$ cd monkey-language
```

- To run the language is as simple as running

```bash
$ go run main.go
Welcome to the Monkey Programming Language.
>>
```

- In order to run the tests run this simple command `bash go test ./...` in order to run all tests written for the monkey language
  You can also test the packages individually by using `bash go test ./<name_of_package_to_test>`

## Sample Language functions

```bash
>> let add = fn(a, b, c, d) { return a + b + c + d };

>> add(1, 2, 3, 4);
10

>> let addThree = fn(x) { return x + 3 };

>> addThree(3);
6

>> let max = fn(x, y) { if (x > y) { x } else { y } };

>> max(5, 10)
10

>> let factorial = fn(n) { if (n == 0) { 1 } else { n * factorial(n - 1) } };

>> factorial(5)
120

>> let callTwoTimes = fn(x, func) { func(func(x)) };

>> callTwoTimes(3, addThree);
9

>> callTwoTimes(3, fn(x) { x + 1});
5

>> let newAdder = fn(x) { fn(n) { x + n } };

>> let addTwo = newAdder(2);

>> addTwo(2);
4

>> let multiply = fn(x, y) { x * y };

>> multiply(50 / 2, 1 * 2)
50

>> fn(x) { x == 10 }(5)
false

>> fn(x) { x == 10}(10)
true

```

## Types in the language

Type | Syntax |  
\_ \_ \_ _|_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_ \_|
null | `null` |
bool | `true false` |
int | `0 42 1234 5` |
str | `"" "foo" "\"something\" and a\n line break"` |
array | `[] [1, 2] [1, 2, 3]` |
hash | `{} {"a": 1} {"a": 1, "b": 2}` |

## Features to be Implemented

- [x] Create a lexer
- [x] Create a REPL for evaluating the language
- [x] Create a Parser
- [x] Evaluating Symbols,functions, function calls etc
- [x] Extending the interpreter to include data structures
- [ ] Adding additional tests
- [ ] Documenting how to use the language
- [ ] Macro systems for the language
- [ ] Other features
- [ ] Writing a compiler for the language
- [ ] Packaging the language as an executable
- [ ] Adding the language to a package registry(e.g snap)

## Reason for the project

The reason for starting the project was just pure curiosity about how interpreters and compilers work and how people create new programming languages and why. The books chosen were due to the fact that I have Go knowledge and because they are well structured and follow a simple enough approach to follow and understand what is going on.

I highly recommend the book for anyone not only Gophers since there are other implementations in different languages.

See [Official Monkey website](https://monkeylang.org/) for implementations in other languages that may be interesting for other developers who are not interested in typing Go code but interested in learning about the monkey language and about interpreters and compilers but in other languages.

## Basic Definitions and Knowledge learnt

coming soon

### Website for presenting my own version of the Monkey language

Coming soon

## Acknowledgements

[Writing an interpreter in Go](https://interpreterbook.com/) for giving me the basics needed in creating and implementing the language in Golang
