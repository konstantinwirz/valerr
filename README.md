# valerr

[![All builds](https://github.com/konstantinwirz/valerr/actions/workflows/main.yaml/badge.svg)](https://github.com/konstantinwirz/valerr/actions/workflows/main.yaml)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/oklog/run/master/LICENSE)

I caught myself implementing the validation error type in every application i implemented, so i decided to put it in a separate library.

## Example

```go
package main

import "github.com/konstantinwirz/valerr"

func createAccount(email string, pwd []byte) error {
	// some validation checks
	return valerr.NewValidationError(
		valerr.NewViolation("email", "not well formed"),
		valerr.NewViolation("password", "insecure"),
	)
}

func main() {

	err := createAccount("email", []byte(""))
	if ok, verr := err.(valerr.ValidationError); ok {
		// handle violations
    }
}
```
