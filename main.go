package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func domainCheck(domain string) {
	var hasMX, hasDMARC, hasSPF bool
	var dmarcRecord, spfRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	textRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	}

	for _, record := range textRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v | %v | %v | %v | %v | %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain | hasMX | hasSPF | sprRecord | hasDMARC | dmarcRecord\n")

	for scanner.Scan() {
		domainCheck(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v\n", err)
	}
}
