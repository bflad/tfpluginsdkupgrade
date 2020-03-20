# tfpluginsdkupgrade

A proof of concept [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) upgrade tool for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code.

## Install

### Local Install

To install into your `$GOBIN` directory (e.g. `$GOPATH/bin`):

```console
$ go get github.com/bflad/tfpluginsdkupgrade
```

## Usage

Additional information about usage and configuration options can be found by passing the `help` argument:

```console
$ tfpluginsdkupgrade help
```

### Local Usage

To report issues, change into the directory of the Terraform Provider code and run:

```console
$ tfpluginsdkupgrade ./...
```

To automatically fix some issues, change into the directory of the Terraform Provider code and run:

```console
$ tfpluginsdkupgrade -fix ./...
```

## Development and Testing

This project is built on the [`go/analysis`](https://godoc.org/golang.org/x/tools/go/analysis) framework and uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

Helpful tooling for development:

* [`astdump`](https://github.com/wingyplus/astdump): a tool for displaying the AST form of Go file

### Updating Dependencies

```console
$ go get URL
$ go mod tidy
$ go mod vendor
```

### Unit Testing

```console
$ go test ./...
```

### Local Install Testing

```console
$ go install .
```
