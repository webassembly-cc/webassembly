package main

import (
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()

	n.UseHandler(initRouter())

	n.Run(":9000")

}
