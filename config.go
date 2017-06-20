package file

import (
	"os"
	"log"
	"encoding/json"
	"strings"
)

type Item struct {
	Source string
	Destination string
	Size bool
	Date bool
	Version bool
}

type Variable struct {
	Key string
	Value string
}

type Configuration struct {
	Variables []Variable
	Items []Item
	WhatIf bool
}

/*
	Loops through variables wrapping variables in % tokens and expanding
	and variables within variables
*/
func resolveVariables(variables *[]Variable) {
	if variables != nil && len(*variables) == 0 {
		return
	}

	// wrap variables within % so we don't need
	// to keep doing for every test
	for i, item := range *variables {
		(*variables)[i].Key = "%" + item.Key + "%"
	}

	// now expand any variables within variables
	for i, item := range *variables {
		(*variables)[i].Key = resolveVariable(variables, item.Key, i)
		(*variables)[i].Value = resolveVariable(variables, item.Value, -1)
	}
}

func resolveVariable(variables *[]Variable, item string, variableIndex int) (result string) {

	result = item
	for i, v := range *variables {
		if variableIndex < 0 || variableIndex != i {
			result = strings.Replace(result, v.Key, v.Value, -1)
		}
	}
	return result
}

func resolveItem(variables *[]Variable, item *Item) {
	item.Source = resolveVariable(variables, item.Source, -1)
	item.Destination = resolveVariable(variables, item.Destination, -1)
}

func resolveItems(configuration *Configuration) {
	if configuration == nil || len(configuration.Variables) == 0 {
		return
	}

	// go through configuration injecting variable data
	for i := range configuration.Items {
		resolveItem(&configuration.Variables, &configuration.Items[i])
	}
}

func NewConfiguration(filename string) (*Configuration, error) {

	configuration := &Configuration {}

	fileHandle, _ := os.Open(filename)
	decoder := json.NewDecoder(fileHandle)
	err := decoder.Decode(configuration)
	if err != nil {
		log.Println(err)
	}

	resolveVariables(&configuration.Variables)
	resolveItems(configuration)

	return configuration, err
}
