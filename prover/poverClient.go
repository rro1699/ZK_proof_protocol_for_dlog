package prover

import (
	"crypto/rand"
	"math/big"
)

var a *big.Int //secret key

var p, q, g, k, r *big.Int

func SetOpenData(localP, localQ, localG *big.Int) {
	p = localP
	q = localQ
	g = localG
}

func GenerateProverKeys(bitSize int) *big.Int {
	var err error
	a, err = rand.Int(rand.Reader, q)
	if err != nil {
		panic(err)
	}
	//fmt.Println("a: ", a)
	a.Neg(a)
	//fmt.Println("a: ", a)
	y := new(big.Int)
	y.Exp(g, a, p)
	a.Neg(a)
	/*	fmt.Println("p: ", p)
		fmt.Println("q: ", q)
		fmt.Println("g: ", g)
		fmt.Println("y: ", y)
		fmt.Println("a: ", a)*/
	return y
}

func Step1() *big.Int {
	var err error
	k, err = rand.Int(rand.Reader, q)
	if err != nil {
		panic(err)
	}
	R := new(big.Int)
	R.Exp(g, k, p)
	//fmt.Println("Prover: k: ", k, " R: ", R)
	return R
}

func Step2(localR *big.Int) {
	r = localR
	//fmt.Println("Prover: r: ", r)
}

func Step3() *big.Int {
	if r.Cmp(big.NewInt(0)) == 0 {
		//fmt.Println("Prover: w: ", k)
		return k
	} else {
		k.Add(k, a)
		k.Mod(k, p)
		//fmt.Println("Prover: w: ", k)
		return k
	}
}
