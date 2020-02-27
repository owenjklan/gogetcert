package main

import (
	"flag"
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"os"
	"strconv"
)

func printCertInfo(cert x509.Certificate) {
	fmt.Printf("Valid From: %q\n", cert.NotBefore)
	fmt.Printf("Expiry:     %q\n", cert.NotAfter)
	fmt.Printf("Subject Alt. Names:\n")
	
	if len(cert.DNSNames) > 0 {
		for index, subAltName := range cert.DNSNames {
			fmt.Printf("\t\033[1;33m#%2d\033[0m %s\n", index, subAltName)
		}
	} else {
		fmt.Printf(" \t\033[33m- None -\033[0m\n")
	}
	fmt.Printf("\n")
}

func saveCertToPEM(cert x509.Certificate, serverName string,  certNumber int) {
	// encode in PEM format
	outFilePath := fmt.Sprintf("%s_%02d", serverName, certNumber)
	outFile, err := os.OpenFile(outFilePath, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed opening output file '%s'! %q", outFilePath, err)
		return
	}
	defer outFile.Close()

	block := &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
	if err := pem.Encode(outFile, block); err != nil {
		fmt.Fprintf(os.Stderr, "Failed encoding certificate in PEM form! %q\n", err)
		return
	}

	outFile.Sync()
}

func downloadServerCert(conn *tls.Conn, serverName string) (*x509.Certificate, error) {
	connState := conn.ConnectionState()

	fmt.Printf("%d certificates in chain\n", len(connState.PeerCertificates))

	for index, cert := range connState.PeerCertificates {
		fmt.Printf("Cert #%2d\n", index)
		fmt.Printf("--------\n")
		printCertInfo(*cert)
		saveCertToPEM(*cert, serverName, index)
		fmt.Printf("\n")
	}

	return nil, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] server\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	insecureCerts := flag.Bool("insecure", false, "Don't validate certificates")
	serverPort := flag.Int("port", 443, "Port number to connect on")

	if len(os.Args) < 2 {
		usage()
	}

	flag.Parse()

	serverPortStr := strconv.FormatInt(int64(*serverPort), 10)

	if len(flag.Args()) < 1 {
		usage()
	}
	serverName := flag.Args()[0]
	serverAddr := serverName + ":" + serverPortStr

	fmt.Printf("Downloading certificate from %s (%s)\n", serverName, serverAddr)

	skipVerifyCerts := *insecureCerts
	if skipVerifyCerts {
		fmt.Fprintf(os.Stderr, "\033[1;31m!!! Not validating certificates !!!\033[0m\n")
	}

	tlsConfig := tls.Config{InsecureSkipVerify: skipVerifyCerts}
	serverConn, err := tls.Dial("tcp4", serverAddr, &tlsConfig)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to %s! %q\n", serverAddr, err)
		os.Exit(1)
	}
	defer serverConn.Close()

	fmt.Printf("Connected\n")

	_, err = downloadServerCert(serverConn, serverName)
}
