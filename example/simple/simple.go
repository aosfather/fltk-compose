package main

import "github.com/aosfather/fltk-compose/compose"

func main() {
	compose.App(compose.Point(10, 10),
		compose.Size(400, 400),
		compose.Title("Hello compose world")).Run()
}
