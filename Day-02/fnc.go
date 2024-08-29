package main

import "fmt"

type test struct {
	name string 
	age int
}
func (test *test) changeName (newName string) {
	test.name = newName;
}
func sum ( a,b int) int  {
	return a + b
}

func sums_array (nums ...int)int {
	value := 0

	for _, num := range nums{
		value = value + num
	}
	return value
}
func main() {
	nums := []int{1, 2, 3, 4, 5}


	fmt.Println("func");
	fmt.Println(sum(3,3));
	yanuar := test{name:"yanuar",age:30}
	fmt.Println(yanuar)
	yanuar.changeName("yanuar tampan")
	fmt.Println(yanuar)
	fmt.Println(sums_array(nums...))
	// sums_array(nums);
}