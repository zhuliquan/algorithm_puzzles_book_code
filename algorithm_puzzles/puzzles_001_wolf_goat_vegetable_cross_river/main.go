package main

import "fmt"

// using a byte value represent S,
// 1、F = S & 8
// 2、W = S & 4
// 3、G = S & 2
// 4、V = S & 1

// F(S)  => S ^ 8
// FW(S) => S ^ 12
// FG(S) => S ^ 10
// FV(S) => S ^ 9
const (
	V byte = 1 << iota
	G
	W
	F
)

func FT(S byte) byte {
	return S ^ F
}

func FWT(S byte) byte {
	return S ^ (F | W)
}

func FGT(S byte) byte {
	return S ^ (F | G)
}

func FVT(S byte) byte {
	return S ^ (F | V)
}

func printSide(b bool) string {
	if b {
		return "R"
	} else {
		return "L"
	}
}

type State struct {
	Value byte
	Path  []byte
}

func printPath(path []byte) {
	for _, s := range path {
		printState(s)
	}
}

func printState(S byte) {
	fmt.Printf("Farmer=%s,Wolf=%s,Goat=%s,Vegetable=%s\n",
		printSide(S&F == F),
		printSide(S&W == W),
		printSide(S&G == G),
		printSide(S&V == V),
	)
}

func main() {
	R := &State{Value: 0, Path: []byte{0}}
	Q := []*State{R}
	for len(Q) != 0 {
		Qn := []*State{}

		for _, Sp := range Q {
			if Sp.Value == F|W|G|V {
				fmt.Println("find solve:")
				printPath(Sp.Path)
				return
			}

			for _, FF := range []func(byte) byte{
				FT, FWT, FGT, FVT,
			} {
				if Sn := FF(Sp.Value); !(Sn&W == Sn&G && Sn&F != Sn&W) && !(Sn&G == Sn&V && Sn&F != Sn&G) {
					path := make([]byte, len(Sp.Path))
					copy(path, Sp.Path)
					children := &State{Value: Sn, Path: append(path, Sn)}
					Qn = append(Qn, children)
				}

			}

		}
		Q = Qn
	}
}
