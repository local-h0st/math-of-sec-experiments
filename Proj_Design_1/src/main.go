package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var id_reverse int = 71112175
	var id int = 57121117
	fmt.Println("my id is", id, "and the reverse of it is", id_reverse)
	fmt.Print("two largest primes of ", id_reverse, " are: ")
	fmt.Println(return2LargestPrime(id_reverse))
	pair := NUM_PAIR{id, id_reverse}
	fmt.Println("the gcd of pair", pair.toStr(), "is", pair.gcd())
	fmt.Println("the lcm of pair", pair.toStr(), "is", pair.lcm())

}

func checkPrime(number int) bool {
	for i := 2; i < int(math.Sqrt(float64(number)))+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func return2LargestPrime(number int) (p1 int, p2 int) {
	i := number
	for {
		if checkPrime(i) {
			break
		}
		i--
	}
	p1 = i
	i--
	for {
		if checkPrime(i) {
			break
		}
		i--
	}
	p2 = i
	return
}

type NUM_PAIR struct {
	a int
	b int
}

func (pair *NUM_PAIR) gcd() int {
	var max, min int
	if pair.a >= pair.b {
		max = pair.a
		min = pair.b
	} else {
		max = pair.b
		min = pair.a
	}
	for {
		//for max >= min {
		//	max = max - min
		//}
		//if max == 0 {
		//	return min
		//} else {
		//	max, min = min, max
		//}
		if max%min == 0 {
			return min
		} else {
			max, min = min, max%min
		}
	}
}
func (pair *NUM_PAIR) lcm() int {
	return pair.a * pair.b / pair.gcd()
}
func (pair *NUM_PAIR) toStr() string {
	return "(" + strconv.Itoa(pair.a) + ", " + strconv.Itoa(pair.b) + ")"
}
