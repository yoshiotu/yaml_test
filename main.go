package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	obj := yaml.MapSlice{}

	// let's read some YAML!

	in, err := ioutil.ReadFile("testdata/sample.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(in, &obj)
	if err != nil {
		panic(err)
	}

	// let's change the password value
	objNew := yaml.MapSlice{}
	for _, item := range obj {
		s, ok := item.Key.(string)
		if ok && s == "password" {
			item.Value = "my_little_secret"
		}
		objNew = append(objNew, item)
	}
	obj = objNew

	// let's add a team tag
	objNew = yaml.MapSlice{}
	for _, item := range obj {
		s, ok := item.Key.(string)
		if ok && s == "team_tags" {
			v, ok2 := item.Value.(yaml.MapSlice)
			if !ok2 {
				panic("failed to parse team tags")
			}
			item.Value = append(v, yaml.MapItem{Key: "team_location", Value: "usa"})
		}
		objNew = append(objNew, item)
	}
	obj = objNew

	// now let's append new properties!

	// start off easy, just a simple string.

	// some_key: some_value
	obj = append(obj, yaml.MapItem{Key: "some_key", Value: "some_value"})

	// getting fancier, let's append a property that is a map/dictionary

	// some_other_key:
	//     map_key_1: map_value_1
	//     map_key_2: map_value_2
	obj = append(obj,
		yaml.MapItem{
			Key: "some_other_key",
			Value: []yaml.MapItem{
				{Key: "map_key_1", Value: "map_value_1"},
				{Key: "map_key_2", Value: "map_value_2"},
			},
		},
	)

	// now let's append a property that's an array of strings

	// array_key:
	//    - array_item_1
	//    - array_item_2
	obj = append(obj, yaml.MapItem{Key: "array_key", Value: []string{"array_item_1", "array_item_2"}})
	for _, item := range obj {
		fmt.Printf("%v %+v\n", item.Key, item.Value)
		fmt.Printf("the value is a %T\n", item.Value)
	}

	fmt.Println()

	// finally, let's print the final YAML!
	out, err := yaml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
