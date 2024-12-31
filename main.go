package main

func main() {
	if err := StartServer(); err != nil {
		panic(err)
	}
}
