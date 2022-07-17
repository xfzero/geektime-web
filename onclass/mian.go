package main

func main() {
	server := NewHttpServer("test-server")

	server.Route("/", home)
	server.Route("/login", login)
}
