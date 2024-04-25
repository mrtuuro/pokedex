package main

func main() {

    cfg := NewConfig()
	server := NewServer(cfg)
	server.Start()

}
