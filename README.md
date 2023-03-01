# lpdice
Go package for dice

[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/lpdice)](https://goreportcard.com/report/github.com/thorstenrie/lpdice)
[![CodeFactor](https://www.codefactor.io/repository/github/thorstenrie/lpdice/badge)](https://www.codefactor.io/repository/github/thorstenrie/lpdice)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorstenrie/lpdice)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/lpdice)](https://pkg.go.dev/mod/github.com/thorstenrie/lpdice)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorstenrie/lpdice)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorstenrie/lpdice)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorstenrie/lpdice)
![GitHub last commit](https://img.shields.io/github/last-commit/thorstenrie/lpdice)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorstenrie/lpdice)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorstenrie/lpdice)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorstenrie/lpdice)
![GitHub](https://img.shields.io/github/license/thorstenrie/lpdice)

The package lpdice provides a simple interface for random results from rolling symmetrical dice with 4, 6, 8, 10, 12, 20 sides. A die is retrieved by
an instance of struct type `Die`. The standard `Die` has six sides. A new die with 4, 6, 8, 10, 12 or 20 sides can also be retrieved with
`NewD{4, 6, 8, 10, 12, 20}`. The underlying random number generator is a
cryptographically secure random number generator based on `crypto/rand`. If the cryptographically secure random
number generator source is not available on the system,
a pseudo-random number generator based on `math/rand` is used. A die can be seeded to generate deterministic random
numbers based on a seed.

- *Simple*: Without configuration, just function calls
- *Easy to use*: Retrieve a die by struct type Die with a simple interface
- *Tested*: Unit tests with high code coverage
- *Dependencies*: Only depends on the [Go standard library](https://pkg.go.dev/std), [lpstats](https://github.com/thorstenrie/lpstats), [tserr](https://github.com/thorstenrie/tserr) and [tsrand](https://github.com/thorstenrie/tsrand)

## Usage

The package is installed with

```
go get github.com/thorstenrie/lpdice
```

In the Go app, the package is imported with

```
import "github.com/thorstenrie/lpdice
```

### Dice

An instance of struct type `Die` has a fixed number of sides. The directly instantiated standard `Die` has six sides.
A 4, 6, 8, 10, 12, 20 sided die can be retrieved using one of the `NewD{4, 6, 8, 10, 12, 20}` functions. A die can be rolled by calling
`Roll`. The result will be randomly generated. To retrieve a random but deterministic series of results, a `Die` can be seeded with `Seed`.
With `NoSeed` the `Die` behavior can be changed back to the non-seeded random number generation.

## Random number generators

A `Die` holds random number generators to generated a result from rolling the die. It contains a pointer to a cryptographically secure random number generator
for non-seeded results and a pointer to a deterministic pseudo-random number generator for seeded results. If the cryptographically secure random number generator
is not available on the platform, a pseudo-random number generator will be used. The deterministic pseudo-random number generator will be used, if the `Die` is seeded
by calling `Seed`.

## Unit tests

The quality of random results from rolling the die depends on the random number generator source. The unit tests cover a basic evaluation of the results. For the basic evaluation covered by the unit tests, the test functions generate random results from rolling all available dice. The test functions compare for each die the aritmetic mean and variance of the retrieved random numbers with the expected values for mean and variance. If the arithmetic mean and variance of the retrieved random numbers are not near equal to expected values, the test fails. Therefore, the unit tests provide an indication if the random number generator sources are providing random values in expected boundaries. However, the unit tests do not evaluate the quality of retrieved random numbers in different dimensions or the implementation of the random number generator source. The output of the random number generator sources might be easily predictable.

## Example

```
TODO
```

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/lpdice)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/lpdice)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Flpdice)
