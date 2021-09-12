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
  defer tryer.WithFatal(recover())
  
  resp := tryer.WithPanic(
    http.Get("https://google.com"))[0].(*http.Response)
    
  io.Copy(os.Stdout, resp.Body)
}
```
