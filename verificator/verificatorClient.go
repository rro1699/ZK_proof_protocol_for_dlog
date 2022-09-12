package verificator

import (
	"crypto/rand"
	"math/big"
)

var p, q, g, rStep1, r, y *big.Int

func SetOpenData(localP, localQ, localG, localY *big.Int) {
	p = localP
	q = localQ
	g = localG
	y = localY
}
func Step1(localR *big.Int) {
	rStep1 = localR
}

func Step2() *big.Int {
	var err error
	r, err = rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(err)
	}
	return r
}

func Step3(w *big.Int) bool {
	g.Exp(g, w, p)
	//fmt.Println("g^w mod p: ", g)
	y.Exp(y, r, p)
	//fmt.Println("y^r mod p: ", y)
	g.Mul(g, y)
	g.Mod(g, p)
	//fmt.Println("(g^w)(y^r) mod p: ", g)
	if rStep1.Cmp(g) == 0 {
		//fmt.Println("Okey")
		return true
	} else {
		//fmt.Println("R:", rStep1, " \\= ", g)
		//fmt.Println("not okey")
		return false
	}
}
