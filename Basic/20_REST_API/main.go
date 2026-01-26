/*
 - by default any method is 'get' if it is not defined.


 # Status Code's :
		- 1xx  -> Informational SC
		- 2xx  -> Success SC
		- 3xx  -> Redirectional SC
		- 4xx  -> Client side Error
		- 5xx  -> Server side Error

  - 100 = continue, server request is reseived but client is still sending the body/request
  - 101 = switching protocol : websocket , http2

  - 200 = ok : request successfully return data
  - 201 = created : new resource creted : POST request
  - 202 = accepted : request is accepted but it is still in processing (gorutines, asynchronous jobs)
  - 204 = No containt : successfully run  (Delete)

  - 301 = Moved permanuntly - resource moved to new position/url (redirect)
  - 302 = Temporarly moved
  - 304 = not modified (caching)

  - 400 = Bad request (invalid data/request. wrong json, missing data)
  - 401 = Unauthorized (Authantication required. Accessing data without login)
  - 403 = Forbidden (You are authanticated but not allowed to access that perticular resource)
  - 404 = not found
  - 405 = method not allowed
  - 409 = conflict (user is allready exist, version missmatch)
  - 429 = to many request

  - 500 = Internal server error
  - 501 = Not implemented (endpoint present but method is not implemented)
  - 502 = Bad Getway ( Serve active as a proxy , invalid response return)
  - 503 = Service unavailable (server temporaly overloaded or down)
  - 504 = Getway timeout (server didn't get response in time)


*/

package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	fmt.Println("Yes api run successfully....")
}

func main() {
	http.HandleFunc("/hello", HelloWorld)

	fmt.Println("Server is running on 7757")
	http.ListenAndServe(":7757", nil)
}
