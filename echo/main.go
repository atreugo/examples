package main

func main() {
	r := Route{
		Endpoint: "/e/{echo}",
	}

	s := NewAtreugoServer("0.0.0.0:8007")
	s.RegisterRoute(r)
	s.Run()
}
