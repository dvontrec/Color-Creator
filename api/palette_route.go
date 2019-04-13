package main

import "net/http"

// Function used to get pallette by the request user or color
func getPallette(w http.ResponseWriter, req *http.Request) {
	// This function takes data from the request
	// If the request has a user it grabs the palettes created by the user
	// If the request has a hex provided it grabs palettes that use that color
	// The palette is saved to an array
	// The array is passed as json

}
