controller
==========

Just the C in MVC (Supports HMVC).

Ok, so I started this repository to explain how I'm doing HMVC with the standard libary without reflection, without taking control of your middleware etc. 

## Getting Started
From your GOPATH:

```bash
go get github.com/d2g/controller
```

Add a file ```server.go``` - for instance, ```src/myapp/server.go```

```go
package main

import (
	"fmt"
	"github.com/d2g/controller"
	"net/http"
	"strings"
)

type ExampleController struct {
	HelloCount int
	base       string
}

func (t *ExampleController) SetBase(base string) controller.HTTPController {
	t.base = base
	return t
}

func (t *ExampleController) Base() string {
	return t.base
}

func (t *ExampleController) Routes() (http.Handler, error) {
	router := http.NewServeMux()
	router.HandleFunc(t.Base(), t.SayHello)
	return router, nil
}

func (t *ExampleController) SayHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", t.HelloCount), "World!")
}

type ExampleMiddleware struct {
	controller.HTTPController
	HelloCount *int
}

func (t *ExampleMiddleware) Routes() (http.Handler, error) {
	return t, nil
}

func (t *ExampleMiddleware) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	handler, err := t.HTTPController.Routes()
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}

	*t.HelloCount = 3
	handler.ServeHTTP(response, request)
}

func main() {

	example := &ExampleController{
		HelloCount: 1,
	}

	exampleControllers := controller.HTTPControllers([]controller.HTTPController{
		&ExampleMiddleware{
			HTTPController: example.SetBase("/"),
			HelloCount:     &example.HelloCount,
		},
	})

	http.ListenAndServe("localhost:3000", exampleControllers.Routes())
}
```

Run the server. It will be available on ```localhost:3000```:

```bash
go run src/myapp/server.go
```