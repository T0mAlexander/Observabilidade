// Monitoramento de certificados SSL

package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SSLInfo struct {
	Certificate CertsInfo `json:"Certificate"`
}

type CertsInfo struct {
	DNS                string `json:"DNS"`
	Issuer             string `json:"Issuer"`
	Expiration         string `json:"Expiration"`
	PubKeyAlgorithm    string `json:"PubKeyAlgorithm"`
	SignatureAlgorithm string `json:"SignatureAlgorithm"`
}

func main() {
	targetURL := "https://google.com.br"
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	response, error := client.Get(targetURL)
	if error != nil {
		fmt.Println("Erro:", error)
		return
	}

	defer response.Body.Close()

	Cert := response.TLS.PeerCertificates[0]

	CertInfo := CertsInfo{
		DNS:                response.TLS.ServerName,
		Issuer:             Cert.Issuer.Organization[0],
		Expiration:         Cert.NotAfter.Format("02 de Janeiro de 2006 - 15:04h"),
		PubKeyAlgorithm:    Cert.PublicKeyAlgorithm.String(),
		SignatureAlgorithm: Cert.SignatureAlgorithm.String(),
	}

	sslInfo := SSLInfo{
		Certificate: CertInfo,
	}

	sslInfoJSON, error := json.MarshalIndent(sslInfo, "", "  ")
	if error != nil {
		fmt.Println("Erro:", error)
		return
	}

	fmt.Println(string(sslInfoJSON))
}
