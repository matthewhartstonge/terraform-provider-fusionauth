# Terraform FusionAuth Provider

> I've decided to archive this project as I've turned my efforts to other matters. Now that FusionAuth have taken control and maintenance over the community Terraform provider I suspect my efforts and aims/goals for this rebuild won't be needed as much.

The [FusionAuth Provider](https://registry.terraform.io/providers/matthewhartstonge/fusionauth/latest/docs)
allows [Terraform](https://terraform.io) to manage [FusionAuth](https://fusionauth.io) resources.

- [Development Roadmap](ROADMAP.md)

## Status

This is a pre-alpha project in a quest to create a FusionAuth provider using best practises and the latest `terraform-plugin-framework`. 

This project is not currently taking contributions. This will be reviewed at a later date after the initial project has well-formed patterns.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

 `TODO`

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
