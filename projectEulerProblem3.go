
/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

/*  this code works on int type input. I couldnt find how can i change for loop iteration in Golang. Because the number is int64

*/


package main

import "fmt"

import "math"
func isPrime(value int) bool {
    for i := 2; i <= int(math.Sqrt(float64(value)/2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}
/*func pressPrime(value int)int{
	var prime int
	prime = 1
	if isPrime(value) == true{
		prime = value;

	}
	return prime
}
*/

func largestPrimeFactor(n int)int{
	var largestPrimeFact int
	largestPrimeFact = n
	var t float64

	
	t = float64(n)
	var p int = int(math.Sqrt(t))
	for i := 2; i <= p; i++ { 
		if isPrime(i)== true {
			for n%i == 0 {
				n=n/i;
				largestPrimeFact = i
			}
		
		}
		if n > largestPrimeFact {
			largestPrimeFact=n
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
