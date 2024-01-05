package main

import (
	"main/cmd"
	"net/http"
)

func main() {
	cmd.InitDB()                                       // DB Connection
	http.ListenAndServe(":3000", cmd.NewHttpHandler()) // Server And Route
}
