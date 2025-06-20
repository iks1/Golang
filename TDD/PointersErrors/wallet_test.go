package wallet

import (
	"testing"
)


func TestWallet(t *testing.T){

	assertError := func(t testing.TB, got error, want error){
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assetBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()

		if got!= want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assetBalance(t, wallet, want)
	})

    t.Run("withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertNoError(t, err)
		assetBalance(t, wallet, want)
	})	

	t.Run("withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
        assertError(t, err, ErrInsufficientFunds)
		assetBalance(t, wallet, startingBalance)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}



// func TestingSome(t *testing.T){

// 	assertSomething1 := func(t testing.TB){
// 		//something
// 	}

// 	t.Run("someting", func(t *testing.T){
// 		assetSomething1(t)
// 		assetSomething2(t)
// 	})

// 	assertSomething2 := func(t testing.TB){
// 		//something
// 	}
// }