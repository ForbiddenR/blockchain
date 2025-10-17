package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	err := wallets.LoadFromFile()
	if err != nil {
		return &wallets, err
	}
	return &wallets, nil
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := string(wallet.GetAddress())
	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); err != nil  {
		if  !os.IsNotExist(err) {
			return err
		}else {
			return nil
		}
	}

	fileContent, err := os.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var serializableWallets SerializableWallets
	gob.Register(SerializableWallet{})
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&serializableWallets)
	if err != nil {
		log.Panic(err)
	}
	wallets, err := serializableWallets.ToWallets()
	if err != nil {
		return nil
	}
	ws.Wallets = wallets.Wallets
	return nil
}

func (ws Wallets) ToSerializable() SerializableWallets {
	sws := make(map[string]*SerializableWallet, len(ws.Wallets))
	for k, v := range ws.Wallets {
		sws[k] = v.ToSerializable()
	}
	return SerializableWallets{sws}
}

type SerializableWallets struct {
	SerializableWallets map[string]*SerializableWallet
}

func (sws SerializableWallets) ToWallets() (*Wallets, error) {
	wallets := Wallets{make(map[string]*Wallet, len(sws.SerializableWallets))}
	for k, v := range sws.SerializableWallets {
		w, err := v.ToWallet()
		if err == nil {
			wallets.Wallets[k] = w
		}
	}
	return &wallets, nil
}

func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(SerializableWallet{})

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws.ToSerializable())
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
