# exegol
Go lang package runner, an `npx` implementation for `go`

## Concept

`exegol` installs the package in temporary location, executes the binary, post run cleans up installation

## Installation

```sh
go get github.com/shubhodeep9/exegol
```

## Usage

> Ensure `$GOBIN` is set in `$PATH`
```sh
exegol -uri package -v latest
```

## To-Do

- [ ] Accept argument for the package to be passed ahead