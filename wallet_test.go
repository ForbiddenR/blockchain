package main

import "testing"

func TestWallet(t *testing.T) {
	wallets := Wallets{}
	err := wallets.LoadFromFile()
	if err != nil {
		panic(err)
	}

	for _, w := range wallets.Wallets {
		t.Log(string(w.GetAddress()))
	}
}
