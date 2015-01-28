package main

import "github.com/MikeFitzgerald/domains/domains"

/*
Create a new MongoDB session, using a database
named "domains". Create a new server using
that session, then begin listening for HTTP requests.
*/
func main() {
	session := domains.NewSession("domains")
	server := domains.NewServer(session)
	server.Run()
}
