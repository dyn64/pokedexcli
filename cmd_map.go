package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeMap struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(conf *config) error {
	res, err := http.Get(conf.mapNext)
	if err != nil {
		return fmt.Errorf("http GET error from url: %s \nError: %v", conf.mapNext, err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed. Status code: %d\nBody: %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("io.readall error: %v", err)
	}
	pokMap := pokeMap{}
	err = json.Unmarshal(body, &pokMap)
	if err != nil {
		return fmt.Errorf("Unmarshal failed with error: %v", err)
	}
	for loc := range pokMap.Results {
		fmt.Println(pokMap.Results[loc].Name)
	}
	if pokMap.Previous == nil {
		conf.mapBack = conf.mapBase
	} else {
		conf.mapBack = *pokMap.Previous
	}
	conf.mapNext = pokMap.Next
	return nil
}

func commandMapBack(conf *config) error {
	if conf.mapBack == conf.mapBase {
		fmt.Println("you're on the first page")
		return nil
	}
	url := conf.mapBack
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http GET error from url: %s \nError: %v", url, err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed. Status code: %d\nBody: %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("io.readall error: %v", err)
	}
	pokMap := pokeMap{}
	err = json.Unmarshal(body, &pokMap)
	if err != nil {
		return fmt.Errorf("Unmarshal failed with error: %v", err)
	}
	for loc := range pokMap.Results {
		fmt.Println(pokMap.Results[loc].Name)
	}
	if pokMap.Previous == nil {
		conf.mapBack = conf.mapBase
	} else {
		conf.mapBack = *pokMap.Previous
	}
	conf.mapNext = pokMap.Next
	return nil

}
