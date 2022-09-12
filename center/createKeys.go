package center

import (
	vC "awesomeProject/Verificator"
	pC "awesomeProject/prover"
	"crypto/rand"
	"math/big"
)

var p, q, g *big.Int

func fact(n *big.Int) *big.Int {
	i := big.NewInt(2)
	tmp := new(big.Int)
	tmp1 := new(big.Int)
	res := new(big.Int)
	for tmp1.Mul(i, i).Cmp(n) < 0 {
		for tmp.Mod(n, i).Cmp(big.NewInt(0)) == 0 {
			n.Div(n, i)
			//fmt.Print(i, " ")
			res = i
		}
		i.Add(i, big.NewInt(1))
	}
	if n.Cmp(big.NewInt(1)) > 0 {
		//fmt.Println(n)
		res = n
	}
	return res
}

func GeneratedKeys() bool {
	var bitSize = 32
	var err error
	p, err = rand.Prime(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}
	p1 := new(big.Int)
	p1.Sub(p, big.NewInt(1))
	q = fact(p1)
	var flag = true
	for flag {
		g, err = rand.Int(rand.Reader, q)
		if err != nil {
			panic(err)
		}
		if len(g.Text(10)) == len(q.Text(10)) {
			flag = false
		}
	}
	return StartAlgoritm(bitSize)
}

func StartAlgoritm(bitSize int) bool {
	pC.SetOpenData(p, q, g)
	y := pC.GenerateProverKeys(bitSize) //get public key y

	vC.SetOpenData(p, q, g, y)

	R := pC.Step1() // prover generate R and send him to Verificator
	vC.Step1(R)

	r := vC.Step2()
	pC.Step2(r)

	w := pC.Step3()
	return vC.Step3(w)
}
