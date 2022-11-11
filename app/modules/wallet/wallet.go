package wallet

type wallet struct{
	objects map[string]*movie
}

func newWallet() *wallet{
	return &wallet{objects: make(map[string]*movie)}
}

func (w *wallet) Size() (size int){
	return len(w.objects)
}

