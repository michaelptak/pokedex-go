package main

import (
	"errors"
	"fmt"
)

func commandMap(conf *config) error {
	locationsResp, err := conf.pokeApiClient.ListLocations(conf.Next)
	if err != nil {
		return err
	}

	conf.Next = locationsResp.Next
	conf.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	if conf.Previous == nil {
		return errors.New("Already on the first page.")
	}

	locationsResp, err := conf.pokeApiClient.ListLocations(conf.Previous)
	if err != nil {
		return err
	}

	conf.Next = locationsResp.Next
	conf.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
