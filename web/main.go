package main

import (
	"os"
	"simpleWebServer/web/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "9999"
)

func main() {
	//get PORT number from environment variables
	port := os.Getenv("PORT")
	//if PORT number not set in enviroment variables then get it default value 8080 as defined above
	if len(port) == 0 {
		port = PORT
	}
	//get input port using flag
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	//use function of negroni Run, Run takes an addr string identical to http.ListenAndServe
	n := service.NewServer()
	n.Run(":" + port)
}
