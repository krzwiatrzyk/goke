package wallet

type account struct{
	name string
	wallet *wallet
	connections []account
}

func NewAccount(name string) *account { 
	return &account{name: name, wallet: newWallet()}
}

func(owner *account) TransferTo(itemName string, receiver *account) {
	receiver.wallet.objects[itemName] = owner.wallet.objects[itemName]
	delete(owner.wallet.objects, itemName)
}

func(a *account) GetItem(itemName string) *movie {
	return a.wallet.objects[itemName]
}

func (a *account) StoreItem(itemName string, item *movie) {
	a.wallet.objects[itemName] = item
}