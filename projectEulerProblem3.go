  /*
  The prime factors of 13195 are 5, 7, 13 and 29.

  What is the largest prime factor of the number 600851475143 ?
  */


  package main

  import "fmt"

  import "math"
  func isPrime(value int64) bool {
      var i int64
      for i := 2; i <= int64(math.Sqrt(float64(value)/2)); i++ {
          if value%i == 0 {
              return false
          }
      }
      return value > 1
  }

  func largestPrimeFactor(n int64)int64{
    var largestPrimeFact int64
    largestPrimeFact = n
    var t float64

    
    t = float64(n)
    var p int64 = int64(math.Sqrt(t))
    var i int64
    for i = 2; i <= p; i++ { 
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
      var n int64
      fmt.Scan(&n)
      fmt.Println(largestPrimeFactor(n))

  }
