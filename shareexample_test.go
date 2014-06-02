package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"testing"
)

/*
 * This example shows how you would pass shared resources to a controller.
 * For example a connection pool
 */
type ShareExampleController struct {
	base  string
	Share string //Shared resource, could be pointer anything I've done a sting as a really basic example.
}

func (t *ShareExampleController) SetBase(base string) HTTPController {
	t.base = base
	return t
}

func (t *ShareExampleController) Base() string {
	return t.base
}

/*
 * Change the Routes this controller will handle
 */
func (t *ShareExampleController) Routes() (http.Handler, error) {
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
	router.HandleFunc(t.Base(), t.share)
	return router, nil
}

/*
 * Create Our own Index.
 */
func (t *ShareExampleController) share(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello.. URL:%q Share:%s", html.EscapeString(request.URL.Path), t.Share)
}

/*
 * Share a resource to the controller
 */
func ExampleShareExample() {
	/*
	 * Create an array of controllers to handle our requests
	 * This adds the controller to handle / and /testing/
	 * This means that the URLs that will work are:
	 * *http://localhost/
	 * *http://localhost/testing/
	 */
	controllers := HTTPControllers{
		(&ShareExampleController{
			Share: "Second Share",
		}).SetBase("/testing/"),
		(&ShareExampleController{
			Share: "First Share",
		}).SetBase("/"),
	}

	log.Fatal(http.ListenAndServe(":80", controllers.Routes()))
}
