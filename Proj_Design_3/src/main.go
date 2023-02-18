package main

import (
	"fmt"
)

func main() {

	g1 := map[string]string{"A": "10", "2": "Q", "3": "A", "4": "5", "5": "2", "6": "6", "7": "K", "8": "3", "9": "8", "10": "9", "J": "7", "Q": "J", "K": "4"}
	g2 := map[string]string{"A": "3", "2": "6", "3": "7", "4": "A", "5": "10", "6": "K", "7": "4", "8": "2", "9": "J", "10": "Q", "J": "9", "Q": "5", "K": "8"}
	g3 := map[string]string{"A": "4", "2": "A", "3": "3", "4": "K", "5": "J", "6": "10", "7": "6", "8": "Q", "9": "8", "10": "9", "J": "5", "Q": "7", "K": "2"}
	poker := map[string]string{"A": "A", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "10": "10", "J": "J", "Q": "Q", "K": "K"}
	printOrder("original order :", poker)
	// 轮换法
	multi(g1, poker)
	multi(g2, poker)
	for i := 1; i <= 7; i++ {
		multi(g3, poker)
	}
	printOrder("G1(1 time) G2(1 time) G3(7 times) :", poker)
	for i := 1; i <= 5; i++ {
		multi(g3, poker)
	}
	for i := 1; i <= 11; i++ {
		multi(g2, poker)
	}
	for i := 1; i <= 34; i++ {
		multi(g1, poker)
	}
	printOrder("resume ", poker)
	// 逆运算法
	g1_rev := make(map[string]string)
	reverseGroup(g1, g1_rev)
	g2_rev := make(map[string]string)
	reverseGroup(g2, g2_rev)
	g3_rev := make(map[string]string)
	reverseGroup(g3, g3_rev)
	multi(g1, poker)
	multi(g2, poker)
	for i := 1; i <= 7; i++ {
		multi(g3, poker)
	}
	printOrder("G1(1 time) G2(1 time) G3(7 times) :", poker)
	for i := 1; i <= 7; i++ {
		multi(g3_rev, poker)
	}
	multi(g2_rev, poker)
	multi(g1_rev, poker)
	printOrder("resume with rev", poker)

}

func multi(group, poker map[string]string) {
	// 必定是群右乘
	for k, v := range poker {
		poker[k] = group[v]
	}
}

func reverseGroup(group, reversed map[string]string) {
	for k, v := range group {
		reversed[v] = k
	}
}

func printOrder(info string, g map[string]string) {
	// 输出当前牌山顺序
	fmt.Println(info)
	fmt.Println(g["A"], g["2"], g["3"], g["4"], g["5"], g["6"], g["7"], g["8"], g["9"], g["10"], g["J"], g["Q"], g["K"])
}
