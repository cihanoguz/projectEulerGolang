package main

import "fmt"
/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
func sum1(x int)int{
	var a int
		a=0
	for i := 0; i < x; i++ {
		
		if (i%3 == 0)&&(i%5 ==0)  {
			a = a+i
		}else if i%3 == 0 {
			a= a+i
		}else if i%5 == 0 {
			a = a+i
		}else{
			a = a+0
		}
	}
	return a 
 }

func main()  {
	fmt.Println("enter a number:")
	var x int
	fmt.Scan(&x)
	fmt.Println(sum1(x))

	
}