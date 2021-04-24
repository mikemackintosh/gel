# gel
---

Gel is a Go utility designed to support a fluid configuration, which includes loading options from Environments, Flags and Configuration files.

## Usage

Install `Gel` with `go get`:

    go get github.com/mikemackintosh/gel

## Example
```
package main

import (
  "fmt"
  "github.com/mikemackintosh/gel"
)
func init() {
    gel.UseOrder(gel.Flags, gel.Env, gel.Config)
    gel.String("FLAG_A", "unset", "This is defaulted to 'unset'")
    gel.String("FLAG_B", "unset", "This is defaulted to 'unset'")
}

func main() {
  gel.Up()

  fmt.Printf("FLAG_A = %s\n", gel.MustGet("FLAG_A").String())
  fmt.Printf("FLAG_B = %s\n", gel.MustGet("FLAG_B").String())
}

```
