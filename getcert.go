package main

import (
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"os"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s server [port]\n", os.Args[0])
		os.Exit(1)
	}

	serverPort := "443"

	// Do we have an extra arg? If so, assume it's a port
	if len(os.Args) > 2 {
		serverPort = os.Args[2]
	}

	serverName := os.Args[1]
	serverAddr := serverName + ":" + serverPort
	fmt.Printf("Downloading certificate from %s (%s)\n", serverName, serverAddr)

	serverConn, err := tls.Dial("tcp4", serverAddr, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to %s! %q\n", serverAddr, err)
		os.Exit(1)
	}
	defer serverConn.Close()

	fmt.Printf("Connected\n")

	_, err = downloadServerCert(serverConn, serverName)
}
