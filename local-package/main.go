package main

import (
	"github.com/KleeTaurus/go-tutorial/local-package/pack1"
	"github.com/KleeTaurus/go-tutorial/local-package/pack1/subpack1"
)

func main() {
	pack1.SayHello()
	subpack1.SayCustomHello()
	// pack2.sayGoodBye()
}
