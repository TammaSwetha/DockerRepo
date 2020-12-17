package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// handler to get all birds
func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	// convert the bird variabel to json
	birdListBytes, err := json.Marshal(birds)
	//fmt.Fprintf(w, "hello Bird")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// write the JSON list of birds to the response
	w.Write(birdListBytes)
}

// Handler to create a new entry of birds
func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// create a new instance of Bird
	bird := Bird{}

	//send all our data as HTML form data
	// the `ParseForm` method of the request
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Get the information about the bird from the form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	birds = append(birds, bird)

	//we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
