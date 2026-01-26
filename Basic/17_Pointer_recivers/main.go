package main

import "fmt"

type Wallet struct {
    balance int
}

// === Pointer Receiver ===
// (w *Wallet) means this method gets a POINTER to the original Wallet.
func (w *Wallet) Deposit(amount int) {
    // Go automatically de-references the pointer,
    // so you can just write w.balance instead of (*w).balance
    w.balance += amount // This modifies the ORIGINAL wallet
    fmt.Printf("  -> Inside Deposit (pointer): Balance is now %d\n", w.balance)
}

func main() {
    myWallet := Wallet{balance: 100}
    
    fmt.Println("--- Calling Pointer Receiver (Deposit) ---")
    fmt.Printf("Original balance (before): %d\n", myWallet.balance)
    
    myWallet.Deposit(50) // Call the method
    
    fmt.Printf("Original balance (after):  %d (Changed!)\n", myWallet.balance)
}
