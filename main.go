package main

import (
	"fmt"
	"time"
)

type temp struct {
	Name string
	Age  int8
}

type tuple func(index int) any

func newTuple(values ...any) tuple {
	return func(index int) any {
		// Passing -1 returns the total length
		if index == -1 {
			return len(values)
		}

		if index < 0 || index >= len(values) {
			panic(fmt.Sprintf("runtime error: index out of range [%d]", index))
		}
		return values[index]
	}
}

func (t tuple) Len() int {
	return t(-1).(int) // from any to int
}

func (t tuple) Slice() []any {
	length := t.Len()
	result := make([]any, length) // pre-allocate memory for speed

	for i := range length {
		result[i] = t(i)
	}
	return result
}

func main() {

	tmp := &temp{
		Name: "bob",
		Age:  22,
	}

	// Initialize the tuple
	t := newTuple(1000, "hello", "pear", time.Now().UTC(), tmp)

	// Read values
	fmt.Println(t(0))              // Outputs: 1
	fmt.Println(t(1))              // Outputs: hello
	fmt.Println(t(4).(*temp).Name) // Outputs: bob
	// fmt.Println(t(20000))          // Panics with index out of range error
	// could use recover() depending on situation

	// t[0] = 99 -> Compile error! Cannot overwrite the data.

	stuff := make([]any, 10)
	fmt.Printf("### capacity of stuff: %d\n", cap(stuff))
	for i := range 10 {
		stuff[i] = fmt.Sprintf("num %d", i)
	}
	combined := []any{}
	combined = append(combined, stuff...)
	combined = append(combined, t.Slice()...)

	fmt.Println("Together is fun...")
	for i, v := range combined {
		fmt.Printf("data [%d] [%v]\n", i, v)
	}
}
