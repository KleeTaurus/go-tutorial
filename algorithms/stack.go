package main

func greet(name string) {
	print("hello, " + name + "!\n")
	greet2(name)
	print("getting ready to say bye...\n")
	bye()
}

func greet2(name string) {
	print("how are you, " + name + "?\n")
}

func bye() {
	print("ok bye!\n")
}
