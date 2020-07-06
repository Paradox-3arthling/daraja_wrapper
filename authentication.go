package daraja_wrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
		return token, err
	}
	fmt.Println(token)
	return token, nil
}
func (a *Auth) setUrl() string {
	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	if a.Prod {
		url = "https://safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	}
	return url
}

// func makeClient() (*http.Client, error) {

// }
