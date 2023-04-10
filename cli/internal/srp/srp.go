package srp

import "math/big"

type SRP struct {
	safePrime *big.Int
	generator *big.Int
	k         *big.Int
}
