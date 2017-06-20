
package file

import (
	"testing"
	"log"
	//"strconv"
	"strconv"
)

func TestResolveItem_WithValidVariable_ExpectItemChanges(t *testing.T) {

	variables := []Variable {
		{
			Key : "key1", Value : "value1",
		},
	}

	item := Item {
		Source : "c:\\src\\%key1%\\subdir",
		Destination : "c:\\dest\\%key1%",
	}

	resolveVariables(&variables)
	resolveItem(&variables, &item)

	if item.Source != "c:\\src\\value1\\subdir" {
		t.Fatal("Variable not assigned to source " + item.Source)
	}
	if item.Destination != "c:\\dest\\value1" {
		t.Fatal("Variable not assigned to destination " + item.Destination)
	}
}

func TestResolveItem_WithValidVariableWithinValidVariable_ExpectItemChanges(t *testing.T) {
	variables := []Variable {
		{
			Key : "key1", Value : "value1",
		},
	}

	item := Item {
		Source : "c:\\src\\%key1%\\subdir",
		Destination : "c:\\dest\\%key1%",
	}

	resolveVariables(&variables)
	resolveItem(&variables, &item)

	if item.Source != "c:\\src\\value1\\subdir" {
		t.Fatal("Variable not assigned to source " + item.Source)
	}
	if item.Destination != "c:\\dest\\value1" {
		t.Fatal("Variable not assigned to destination " + item.Destination)
	}
}

func TestResolveVariable_WithSingleKeyValue_ExpectKeyReplaceWithValue(t *testing.T) {
	variables := []Variable {
		{
			Key : "key1", Value : "value1",
		},
		{
			Key : "key2", Value : "value2",
		},
		{
			Key : "key3", Value : "value3",
		},
	}

	resolveVariables(&variables)
	result := resolveVariable(&variables, "c:\\root\\%key1%\\subdir", -1)

	if result != "c:\\root\\value1\\subdir" {
		t.Fatal("Variable not resolved " + result)
	}
}

func TestResolveItems_NilConfiguration_ExpectNoFailures(t *testing.T) {
	resolveItems(nil)
}

func TestResolveItems_EmptyConfiguration_ExpectNoFailures(t *testing.T) {
	configuration := &Configuration {}

	resolveItems(configuration)
}

func TestResolveItems_WithValidVariable_ExpectConfigurationChanges(t *testing.T) {
	configuration := &Configuration {}
	configuration.Variables = []Variable {
		{
			Key : "key1", Value : "value1",
		},
	}
	configuration.Items = []Item {
		{
			Source : "c:\\src\\%key1%\\subdir",
			Destination : "c:\\dest\\%key1%",
		},
	}

	resolveVariables(&configuration.Variables)
	resolveItems(configuration)

	for _, item := range configuration.Items {
		log.Println(item.Source)
	}

	if configuration.Items[0].Source != "c:\\src\\value1\\subdir" {
		t.Fatal("Variable not assigned to source " + configuration.Items[0].Source )
	}
	if configuration.Items[0].Destination != "c:\\dest\\value1" {
		t.Fatal("Variable not assigned to destination " + configuration.Items[0].Destination)
	}
}

func TestResolveVariables_WithoutVariablesWithinVariables_ExpectOnlyKeyChanges(t *testing.T) {
	variables := []Variable{
		{
			Key: "key1", Value: "value1",
		},
		{
			Key: "key2", Value: "value2",
		},
		{
			Key: "key3", Value: "value3",
		},
	}

	resolveVariables(&variables)

	for i, v := range variables {
		tokenized := "%key" + strconv.Itoa(i+1) + "%"
		if v.Key != tokenized {
			t.Fatal("Variable key not wrapped in tokens " + v.Key + " != " + tokenized)
		}
	}
}
