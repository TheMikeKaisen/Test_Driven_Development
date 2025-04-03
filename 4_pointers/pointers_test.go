package pointers

import (
	"fmt"
	"testing"
)


func TestPointers(t *testing.T) {
	t.Run("testing banking systems", func(t *testing.T) {

		w := Wallet{}


		w.Deposit(Bitcoin(10))

		got := w.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		fmt.Println()


	})

	t.Run("withdraw btc", func (t *testing.T){

		w := Wallet{
			balance: 20,
		}

		fmt.Println(w.balance)

		err := w.Withdraw(20);
		if err!=nil {
			fmt.Println("An error has occured ", err)
		}


		got := w.Balance()
		want := Bitcoin(0)

		if(got != want ) {
			t.Errorf("got %d, want %d", got, want);
		}
		
		fmt.Println("after withdraw ", w.balance)

	})
}