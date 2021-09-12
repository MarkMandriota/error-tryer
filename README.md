# error-tryer
provides new simply error handling method

### Example:
```go
package main

import (
  "io"
  "net/http"
  
  "github.com/MarkMandriota/error-tryer"
)

func main() {
  defer tryer.TryWithPrint(recover())
  
  resp := tryer.TryWithPanic(http.Get("https://google.com"))[0].(*http.Response)
  fmt.Printf("\nWritten: %d\n", *tryer.TryWithPanic(io.Copy(os.Stdout, resp.Body))[0].(*int64))
}
```
