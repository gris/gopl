// Ex4-2 imprime o SHA256 da entrada por default, ou exibe hashes SHA384 ou SHA512 baseado em flags
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hash = flag.Int("b", 256, "quantity of bits on the hash")

func main() {
	flag.Parse()
	input := flag.CommandLine.Arg(0)
	switch *hash {
	case 256:
		v := sha256.Sum256([]byte(input))
		fmt.Printf("%x\n", v)

	case 384:
		v := sha512.Sum384([]byte(input))
		fmt.Printf("%x\n", v)

	case 512:
		v := sha512.Sum512([]byte(input))
		fmt.Printf("%x\n", v)
	}
}
