package main

import "sertif_validator/app/routes"

func main() {

	s := routes.ServiceTKBAI()

	s.Start(":9070")
}
