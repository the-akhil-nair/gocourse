package advanced

import (
	"fmt"
	"reflect"
)

// Working with Methods using Reflection
type greeter struct{}

func (g greeter) Greet(fname, lname string) string {
	return "Hello, " + fname + " " + lname
}

func (greeter) Farewell(name string) string {
	return "Goodbye, " + name
}

func main() {
	g := greeter{}
	v := reflect.ValueOf(g)
	t := reflect.TypeOf(g)
	fmt.Println("Type:", t)

	for i := range t.NumMethod() {
		method := t.Method(i)
		fmt.Println("Method Name:", method.Name)
		fmt.Println("Method Type:", method.Type)
	}
	// To find and call a method by name
	method := v.MethodByName("Greet")
	// []string{"World", "Everyone"}
	// []type{typeOf("string"), typeOf("string")}
	args := []reflect.Value{reflect.ValueOf("John"), reflect.ValueOf("Doe")}
	result := method.Call(args)
	fmt.Println(result[0].Interface())
	fmt.Println(result[0].String())

	j := map[string]any{
		"name": "Alice",
		"age":  30,
		"school": map[string]any{
			"name":    "XYZ High School",
			"address": "123 Main St",
			"year":    2024,
		},
		"grades": []int{90, 85, 88},
	}

	for key, value := range j {
		v := reflect.ValueOf(value)
		// Type: map[string]interface {}, Kind: map
		// Type: []int, Kind: slice
		fmt.Printf("Key: %s, Value: %v, Type: %s, Kind: %s\n", key, value, v.Type(), v.Kind())

		if v.Type() == reflect.TypeOf(map[string]any{}) {
			// Need to convert back to map[string]any to iterate as we cannot iterate reflect.Value directly
			for k, i := range v.Interface().(map[string]any) {
				fmt.Println("Map Key:", k, "Map Value:", i)
			}
		}

		if v.Kind() == reflect.Slice {
			for i, _ := range v.Interface().([]int) {
				fmt.Println("Slice Index:", i, "Slice Value:", v.Index(i).Interface())
			}
		}
	}

}

// Working with Structs and Fields using Reflection
// type Person struct {
// 	Name string // For reflection, fields must be exported
// 	Age  int
// }

// func main() {
// 	p := Person{Name: "Alice", Age: 30}
// 	v := reflect.ValueOf(p)

// 	for i := range v.NumField() {
// 		vField := v.Field(i)
// 		typeField := v.Field(i)
// 		fmt.Println("Field Name:", typeField.Type().Name())
// 		fmt.Println("Field Value:", vField)
// 	}

// 	v1 := reflect.ValueOf(&p).Elem() // Get the addressable value
// 	nameField := v1.FieldByName("Name")

// 	if nameField.CanSet() {
// 		nameField.SetString("Bob")
// 	}

// 	for i := 0; i < v1.NumField(); i++ {
// 		vField := v1.Field(i)
// 		typeField := v1.Type().Field(i)
// 		fmt.Println("Field Name:", typeField.Name)
// 		fmt.Println("Field Value:", vField)
// 	}
// }

// func main() {
// 	x := 42
// 	v := reflect.ValueOf(x)

// 	// If you want to convert the reflect.Value back to int
// 	p := v.Interface().(int) // type assertion to int
// 	fmt.Println(p)

// 	fmt.Println("Value:", v)
// 	fmt.Println("Kind:", v.Kind())
// 	fmt.Println("Type:", v.Type())
// 	fmt.Println("IsValid:", v.IsValid())
// 	fmt.Println("IsZero:", v.IsZero())
// 	fmt.Println("Interface:", v.Interface())
// 	fmt.Println("Is Int:", v.Kind() == reflect.Int)

// 	y := 10
// 	// The following line will panic because v is not addressable or an inrerface
// 	// panic: reflect: call of reflect.Value.Elem on int Value
// 	// v = reflect.ValueOf(y).Elem()
// 	v = reflect.ValueOf(&y).Elem()
// 	// This will give pointer to y
// 	v1 := reflect.ValueOf(&y)

// 	fmt.Println("v1:", *v1.Interface().(*int))
// 	fmt.Println("v1 Kind:", v1.Kind())
// 	fmt.Println("v1 Type:", v1.Type())
// 	fmt.Println("v in int:", v.Int())
// 	v.SetInt(100)
// 	fmt.Println("Modified y:", y)

// 	var itf interface{} = "hello"
// 	v = reflect.ValueOf(itf)
// 	fmt.Println("itf Value:", v)
// 	fmt.Println("itf Kind:", v.Kind())
// 	fmt.Println("itf Type:", v.Type())
// 	fmt.Println("itf Interface:", v.Interface())
// }
