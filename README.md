# How to Install

This project requires `gomail` as a dependency.

```bash
go get "gopkg.in/gomail.v2"
```

# How to Use

A Sendmail/Postfix server must be running and accepting connections on port 587.

Then, run `go run mailpimch.go -input=./path/to/file/listing/email_addresses`.

# Tests

Run `go test`.
