package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	MailTo = os.Getenv("MAILTO")
	url    = "https://api-app.qtrac.com/scheduler-template/api/external/companies/156068c8-1a37-4dcb-ad6b-82ab7691898f/scheduler-templates/branches/4ef1665c-8724-4cce-933b-8b4e3698bdbc/workflows/c2f6d68e-77a1-4bfd-8604-b482b024f3bd/services/d46583a3-714a-4554-8c80-a9f22f5f64f0/first-available"
)

func main() {
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("error getting url: %s\n", err)
			continue
		}
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error reading body: %s\n", err)
			continue
		}
		if string(buf) == "[]" {
			log.Println("No appointment")
			continue
		}
		err = ioutil.WriteFile("email.txt", buf, 0644)
		if err != nil {
			log.Printf("couldn't write file: %s\n", err)
			continue
		}
		err = exec.Command("sendmail", MailTo, "<", "email.txt").Run()
		if err != nil {
			log.Printf("error running command: %s\n", err)
			continue
		}
		log.Println("Mail sent!")

	}
}
