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
	// bufio.NewScanner creates a scanner that reads input line by line.
	// os.Stdin means it reads from the terminal/input pipe.
	//
	// Example:
	// go run main.go
	// then type:
	// google.com
	// yahoo.com
	scanner := bufio.NewScanner(os.Stdin)

	// Print the CSV header.
	// This tells us what each column means.
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	// scanner.Scan() keeps returning true as long as there is another line to read.
	// When input ends, it returns false and the loop stops.
	for scanner.Scan() {

		// scanner.Text() gives us the current line as a string.
		// strings.TrimSpace removes spaces, tabs, and newline characters around the domain.
		domain := strings.TrimSpace(scanner.Text())

		// If the line is empty, skip it.
		if domain == "" {
			continue
		}

		// Check DNS records for this domain and print the result.
		checkDomain(domain)
	}

	// If the scanner stopped because of an error, print the error and exit.
	// For example: if input could not be read.
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

// checkDomain checks one domain for MX, SPF, and DMARC DNS records.
// After checking, it prints one line of output.
func checkDomain(domain string) {

	// These variables store the results we will print later.
	//
	// hasMX      = true if domain has MX records
	// hasSPF     = true if domain has an SPF TXT record
	// hasDMARC   = true if domain has a DMARC TXT record
	//
	// spfRecord   = the actual SPF record text, if found
	// dmarcRecord = the actual DMARC record text, if found
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// ---------------------------------------------------------------
	// 1. Check MX records
	// ---------------------------------------------------------------
	//
	// net.LookupMX asks DNS for MX records for the given domain.
	//
	// Example:
	// If domain is "gmail.com", DNS will return MX records that point
	// to Google's mail servers.
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		// If the lookup fails, log the error.
		// We do not stop the program because we still want to check other domains.
		log.Printf("Error looking up MX for %s: %v\n", domain, err)
	}

	// If the MX record list has at least one item, then the domain has MX records.
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// ---------------------------------------------------------------
	// 2. Check SPF record
	// ---------------------------------------------------------------
	//
	// SPF records are stored as TXT records on the domain itself.
	//
	// Example:
	// TXT record for example.com might contain:
	// "v=spf1 include:_spf.google.com ~all"
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		// If TXT lookup fails, log the error and continue.
		log.Printf("Error looking up TXT for %s: %v\n", domain, err)
	}

	// A domain can have many TXT records.
	// We loop through each one and look for the SPF record.
	for _, record := range txtRecords {

		// SPF records start with "v=spf1".
		// strings.HasPrefix checks whether the record begins with that text.
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record

			// We already found the SPF record, so stop checking other TXT records.
			break
		}
	}

	// ---------------------------------------------------------------
	// 3. Check DMARC record
	// ---------------------------------------------------------------
	//
	// DMARC is not checked directly on "example.com".
	// It is checked on a special subdomain:
	// "_dmarc.example.com"
	//
	// So if the domain is "example.com", we look up:
	// "_dmarc.example.com"
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		// If DMARC lookup fails, log the error and continue.
		log.Printf("Error looking up DMARC for %s: %v\n", domain, err)
	}

	// Loop through the TXT records found at "_dmarc.<domain>".
	for _, record := range dmarcRecords {

		// DMARC records start with "v=DMARC1".
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record

			// Found the DMARC record, no need to check the rest.
			break
		}
	}

	// ---------------------------------------------------------------
	// 4. Print the final result
	// ---------------------------------------------------------------
	//
	// This prints one CSV-style line for the domain.
	//
	// %v means "print the value using its default format".
	// \n adds a newline so each domain result appears on its own line.
	fmt.Printf("%v, %v, %v, %v, %v, %v\n",
		domain,
		hasMX,
		hasSPF,
		spfRecord,
		hasDMARC,
		dmarcRecord,
	)
}

// This program reads domain names from standard input, one domain per line.
// For each domain, it checks:
//
//  1. Whether the domain has MX records.
//     MX records are DNS records that say where email for the domain is handled.
//
//  2. Whether the domain has an SPF record.
//     SPF is a TXT record that starts with "v=spf1".
//
//  3. Whether the domain has a DMARC record.
//     DMARC is a TXT record located at "_dmarc.yourdomain.com"
//     and usually starts with "v=DMARC1".
//
// The output is printed in a CSV-like format:
// domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
