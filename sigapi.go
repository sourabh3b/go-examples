package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/tekion/apiframework"
)

func main() {
	// create an empty context, you can configure a set of global flags here
	ctx := apiframework.APIContext{}

	// create a router,
	// pass the address of the context,
	// pass the APIRoutes collection
	// set the url route prefrix
	router := apiframework.NewRouter(
		&ctx,
		apiframework.APIRoutes{
			apiframework.APIRoute{
				Name:        "test",
				Method:      "GET",
				Pattern:     "/test",
				HandlerFunc: testRequest,
			},
		},
		"")

	// set the http server based on the router listening on port 9000
	// you can run this app http://localhost:9000/test
	err := http.ListenAndServe("localhost:9000", router)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func testRequest(ctx *apiframework.APIContext, w http.ResponseWriter, r *http.Request) (errorCode int, err error) {
	w.Write([]byte("Hello World!"))
	return 0, nil
}
