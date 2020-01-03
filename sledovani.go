package main

import "fmt"
import "math/big"

const LIMIT = 50

func pairs(n *big.Int) (p *big.Int) {
	switch {
	case n.Cmp(big.NewInt(0)) == 0:
		fallthrough
	case n.Cmp(big.NewInt(1)) == 0:
		return big.NewInt(0)
	case n.Cmp(big.NewInt(2)) == 0:
		return big.NewInt(2)
	}
	p = big.NewInt(2)
	var previousWoman, woman bool = true, false
	var fibA, fibP *big.Int
	fibA, fibP = big.NewInt(3), big.NewInt(2)
	for i := int64(4); i <= n.Int64(); i++ {
		woman = big.NewInt(0).Mod(fibA, big.NewInt(2)).Cmp(big.NewInt(0)) == 0
		if woman {
			if gcd(fibA, fibP) {
				p.Add(p, big.NewInt(1))
			}
		} else if previousWoman {
			if gcd(fibA, fibP) {
				p.Add(p, big.NewInt(1))
			}
		}
		previousWoman = woman
		swp := &big.Int{}
		swp.Add(fibA, fibP)
		fibP = fibA
		fibA = swp
	}
	return
}

func gcd(x, y *big.Int) bool {
	var z *big.Int
	z = &big.Int{}
	z.GCD(nil, nil, x, y)
	if z.Cmp(big.NewInt(1)) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		var M big.Int
		fmt.Scan(&M)
		if M.Uint64() > LIMIT {
			/*
				for j := i; j < N; j++ {
					fmt.Println(0)
				}
			*/
			return
		}
		fmt.Println(pairs(&M))
	}
}
