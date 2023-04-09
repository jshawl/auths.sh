package main

type Session struct {
	Id string
}

type User struct {
	Name      string
	PublicKey string
}

func main() {
	HTTPServe()
	SSHServe()
}
