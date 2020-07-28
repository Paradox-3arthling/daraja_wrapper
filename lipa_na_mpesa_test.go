package daraja_wrapper

import "testing"

const correct_pass = "MTIzNDIwc29tZXJhbmRwYXNzMjAyMDA0MjAwNDIwMDA="

func TestPasswordGeneration(t *testing.T) {
	lipa_mpesa := LipaNaMpesaPayStruct{
		BusinessShortCode: "123420",
		Timestamp:         "20200420042000",
	}
	pass := "somerandpass"
	lipa_mpesa.GenPasswordAndAssign(pass)
	if lipa_mpesa.Password != correct_pass {
		t.Errorf("\n`lipa_mpesa.GenPassword` error Got: '%v'\nexpected '%v'", lipa_mpesa.Password, correct_pass)
	}
}
