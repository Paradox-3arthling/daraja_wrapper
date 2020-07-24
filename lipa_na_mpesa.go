package daraja_wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const Lipa_na_url = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

type LipeNaMpesaPayStruct struct {
	BusinessShortCode, Password, Timestamp                      string
	TransactionType, Amount, PartyA, PartyB                     string
	PhoneNumber, CallBackURL, AccountReference, TransactionDesc string
}

// `LipaNaMpesaPayment` -
// 'Lipa na M-Pesa Online Payment'
// Lipa na M-Pesa Online Payment API is used to initiate a M-Pesa
// transaction on behalf of a customer using STK Push. This is the same technique mySafaricom App
// uses whenever the app is used to make payments.
func (a *Auth) LipaNaMpesaPayment(l *LipeNaMpesaPayStruct) (bool, error) {
	token, err := a.GetAuthKey()
	if err != nil {
		return false, err
	}
	client := &http.Client{}
	var url string
	if a.Prod {
		url = strings.Replace(Lipa_na_url, "sandbox.", "", 1)
	} else {
		url = Lipa_na_url
	}
	saf_req, err := json.Marshal(l)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(saf_req))
	if err != nil {
		return false, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	// var token Token
	// err = json.Unmarshal(body, &token)
	// if err != nil {
	// 	return token, err
	// }
	return false, nil
}
