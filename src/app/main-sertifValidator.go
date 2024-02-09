package main

import "sertif_validator/app/routes"

func main() {

	s := routes.ServiceVALIDATOR()

	s.Start(":9070")
}
