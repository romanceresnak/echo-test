package main

func main() {
	RunServer()
}

func RunServer() {
	echo := NewApp()
	echo.Start(":3001")
}
