package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Records")

	for scanner.Scan() {
		checkEmail(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: could not read from input %v\n", err)
	}
}

func checkEmail(email string) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		fmt.Printf("Invalid email format: %s\n", email)
		return
	}

	domain := parts[1]
	checkDomain(domain, email)
}

func checkDomain(domain string, email string) {
	var hasMX, hasSPF, hasDMRC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dm := "_dmarc." + domain
	dmarcRecords, err := net.LookupTXT(dm)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMRC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v\n", email, hasMX, hasSPF, spfRecord, hasDMRC, dmarcRecord)
}
