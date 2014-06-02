package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

/*
 * This is a basic implementation of the HTTPController Interface.
 */
type BasicExampleController struct {
	base string
}

func (t *BasicExampleController) SetBase(base string) HTTPController {
	t.base = base
	return t
}

func (t *BasicExampleController) Base() string {
	return t.base
}

/*
 * The routes that this controller supports
 */
func (t *BasicExampleController) Routes() (http.Handler, error) {
	/*
	 * Lets use the Gorrila Mux
	 */
	router := mux.NewRouter()

	/*
	 * Lets Setup the routes.
	 *
	 * This route is at the base (i.e. / or /user/ etc.).
	 * The controller is only responsible for the base of the url and below.
	 */
	router.HandleFunc(t.Base(), t.index)
	router.HandleFunc(t.Base()+"magic", t.index)
	return router, nil
}

/*
 * Sample handler function.
 */
func (t *BasicExampleController) index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(request.URL.Path))
}

/*
 * Simplistic example, which doesn't really show the advantages.
 */
func ExampleBasicExample() {
	/*
	 * Create an array of controllers to handle our requests
	 * This adds the controller to handle / and /testing/
	 * This means that the URLs that will work are:
	 * *http://localhost/
	 * *http://localhost/magic
	 * *http://localhost/testing/
	 * *http://localhost/testing/magic
	 */
	controllers := HTTPControllers{
		(&BasicExampleController{}).SetBase("/"),
		(&BasicExampleController{}).SetBase("/testing/"),
	}

	log.Fatal(http.ListenAndServe(":80", controllers.Routes()))
}
