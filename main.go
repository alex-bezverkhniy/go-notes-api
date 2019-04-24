package main

func main() {
	a := App{}
	// TODO: move args into enveronments
	a.Initialize("gonotes", "1Q2w3e4r", "gonotes")
	a.Run(":8080")
}
