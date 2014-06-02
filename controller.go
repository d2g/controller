package controller

import "net/http"

/*
 * Controllers need to respond with the routes they handle.
 * This means all routing information is self contained.
 */
type HTTPController interface {
	/*
	 * Return the Routes that are managed by this Controller.
	 */
	Routes() (http.Handler, error)

	/*
	 * Sets the Base URL (i.e. /devices/ or /users/ or / ...)
	 */
	SetBase(url string) HTTPController
	/*
	 * Return The Base URL (Needed for Setting Up routing)
	 */
	Base() string
}

type HTTPControllers []HTTPController

/*
 * Return the standard http.Handler for use with the standard net/http
 */
func (t *HTTPControllers) Routes() http.Handler {
	router := http.NewServeMux()

	for _, pathcontroller := range []HTTPController(*t) {
		subRoutes, err := pathcontroller.Routes()
		if err != nil {
			panic("Error in Controller (Path:" + pathcontroller.Base() + ")")
		}

		router.Handle(pathcontroller.Base(), subRoutes)
	}

	return router
}
