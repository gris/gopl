// Netflag demonstra um tipo inteiro usado como um campo de bit.
package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // está ativo
	FlagBroadcast                      // suporta recurso de acesso broadcast
	FlagLoopback                       // é uma interface loopback
	FlagPointToPoint                   // pertence a um link ponto a ponto
	FlagMulticast                      // suporta recurso de acesso multicast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
