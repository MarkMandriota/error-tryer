# error-tryer
provides new simply error handling method

### Using default tryers:
```go
package main

import (
	"io"
	"net/http"
	"os"

	tryer "github.com/MarkMandriota/error-tryer"
)

func main() {
	tryer.WithFatal(io.Copy(os.Stdout, tryer.WithFatal(
		http.Get("https://google.com"))[0].(*http.Response).Body))
}
```

### Create custom tryer:
```go
http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	mytryer := tryer.Compose(tryer.New(func(err error) {
		rw.WriteHeader(http.StatusInternalServerError)
	}, tryer.Printer(log.Default()))

	...
})
```