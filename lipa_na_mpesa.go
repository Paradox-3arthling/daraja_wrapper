package daraja_wrapper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Lipa_na_mpesa_url = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

type LipaNaMpesaPayStruct struct {
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
func (l *LipaNaMpesaPayStruct) LipaNaMpesaPayment(a *Auth, pass_key string) (LipaNaMpesaResp, error) {
	l.TransactionType = "CustomerPayBillOnline"
	l.GenPasswordAndAssign(pass_key)
	var daraja_resp LipaNaMpesaResp

	token, err := a.GetAuthKey()
	if err != nil {
		return daraja_resp, err
	}

	saf_req, err := json.Marshal(l)
	if err != nil {
		return daraja_resp, fmt.Errorf("`json.Marshal/1` got the err -> '%s' with args -> '%s'", err, l)
	}
	url := a.setUrl(Lipa_na_mpesa_url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(saf_req))
	if err != nil {
		return daraja_resp, fmt.Errorf("`http.NewRequest/3` got the err -> '%s'.", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return daraja_resp, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return daraja_resp, err
	}
	err = json.Unmarshal(body, &daraja_resp)
	if err != nil {
		return daraja_resp, err
	}
	return daraja_resp, nil
}
func (l *LipaNaMpesaPayStruct) GenPasswordAndAssign(pass_key string) {
	var msg string = l.BusinessShortCode + pass_key + l.Timestamp
	l.Password = string(base64.StdEncoding.EncodeToString([]byte(msg)))
}
