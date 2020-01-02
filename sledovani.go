package main

import "fmt"
import "math/big"

const LIMIT = 50

func pairs(n *big.Int) (p *big.Int) {
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(0)
	}
	p = &big.Int{}
	fibs := make([]*big.Int, n.Uint64())
	for i := range fibs[2:] {
		fibs[i+2] = big.NewInt(0)
	}
	fibs[0] = big.NewInt(1)
	fibs[1] = big.NewInt(1)
	for i := range fibs[2:] {
		fibs[i+2].Add(fibs[i+1], fibs[i])
	}
	for i, fib := range fibs {
		for _, friend := range fibs[i+1:] {
			IsManFriend := big.NewInt(0).Mod(friend, big.NewInt(2)).Cmp(big.NewInt(1)) == 0
			IsManFib := big.NewInt(0).Mod(fib, big.NewInt(2)).Cmp(big.NewInt(1)) == 0
			if IsManFib == IsManFriend {
				continue
			}
			if isFib(big.NewInt(0).Add(fib, friend)) {
				if gcd(fib, friend) {
					continue
				} else {
					p.Add(p, big.NewInt(1))
				}
			}
		}
	}
	return
}

func isFib(n *big.Int) bool {
	var ques1, ques2 *big.Float
	ques1 = &big.Float{}
	ques2 = &big.Float{}
	ques1.Mul(big.NewFloat(5.0), big.NewFloat(0.0).Mul(big.NewFloat(0.0).SetInt(n), big.NewFloat(0.0).SetInt(n)))
	ques1.Add(ques1, big.NewFloat(4.0))
	ques1.Sqrt(ques1)
	ques2.Mul(big.NewFloat(5.0), big.NewFloat(0).Mul(big.NewFloat(0).SetInt(n), big.NewFloat(0).SetInt(n)))
	ques2.Sub(ques2, big.NewFloat(4.0))
	ques2.Sqrt(ques2)
	if ques1.IsInt() || ques2.IsInt() {
		return true
	}
	return false
}

func gcd(x, y *big.Int) bool {
	var z big.Int
	z.GCD(nil, nil, x, y)
	if z.Cmp(big.NewInt(1)) == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		var M big.Int
		fmt.Scan(&M)
		if M.Uint64() > LIMIT {
			for j := i; j < N; j++ {
				fmt.Println(0)
			}
			return
		}
		fmt.Println(pairs(&M))
	}
}
