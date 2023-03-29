package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/myRpc"
	"log"
)

func init() {
	myRpc.SetupMyRpcDepositReply(depositReply)
	myRpc.SetupMyRpcWithdrawReply(withdrawReply)
	myRpc.SetupMyRpcBalanceReply(balanceReply)
	myRpc.SetupMyRpcTransferReply(transferReply)
}

func depositReply(message string) (myRpc.MyRpcProcedure) {
	log.Printf("Received deposit reply: " + message)
	return nil
}

func withdrawReply(message string) (myRpc.MyRpcProcedure) {
	log.Printf("Received withdraw reply: " + message)
	return nil
}

func balanceReply(balanceval uint64) (myRpc.MyRpcProcedure) {
	log.Printf("The Balance in the account is: %d", balanceval)
	return nil
}

func transferReply(message string) (myRpc.MyRpcProcedure) {
	log.Printf("The transfer status: " + message)
	return nil
}



func main () {

	altEthos.LogToDirectory("test/bankingclient")
	
	log.Println("********************* Entering Client ************************")

	var user = altEthos.GetUser()

	log.Println("The current user is " + user)

	log.Println("********************* Entering Deposit ************************")
	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	// RPC call for deposit
	call := myRpc.MyRpcDeposit{user, 60}
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	log.Println("********************* Exiting Deposit ************************")
	log.Println("********************* Entering Withdraw ************************")
	fd1, status1 := altEthos.IpcRepeat("myRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	// RPC call for withdraw
	call1 := myRpc.MyRpcWithdraw{user, 60}
	status1 = altEthos.ClientCall(fd1, &call1)
	if status1 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	log.Println("********************* Exiting Withdraw ************************")
	log.Println("********************* Entering Transfer ************************")
	fd3, status3 := altEthos.IpcRepeat("myRpc", "", nil)
	if status3 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status3)
		altEthos.Exit(status3)
	}

	// RPC call for transfer
	call3 := myRpc.MyRpcTransfer{user, "bennett", 10}
	status3 = altEthos.ClientCall(fd3, &call3)
	if status3 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status3)
		altEthos.Exit(status3)
	}

	log.Println("********************* Exiting Transfer ************************")
	log.Println("********************* Entering Balance ************************")
	fd2, status2 := altEthos.IpcRepeat("myRpc", "", nil)
	if status2 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	// RPC call for balance
	call2 := myRpc.MyRpcBalance{user}
	status2 = altEthos.ClientCall(fd2, &call2)
	if status2 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	log.Println("********************* Exiting Balance ************************")

	log.Println("********************* Exiting Client ************************")
}