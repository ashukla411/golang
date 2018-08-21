package main 
import "fmt" 
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func vals() (int, int) {
    return 3, 7
}
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    // We call `intSeq`, assigning the result (a function)
    // to `nextInt`. This function value captures its
    // own `i` value, which will be updated each time
    // we call `nextInt`.
    nextInt := intSeq()

    // See the effect of the closure by calling `nextInt`
    // a few times.
	fmt.Print("next sequence int: ")
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // To confirm that the state is unique to that
    // particular function, create and test a new one.
    newInts := intSeq()
    fmt.Println(newInts())
	fmt.Println("new func")
    one,two  := vals()

    fmt.Println(one)
    fmt.Println(two)
    _, three := vals()
    fmt.Println(three)

    sum(1, 2)
    sum(1, 2, 3)
	
	fmt.Println("Checking false and 0 and other string stuff below....")
	const a bool = false
	if(!a){
		fmt.Println("not false is true")
}
	var e int
	fmt.Println(e)
    s := make([]bool, 3)
	fmt.Println(s,len(s))
    b := [3]string{"0","b","c"}
	fmt.Println(b)
	c := make([]string, 3)
	c[0] = "a"
	c = append(c,"e","f","g","h","i","asdasdasd")
	fmt.Println(c)
	d := make([]string,len(c))
	copy(d,c)
	fmt.Println(d)
	t := []string{"g", "h", "i"}
	t = append(t,"z")
	fmt.Println(t)
	}