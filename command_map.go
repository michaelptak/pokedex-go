package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(conf *config) error {
	var url string
	if conf.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *conf.Next
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locations Locations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	conf.Next = locations.Next
	conf.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	var url string
	if conf.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else if conf.Previous == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *conf.Previous
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locations Locations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	conf.Next = locations.Next
	conf.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
