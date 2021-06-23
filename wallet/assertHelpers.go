package wallet

import "testing"

func AssertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func AssertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but got nil")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
