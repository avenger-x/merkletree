module github.com/authur117/merkletree

go 1.23.0

toolchain go1.24.2

require (
	github.com/ethereum/go-ethereum v1.15.11
	golang.org/x/crypto v0.35.0
)

require (
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	golang.org/x/sys v0.30.0 // indirect
)

replace github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef => github.com/ebellocchia/go-bip39 v0.1.0
