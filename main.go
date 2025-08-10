package main

import "fmt"

// This is the interface definition. It's a contract with no code.
type HealthChecker interface {
    checkHealth() string
}

// This is the struct definition. It's a blueprint for our data.
type Server struct {
    Name string
}

// This is the method. It's the implementation of the interface's contract.
func (s Server) checkHealth() string {
    return "Server " + s.Name + " is healthy."
}

func main() {
    // We create a variable of type Server.
    webServer := Server{Name: "Web-Server-01"}
    
    // We can now call the method on our server variable.
    fmt.Println(webServer.checkHealth())
}