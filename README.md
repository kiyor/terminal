## Terminal ##
Terminal is a simple golang package that provides basic terminal handling.

Terminal wraps color/format functions provided by [ANSI escape code](http://en.wikipedia.org/wiki/ANSI_escape_code)

## Usage ##
```go
package main

import (
  "github.com/kiyor/terminal"
  "github.com/kiyor/terminal/color"
)

func main() {
  terminal.Stdout.Color("y").
    Print("Hello world").Nl().
    Reset().
    Colorf("@{kW}Hello world\n")

  color.Print("@rHello world")
  color.Print("@{216 8}Hello world")
}
```
Check the [godoc result](https://godoc.org/github.com/kiyor/terminal) for more details.
