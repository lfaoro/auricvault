# AuricVault API library

Simple Go-API used to store sensitive data using the AuricVault PCI/HIPAA compliant encryption service.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/lfaoro/creditcard)](https://goreportcard.com/report/github.com/lfaoro/auricvault)
[![Go Doc](https://img.shields.io/badge/go%20doc-read-blue.svg?style=flat)](https://godoc.org/github.com/lfaoro/auricvault)

## Installation
```bash
$ go get -u github.com/lfaoro/auricvault
```

```go
import "github.com/lfaoro/auricvault"
```

## Quick start
```bash
# Provide a .env file in your project with the following variables or export them.
# the .env file will be automatically parsed.
$ cat > .env << EOF
AURIC_URL="https://vault01.auricsystems.com/vault/v2/"
AURIC_URL2="https://vault02.auricsystems.com/vault/v2/" # optional
AURIC_CONFIGURATION=""
AURIC_MTID=""
AURIC_MTID_SECRET=""
AURIC_SEGMENT=""
EOF
```

```go
// Instance a new Vault, choose the retention period
vault := auricvault.New(auricvault.Forever)

// If you want to see Debug information
auricvault.SetDebug()

data := "VISA,475055XXXX314032,0818"

// Encrypt using an auto-generated token
token, err := vault.Encrypt(data)
if err != nil {
    log.Fatal(err)
}
fmt.Println("token: ", token)

// Encrypt using your own token
token, err := vault.Encrypt(data, "khR8pew41q0URCxtivea")
if err != nil {
    log.Fatal(err)
}
fmt.Println("token: ", token)

// Retrieve the string data from the vault using the token
data, err := vault.Decrypt("khR8pew41q0URCxtivea")
if err != nil {
    log.Fatal(err)
}
fmt.Println(data)
```

## Disclaimer
Some AuricVault API features and settings have been omitted in favor of simplicity. If you require any of the omitted features implemented, please open an [issue](https://github.com/lfaoro/auricvault/issues/new).

# Contibuting
> Any help and suggestions are very welcome and appreciated. Start by opening an [issue](https://github.com/lfaoro/auricvault/issues/new).

- Fork the project
- Create your feature branch `git checkout -b my-new-feature`
- Commit your changes `git commit -am 'Add my feature'`
- Push to the branch `git push origin my-new-feature`
- Create a new pull request against the master branch
