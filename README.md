# Pivot Security - Go

GO API interface for Pivot Security API.

All updates to this library is documented in our [CHANGELOG](https://github.com/pivotsecurity/pivotsecurity-go/blob/master/CHANGELOG.md).

# Table of Contents
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [License](#license)

<a name="installation"></a>
# Installation

## Prerequisites

- Go version 1.6.X, 1.7.X, 1.8.X, 1.9.X or 1.10.X

## Install Package

```bash
go get github.com/pivotsecurity/pivotsecurity-go
```

## Setup Environment Variables

### Initial Setup

```bash
cp .env_sample .env
```

### Environment Variable

Update the development environment with your keys(https://api.pivotsecurity.com/settings/), for example:

```bash
echo "export PUBLIC_API_KEY='YOUR_PUBLIC_API_KEY'" > pivotsecurity.env
echo "export PRIVATE_API_KEY='YOUR_PRIVATE_API_KEY'" > pivotsecurity.env
echo "pivotsecurity.env" >> .gitignore
source ./pivotsecurity.env
```

<a name="quick-start"></a>
# Quick Start

`Info call` with user id (or email).

```go
package main

import "github.com/pivotsecurity/pivotsecurity-go"
import "fmt"

func main() {
	const host = "https://api.example.com"
	uid := "uid"
	email := "email"
	response, err := account.Send(uid, email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```

<a name="license"></a>
# License
[The MIT License (MIT)](LICENSE)


