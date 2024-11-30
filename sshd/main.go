package main

import (
	"log"
	"os/exec"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
	
	   cmd := exec.Command("/bin/rc")
	   cmd.Stdin = s
	   cmd.Stdout = s
	   cmd.Stderr = s
	   
	   if err := cmd.Start(); err != nil {
	       log.Println("Failed to start command:", err)
	       return
	   }
	   
	   if err:= cmd.Wait(); err != nil {
	       log.Println("Command failed:", err)
	   }
	})

	log.Println("starting ssh server on port 22...")
	log.Fatal(ssh.ListenAndServe(":22", nil))
}