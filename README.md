# exegol
Go lang package runner, an `npx` implementation for `go`

## Concept

`exegol` installs the package in temporary location, executes the binary, post run cleans up installation

## Installation

```sh
go get github.com/shubhodeep9/exegol

// latest versions of go
go install github.com/shubhodeep9/exegol@latest
```

## Usage

> Ensure `$GOBIN` is set in `$PATH`
```sh
exegol -uri package -v latest -arg "-sample arg"
```