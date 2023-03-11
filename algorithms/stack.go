package algorithms

func Greet(name string) {
	print("hello, " + name + "!\n")
	Greet2(name)
	print("getting ready to say bye...\n")
	Bye()
}

func Greet2(name string) {
	print("how are you, " + name + "?\n")
}

func Bye() {
	print("ok bye!\n")
}
