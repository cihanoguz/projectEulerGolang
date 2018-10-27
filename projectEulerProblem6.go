/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"

func difference(n int)int{
	var total1 int
	total1=0
	
	
       for i:=1; i<=n; i++ {
		   total1= total1+i*i
	   }
	   return ((n*(n+1))/2)*((n*(n+1))/2) - total1

}


func main(){
	fmt.Println("give me number")
	var n int
	fmt.Scan(&n)
	fmt.Println(difference(n))



}