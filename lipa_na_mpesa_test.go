package daraja_wrapper

import "testing"

const correct_pass = ""

func TestPasswordGeneration(t *testing.T) {
	lipa_mpesa := LipaNaMpesaPayStruct{
		BusinessShortCode: "123",
		Timestamp:         "20200420042000",
	}
	pass := "somerandpass"
	lipa_mpesa.GenPassword(pass)
	if lipa_mpesa.Password != correct_pass {
		t.Errorf("`lipa_mpesa.GenPassword` error Got: %v, expected %v", lipa_mpesa.Password, correct_pass)
	}
}
