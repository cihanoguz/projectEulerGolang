
/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

import "math"

func largestPrimeFactor(n int) int{
	var largestPrimeFact int
	largestPrimeFact = 0
	var t float64
	t = float64(n)
	var p int = int(math.Sqrt(t))
	for i := 2; i <= p; i++ { 
        if (n % i == 0) { 
            largestPrimeFact = i; 
		} 
	} 
	var k int
	k = n / largestPrimeFact
	var l float64
	l = float64(k)
	for i := 2; i <= int(math.Sqrt(l)); i++{
		if int(k) % i !=0 {
			largestPrimeFact=k
		}

    }

	
	
	return largestPrimeFact


}
func main(){
	fmt.Println("enter a number:")
	var n int
	fmt.Scan(&n)
	fmt.Println(largestPrimeFactor(n))



}
