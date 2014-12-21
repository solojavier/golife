# Golife

[![travis](https://travis-ci.org/solojavier/golife.svg)](https://travis-ci.org/solojavier/golife)

An implementation of [Conway's Game of life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in golang

Main program has an infinite loop that uses `Conway`'s rules to generate
a next universe generation and print it to the screen.

It uses a grid universe with wrapped edges, but it could be switched
to another type of universe thanks to the usage of interfaces.

You could also write your own rules and use them instead of default ones.

## Build and run

If you are not familiar with golang code, visit this page for more information: https://golang.org/cmd/go/

## Test

Use standard command:

```
$ go test
```

## Feeback/Issues

Feel free to open an issue in this repository for any feeback of issues you may find.
