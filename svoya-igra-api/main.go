package main

import "github.com/muhammadjon1304/svoya-igra-api/app"

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}
