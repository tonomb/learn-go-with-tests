package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat( t * testing.T){

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected{
		t.Errorf(" expected %q to be %q", repeated, expected)
	}
}


func BenchmarkRepeat(b * testing.B){
	for i:= 0 ; i < b.N ; i++{
		Repeat("a", 100)
	}
}


func ExampleRepeat(){
	repeat := Repeat("a",5)
	fmt.Println(repeat)
	// Output: "aaaaa"
}