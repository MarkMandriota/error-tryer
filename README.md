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
  defer tryer.TryWithFatal(recover())
  
  try := tryer.TryWithPanic
  
  resp := try(
    http.Get("https://google.com"))[0].(*http.Response)
    
  written := try(
    io.Copy(os.Stdout, resp.Body))[0].(*int64)
    
  fmt.Printf("\nWritten: %d\n", *written)
}
```
