package main

func main() {
	var server = Server{}
	server.ssl = []string{"cert.pem", "key.pem"}
	server.Start("huy", "", "mdpad")
}
