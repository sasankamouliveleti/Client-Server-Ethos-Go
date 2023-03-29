package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/myRpc"
	"log"
	"strconv"
)

var accountsArr [5]Account

type Account struct {
	ID      uint64
	Name    string
	Balance uint64
}

func init() {
	myRpc.SetupMyRpcDeposit(deposit)
	myRpc.SetupMyRpcWithdraw(withdraw)
	myRpc.SetupMyRpcBalance(balance)
	myRpc.SetupMyRpcTransfer(transfer)
}

// RPC method for deposit
func deposit(user string, amount uint64) (myRpc.MyRpcProcedure) {
	var responsereply string
	responsereply = "Invalid deposit"
	for i := 0; i < len(accountsArr); i++ {
		if accountsArr[i].Name == user {
			accountsArr[i].Balance = accountsArr[i].Balance + amount
			responsereply = "The deposit of " + strconv.Itoa(int(amount)) + " is done and the balance is " + strconv.Itoa(int(accountsArr[i].Balance))
		}
	}
	return &myRpc.MyRpcDepositReply{responsereply}
}

// RPC method for withdraw
func withdraw(user string, amount uint64) (myRpc.MyRpcProcedure) {
	var responsereply1 string
	responsereply1 = "Invalid withdrawal"
	for i := 0; i < len(accountsArr); i++ {
		if accountsArr[i].Name == user {
			if amount <= accountsArr[i].Balance {
				accountsArr[i].Balance = accountsArr[i].Balance - amount
				responsereply1 = "The withdrwal of " + strconv.Itoa(int(amount)) + " is done and the balance is " + strconv.Itoa(int(accountsArr[i].Balance))
			}
		}
	}
	return &myRpc.MyRpcWithdrawReply{responsereply1} 
}

// RPC method for balance
func balance(user string) (myRpc.MyRpcProcedure){
	var responsereply2 uint64
	responsereply2 = 0
	for i := 0; i < len(accountsArr); i++ {
		if accountsArr[i].Name == user {
			return &myRpc.MyRpcBalanceReply{accountsArr[i].Balance} 
		}
	}
	return &myRpc.MyRpcBalanceReply{responsereply2} 
}

// RPC method for transfer
func transfer(from string, to string, amount uint64) (myRpc.MyRpcProcedure) {
	var responsereply3 string
	responsereply3 = "The transfer could not be made"
	for i := 0; i < len(accountsArr); i++ {
		if accountsArr[i].Name == from {
			if amount <= accountsArr[i].Balance {
				accountsArr[i].Balance = accountsArr[i].Balance - amount
				for j := 0; j < len(accountsArr); j++ {
					if accountsArr[j].Name == to{
						accountsArr[j].Balance = accountsArr[j].Balance + amount
						responsereply3 = "The amount of " + strconv.Itoa(int(amount)) + " has been transfered to " + to + " from " + from 
						return &myRpc.MyRpcTransferReply{responsereply3} 
					}
				}
			}
		}
	}
	return &myRpc.MyRpcTransferReply{responsereply3} 
}


func main () {
	altEthos.LogToDirectory("test/bankingserver")

	// Initialise accounts
	accountsArr[0] = Account{ID: 1, Name: "me", Balance: 1000}
	accountsArr[1] = Account{ID: 2, Name: "jlong", Balance: 500}
	accountsArr[2] = Account{ID: 3, Name: "mouli", Balance: 750}
	accountsArr[3] = Account{ID: 4, Name: "gabriel", Balance: 150}
	accountsArr[4] = Account{ID: 5, Name: "bennett", Balance: 300}

	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection is established")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}
}