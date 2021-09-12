# error-tryer
provides new simply error handling method

### Example:
```go
package main

import (
	"io"
	"net/http"
	"os"

	tryer "github.com/MarkMandriota/error-tryer"
)

func main() {
	resp := tryer.WithFatal(
		http.Get("https://google.com"))[0].(*http.Response)

	tryer.WithFatal(io.Copy(os.Stdout, resp.Body))
}
```

### Create custom Tryer:
```go
http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	try := tryer.New(func(err error) {
		rw.WriteHeader(http.StatusInternalServerError)

		fmt.Print(rw, err)
	}
	
	...
})
```
