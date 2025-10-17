package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"
)

func TestWallet(t *testing.T) {
	wallets := Wallets{}
	err := wallets.LoadFromFile()
	if err != nil {
		panic(err)
	}

	t.Log(wallets)
	for _, w := range wallets.Wallets {
		t.Log(string(w.GetAddress()))
	}
}

func TestSignature(t *testing.T) {
	wallets := Wallets{}
	err := wallets.LoadFromFile()
	if err != nil {
		panic(err)
	}

	addresses := wallets.GetAddresses()
	if len(addresses) < 1 {
		panic("no address in the wallet")
	}

	wallet := wallets.GetWallet(addresses[0])

	bs := []byte{0x01, 0x03, 0x64}
	r, s, err := ecdsa.Sign(rand.Reader, &wallet.PrivateKey, bs)
	if err != nil {
		panic(err)
	}

	signature := append(r.Bytes(), s.Bytes()...)


	lenS := len(signature)

	r = new(big.Int).SetBytes(signature[:lenS/2])
	s = new(big.Int).SetBytes(signature[lenS/2:])

	pubKey, _ := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), wallet.PublicKey)

	result := ecdsa.Verify(pubKey, bs, r, s)
	t.Log(result)
}
