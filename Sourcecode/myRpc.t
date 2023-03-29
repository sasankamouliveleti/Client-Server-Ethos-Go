MyRpc interface {
	Deposit(user string, amount uint64) (message string)
	Withdraw(user string, amount uint64) (message string)
	Balance(user string) (bal uint64)
	Transfer(from string, to string, amount uint64) (message string)
}
