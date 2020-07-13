package daraja_wrapper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Token - This is the token received from the
// daraja API. The `Token` struct is used to validate
// various actions done on the API
type Token struct {
	Token  string `json:"access_token"`
	Expiry string `json:"expires_in"`
}

func (t Token) String() string {
	return fmt.Sprintf("Token: '%s', Expires in(seconds): '%s'", t.Token, t.Expiry)
}

type Auth struct {
	Key, Secret string
	Prod        bool
}

func (a Auth) String() string {
	return fmt.Sprintf("Key: '%s', Secret: '%s', Prod: '%v'", a.Key, a.Secret, a.Prod)
}
func (a *Auth) GetAuthKey() (Token, error) {
	if a.Secret == "" || a.Key == "" {
		no_args := fmt.Errorf("The `Auth` struct is missing the `key` or `secret` field! \n Auth: '%v'", a)
		return Token{}, no_args
	}
	client := &http.Client{}
	url := a.setUrl()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Token{}, err
	}
	req.SetBasicAuth(a.Key, a.Secret)
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return Token{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Token{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Token{}, err
	}
	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("json.Unmarshal :", token)
		return token, err
	}
	return token, nil
}
func (a *Auth) setUrl() string {
	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	if a.Prod {
		url = "https://safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	}
	return url
}

// `SecurityCredentials` - This function is for generating the `SecurityCredentials`
// used in the parameters for mpesa API
// `password` is the 'Initiator Security Password'
// `cert_file` is the contents of the cetificate(Leaving the getting of the file to the user)
func SecurityCredentials(password, cert_file []byte) (string, error) {
	block, _ := pem.Decode(cert_file)
	if block == nil {
		return "", fmt.Errorf("There is no PEM data found!")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}
	public := cert.PublicKey.(*rsa.PublicKey)
	security_cred, err := rsa.EncryptPKCS1v15(rand.Reader, public, password)
	security_cred_b64 := b64.StdEncoding.EncodeToString([]byte(security_cred))
	return security_cred_b64, nil
}
