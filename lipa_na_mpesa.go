package daraja_wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const Lipa_na_mpesa_url = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

type LipeNaMpesaPayStruct struct {
	BusinessShortCode, Password, Timestamp                      string
	TransactionType, Amount, PartyA, PartyB                     string
	PhoneNumber, CallBackURL, AccountReference, TransactionDesc string
}
type LipaNaMpesaResp map[string]interface{}

// `LipaNaMpesaPayment` -
// 'Lipa na M-Pesa Online Payment'
// Lipa na M-Pesa Online Payment API is used to initiate a M-Pesa
// transaction on behalf of a customer using STK Push. This is the same technique mySafaricom App
// uses whenever the app is used to make payments.
func (a *Auth) LipaNaMpesaPayment(l *LipeNaMpesaPayStruct) (LipaNaMpesaResp, error) {
	var daraja_resp LipaNaMpesaResp
	token, err := a.GetAuthKey()
	if err != nil {
		return daraja_resp, err
	}

	client := &http.Client{}

	saf_req, err := json.Marshal(l)
	if err != nil {
		return daraja_resp, err
	}
	url := a.setUrl(Lipa_na_mpesa_url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(saf_req))
	if err != nil {
		return daraja_resp, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)
	if err != nil {
		return daraja_resp, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return daraja_resp, err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return daraja_resp, err
	}
	// var token Token
	// err = json.Unmarshal(body, &token)
	// if err != nil {
	// 	return token, err
	// }
	return daraja_resp, nil
}
