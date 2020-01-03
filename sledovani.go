package main

import "fmt"
import "math/big"

const LIMIT = 1e18

func pairs(n *big.Int) (p *big.Int) {
	switch {
	case n.Cmp(big.NewInt(0)) == 0:
		fallthrough
	case n.Cmp(big.NewInt(1)) == 0:
		return big.NewInt(0)
	case n.Cmp(big.NewInt(2)) == 0:
		return big.NewInt(2)
	}
	p = &big.Int{}
	p.Div(n, big.NewInt(3))
	p.Mul(p, big.NewInt(2))
	p.Add(p, big.NewInt(1))
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
			for j := i; j < N; j++ {
				fmt.Println(0)
			}
			return
		}
		fmt.Println(pairs(&M))
	}
}
