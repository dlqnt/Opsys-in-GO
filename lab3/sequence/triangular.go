package sequence

// Task 1: Triangular numbers
//
// triangular(n) returns the n-th Triangular number, and is defined by the
// recurrence relation F_n = n + F_n-1 where F_0=0 and F_1=1
//
// Visualization of numbers:
// n = 1:    n = 2:     n = 3:      n = 4:    etc...
//   o         o          o           o
//            o o        o o         o o
//                      o o o       o o o
//                                 o o o o
func triangular(n uint) uint {

	var y uint
	var x uint = 0
	for y = uint(1); y <= n; y++ {
		x = y + x

	}

	/* 	fmt.Println("n =", x)
	   	for i = uint(1); i <= x; i++ {
	   		for j = uint(1); j <= x-i; j++ {
	   			fmt.Print(" ")
	   		}
	   		for symbol = uint(1); symbol <= i; symbol++ {
	   			fmt.Print("o ")
	   		}
	   		fmt.Print("\n")

	   	} */
	return x
}
