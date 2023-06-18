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

func right(S, P byte) bool {
	return S&P == P
}

func valid(S byte) bool {
	return !(right(S, W) == right(S, G) && right(S, F) != right(S, W)) && !(right(S, G) == right(S, V) && right(S, F) != right(S, G))
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
			if nextValue := Sp.Value ^ F; valid(nextValue) {
				path := make([]byte, len(Sp.Path))
				copy(path, Sp.Path)
				Qn = append(Qn, &State{Value: nextValue, Path: append(path, nextValue)})
			}
			if right(Sp.Value, F) == right(Sp.Value, W) { // farmer and wolf in a side
				if nextValue := Sp.Value ^ (F | W); valid(nextValue) {
					path := make([]byte, len(Sp.Path))
					copy(path, Sp.Path)
					Qn = append(Qn, &State{Value: nextValue, Path: append(path, nextValue)})
				}
			}
			if right(Sp.Value, F) == right(Sp.Value, G) { // farmer and goat in a side
				if nextValue := Sp.Value ^ (F | G); valid(nextValue) {
					path := make([]byte, len(Sp.Path))
					copy(path, Sp.Path)
					Qn = append(Qn, &State{Value: nextValue, Path: append(path, nextValue)})
				}
			}
			if right(Sp.Value, F) == right(Sp.Value, V) { // farmer and vegetable in a side
				if nextValue := Sp.Value ^ (F | V); valid(nextValue) {
					path := make([]byte, len(Sp.Path))
					copy(path, Sp.Path)
					Qn = append(Qn, &State{Value: nextValue, Path: append(path, nextValue)})
				}
			}

		}
		Q = Qn
	}
}
