package main

import (
	web "github.com/arifinnor/echo-app/routes"
)

func main() {
	e := web.Web()

	e.Logger.Fatal(e.Start(":1323"))
}
