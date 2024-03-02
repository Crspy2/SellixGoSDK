# Sellix Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/crspy2/SellixGoSDK)](https://pkg.go.dev/github.com/crspy2/SellixGoSDK)
[![Go Report Card](https://goreportcard.com/badge/github.com/crspy2/SellixGoSDK)](https://goreportcard.com/report/github.com/crspy2/SellixGoSDK)
[![License](https://img.shields.io/github/license/crspy2/SellixGoSDK)](https://img.shields.io/github/license/crspy2/SellixGoSDK)


## Introduction

Sellix public API for developers to access merchant resources

## Requirements

- Go 1.11 or later

## Installation

Install the package through NPM.

```
go get github.com/crspy2/SellixGoSDK
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/crspy2/SellixGoSDK"
)

func main() {
	// Get your API Key from https://dashboard.sellix.io/settings/security
	client := SellixGoSDK.NewSellixClient("<API_KEY>")
	// Only call this method if you have multiple shops on your sellix account
	client.SetStoreName("<STORE_NAME>")
	info, err := client.ListProducts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}
```

## Documentation

[Sellix Developers API](https://developers.sellix.io)
