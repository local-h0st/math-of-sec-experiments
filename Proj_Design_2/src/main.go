package main

import (
	"Proj_Design_2/src/tool"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var id int64 = 57121117
	var id_reverse int64 = 71112175
	p, q := tool.Return2LargestPrime(id_reverse)
	prime_pair := tool.NUM_PAIR{p, q} // 得到两个大质数
	n := prime_pair.P * prime_pair.Q  // 两个质数乘积n
	phi_n := tool.Phi2(prime_pair)    // 得到φ(n)
	var e int64
	for true {
		e = int64(rand.Intn(int(phi_n-1)) + 1)
		tmp_pair := tool.NUM_PAIR{e, phi_n}
		if tmp_pair.Gcd() == 1 {
			break
		}
	} // 随机生成一个e，满足(e,φ(n)) = 1
	d := tool.GetReverseOfMod(e, phi_n) // 计算d，其中 ed ≡ 1 (mod φ(n))

	pk := tool.KEY{e, n}
	sk := tool.KEY{d, n}

	encrypted_data := tool.RSAEncrypt(id, pk)
	fmt.Println("original:", id)
	fmt.Println("encrypted with RSA:", encrypted_data)
	fmt.Println("decrypted:", tool.RSADecrypt(encrypted_data, sk))
	fmt.Println("decrypted with Chinese Remainder Theorem:", tool.RSADecryptSpeedUp(encrypted_data, sk, prime_pair))
	fmt.Println("public key (e, n):", pk.M, pk.N)
	fmt.Println("secret key (d, n):", sk.M, sk.N)
	fmt.Println(n, "=", prime_pair.P, "*", prime_pair.Q, "√")
	fmt.Println("φ(n) =", phi_n, "√")
	fmt.Println("e * d mod φ(n) =", (pk.M*sk.M)%phi_n+phi_n, "×") // 计算器可以算，结果为1，但是golang溢出了
	fmt.Println("e * d =", pk.M*sk.M, "×")                        // 普通方式直接溢出

	fmt.Println(tool.RSADecryptSpeedUp(468, tool.KEY{237, 667}, tool.NUM_PAIR{23, 29}))
}
