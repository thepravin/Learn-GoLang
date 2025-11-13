package main

import "fmt"

type Wallet struct {
    balance int
}

// === Value Receiver ===
// (w Wallet) means this method gets a COPY of the Wallet.
func (w Wallet) TryDeposit(amount int) {
    w.balance += amount // This only modifies the copy
    fmt.Printf("  -> Inside TryDeposit (copy): Balance is now %d\n", w.balance)
}

// This method just reads data, so a value receiver is perfect.
func (w Wallet) GetBalance() int {
    // w is a copy, but we are just reading from it.
    return w.balance
}

func main() {
    myWallet := Wallet{balance: 100}
    
    fmt.Println("--- Calling Value Receiver (TryDeposit) ---")
    fmt.Printf("Original balance (before): %d\n", myWallet.balance)
    
    myWallet.TryDeposit(50) // Call the method
    
    fmt.Printf("Original balance (after):  %d (Unchanged!)\n", myWallet.balance)
    
    fmt.Printf("\nGetBalance (value receiver): %d\n", myWallet.GetBalance())
}
