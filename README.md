# MTCS-Go

This is a Go project that includes a server implementation. The server listens for incoming TCP connections on port 12345 and handles each connection in a separate goroutine.

## WIP - TODO
- encryption
- client code
- clean up

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).

### Installing

To install the project, you can clone the repository using the following command:

```sh
git clone git@github.com:aeegvk/mtcs-go.git
```

Then, navigate to the project directory and download the necessary Go modules:

```sh
cd mtcs-go
go mod download
```

### Running the Server

To start the server, navigate to the `cmd/server` directory and run the `server.go` file:

```sh
cd cmd/server
go run server.go
```

The server will start listening for incoming connections on `localhost:12345`.

## Running the Tests

To run the tests for the server, navigate to the `cmd/server` directory and run the following command:

```sh
go test
```
or
```sh
go test ./cmd/server
```

## Built With

* [Go](https://golang.org/) - The programming language used

## Authors

* **Aeegvk** - *Initial work* - [aeegvk](https://github.com/aeegvk)

## License

This project is licensed under the [MIT License](LICENSE)