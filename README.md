# Learning Go

The Go programming language is perfect for console applications as well as servers/databases.

This repository is where I have been learning the Go programming language. Feel free to look around!

Not only can this teach you some fundementals of Go, but it can be a great revision guide if you forget something in the future.

# Notes

Semicolons are used to seperate lines of code but are not required in source code, the lexer will automatically insert them at the end of lines. The only time they must be included is when writing multiple blocks of code on one line, for example `count += 1; return count`.

Declaring a variable uses the `var` keyword but this can be ommitted if immediately initializing the variable with a value by using the `:=` operator. `var count; count = 1;` becomes `count := 1`.

A while loop is declared using `for (condition)`.

Golang OOP uses some golang-specific vocabulary. This document goes over terms in previously known concepts: https://github.com/luciotato/golang-notes/blob/master/OOP.md

# Commands

| Commands | Description |
| --- | --- |
| go run path/to/file | Runs the file |
| go run . | Run .go file in current directory (can only be one file in working directory) |
| go build path/to/file | Builds the file
| go mod init example/hello | Creates a module under name 'example/hello' |
