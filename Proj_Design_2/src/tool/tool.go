package tool

import (
	"math"
	"math/big"
)

func checkPrime(number int64) bool {
	var i int64 = 2
	for ; i < int64(math.Sqrt(float64(number)))+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func Return2LargestPrime(number int64) (p1 int64, p2 int64) {
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
	P int64
	Q int64
}

func (pair *NUM_PAIR) Gcd() int64 {
	var max, min int64
	if pair.P >= pair.Q {
		max = pair.P
		min = pair.Q
	} else {
		max = pair.Q
		min = pair.P
	}
	for {
		if max%min == 0 {
			return min
		} else {
			max, min = min, max%min
		}
	}
}
func (pair *NUM_PAIR) Lcm() int64 {
	return pair.P * pair.Q / pair.Gcd()
}

// 欧拉函数
func Phi(n int64) int64 {
	if checkPrime(n) {
		return n - 1
	} else {
		var count int64
		var m int64 = 1
		for ; m < n; m++ {
			pair := NUM_PAIR{m, n}
			if pair.Gcd() == 1 {
				count++
			}
		}
		return count
	}

} // 暴力枚举
func Phi2(pair NUM_PAIR) int64 { // n表示为两个互素数的乘积
	if pair.Gcd() != 1 {
		return -1
	} else {
		return Phi(pair.Q) * Phi(pair.P)
	}
}

func GetReverseOfMod(n, mod int64) int64 {
	var s, ss, t, tt, q, r, rr, rrr int64
	rrr = mod
	rr = n
	ss = 1
	tt = 0
	s = 0
	t = 1
	for true {
		q = rrr / rr
		r = rrr % rr
		if r != 0 {
			ss, s = s, ss-s*q
			tt, t = t, tt-t*q
			rrr, rr, r = rr, r, 0
		} else {
			break
		}
	}
	if t < 0 {
		t = t + mod
	}
	return t
} // 广义欧几里得除法求n关于mod的逆元

func GetPrimeByIndex(index int) int64 {
	count := 0
	var num int64 = 2
	for count != index {
		if checkPrime(num) {
			count++
		}
		num++
	}
	return num - 1
}

func getPrimeFactors(mod int64) []int64 {
	// 首先对mod质因数分解
	prime_factors := make([]int64, 0)
	tmp := mod
	for true {
		for i := 1; GetPrimeByIndex(i) <= tmp; i++ {
			if tmp%GetPrimeByIndex(i) == 0 {
				prime_factors = append(prime_factors, GetPrimeByIndex(i))
				tmp = tmp / GetPrimeByIndex(i)
				break
			}
		}
		if tmp == 1 {
			break
		}
	}
	return prime_factors
}
func checkBinaryEqualsToDecimo(binary_reverse []int, decimo int) bool {
	power := func(i int) int {
		result := 1
		for j := 0; j < i; j++ {
			result *= 2
		}
		return result
	}
	for i := 0; i < len(binary_reverse); i++ {
		decimo -= binary_reverse[i] * power(i)
	}
	if decimo == 0 {
		return true
	} else {
		return false
	}
}

type KEY struct {
	M int64
	N int64
}

// 由于公钥的n是很大的，而且不能被暴力质因数分解，因此不能用中国剩余定理求解，所以考虑用模重复平方法计算
func RSAEncrypt(data int64, pk KEY) int64 {
	// find out the result: data^pk.M mod pk.N
	// data < pk.N 不然会出问题
	var e = pk.M
	e_binary_reverse := make([]int, 0)
	for e != 0 {
		//e_binary = append([]int{e % 2}, e_binary...)
		e_binary_reverse = append(e_binary_reverse, int(e%2))
		e = e / 2
	}
	// e_binary_reverse 正确获得

	// 开始模重复平方计算法
	var a = big.NewInt(1)
	var b = big.NewInt(int64(data))
	for i := 0; i < len(e_binary_reverse); i++ {
		//fmt.Println("[debug] a/b", a, b)
		if e_binary_reverse[i] == 0 {
			//a = a % pk.N
			a = big.NewInt(0).Mod(a, big.NewInt(int64(pk.N)))
		} else {
			//a = (a * b) % pk.N
			a = big.NewInt(0).Mod(big.NewInt(0).Mul(a, b), big.NewInt(int64(pk.N)))
		}
		//b = (b * b) % pk.N
		b = big.NewInt(0).Mod(big.NewInt(0).Mul(b, b), big.NewInt(int64(pk.N)))
	}
	// 最终的a就是结果
	var encrypted int64 = a.Int64()
	if encrypted < 0 {
		encrypted += int64(pk.N)
	}
	return encrypted

}
func RSADecrypt(encrypted int64, sk KEY) int64 {
	// encrypted^pk.M mod pk.N 这个结构和加密过程完全一样，故直接调用就行
	return RSAEncrypt(encrypted, sk)
}
func RSADecryptSpeedUp(encrypted int64, sk KEY, prime_pair NUM_PAIR) int64 {
	// 中国剩余定理加速解密
	// 本质还是求解 encrypted^sk.m mod sk.n = ? 但是由于pq是自己生成的，因此直到n如何分解
	m1 := big.NewInt(prime_pair.P)
	m2 := big.NewInt(prime_pair.Q)
	// 欧拉定理(2.4.1)
	i1 := big.NewInt(0).Mod(big.NewInt(sk.M), big.NewInt(0).Sub(m1, big.NewInt(1)))
	i2 := big.NewInt(0).Mod(big.NewInt(sk.M), big.NewInt(0).Sub(m2, big.NewInt(1)))
	b1 := big.NewInt(RSAEncrypt(encrypted, KEY{i1.Int64(), m1.Int64()})) // 仅仅是这个函数刚好提供模重复平方法
	b2 := big.NewInt(RSAEncrypt(encrypted, KEY{i2.Int64(), m2.Int64()}))
	M1 := m2
	M2 := m1
	M11 := GetReverseOfMod(M1.Int64(), m1.Int64())
	M22 := GetReverseOfMod(M2.Int64(), m2.Int64())
	//result := (b1*M1*M11 + b2*M2*M22) % sk.N
	result := big.NewInt(0).Mod(
		big.NewInt(0).Add(
			big.NewInt(0).Mul(big.NewInt(0).Mul(b1, M1), big.NewInt(M11)),
			big.NewInt(0).Mul(big.NewInt(0).Mul(b2, M2), big.NewInt(M22)),
		),
		big.NewInt(sk.N),
	)
	return result.Int64()
}
