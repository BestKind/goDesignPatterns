package facade

import (
	"log"
	"testing"
)

func TestFacade(t *testing.T) {
	walletFacade := newWalletFacade("test", 1234)
	err := walletFacade.addMoneyToWallet("test", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s \n", err.Error())
	}
	err = walletFacade.deductMoneyFromWallet("test", 1234, 3)
	if err != nil {
		log.Fatalf("Error: %s \n", err.Error())
	}
}
