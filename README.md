![Go](https://github.com/Paradox-3arthling/daraja_wrapper/workflows/Go/badge.svg?branch=master)
# daraja_wrapper
## What it does
This is a wrapper library to the daraja API endpoints from safaricom Mpesa/Lipa na Mpesa

> This is a project is still in dev, though working parts
> should be ok.

## Why is it useful
- Making it easier to access the daraja API endpoints without actually having to know much about them.
- Hopefully the project can be used to learn more about the endpoints and the 'net/http' package from golang.
- Hopes to create a package that is able scale and while still being very fast.

## How to get started
> !!IMPORTANT!! do check for errors, am ignoring them for the sake of short examples.
This will be sandbox examples:
### Lipa na M-pesa
Start by importing the project `import "github.com/paradox-3arthling/daraja_wrapper"`:
```
	lipa_mpesa := daraja_wrapper.LipaNaMpesaPayStruct{
		BusinessShortCode: "",
		Timestamp:         "yyyymmddhhmmss",
		Amount:            "420",
		PartyA:            "254XXXXXXXXX",
		PartyB:            "",
		PhoneNumber:       "254XXXXXXXXX",
		CallBackURL:       "https://ip_address/response_path",
		AccountReference:  "test",
		TransactionDesc:   "test",
	}
	// lipa na mpesa acc
	mpesa_auth := daraja_wrapper.Auth{
		Key:    "some_key",
		Secret: "some_secret",
	}
	resp, _ := lipa_mpesa.LipaNaMpesaPayment(&mpesa_auth, "some_pass")
    // check for error 1st!
	fmt.Println("resp ->", resp)
```
### Setting up callback url
- A lot of the API endpoints from daraja have a field called `CallBackURL` this is, where if the process you are trying to do gets successful processed, safaricom will send extra information on the transaction you started was completed on the user side.
- Side note is that you don't have to use my callback to make the transaction there are a lot of alternatives.
> I'm using a side project(import `import "github.com/paradox-3arthling/async_response_server"`) for this and requires a little set up for the tunnelling(done by node.js) but you can skip the tunnelling process if you are not behind a proxy.
```
var succ_mess string = `{
	"ResponseCode": "00000000",
	"ResponseDesc": "success"
	}`
```
This is the response we'll send back to safaricom for successfully accepting the information sent.(Could do better than this probably :D)
```
	port := ":3310"
	ts := async_response_server.CreateHookServerAsync(port)
	fmt.Printf("The callback URL is '%s'\n", ts.Url)
	defer ts.Close()
	fmt.Println("waiting on info. from our saf.")
	feedback := <-ts.Feed_back
	ts.Feed_back <- succ_mess
	fmt.Println("saf message:\n", feedback)
```
This creates a server that listens on `ts.Url` variable which listens for information and responds with the `succ_mess` message.

The tunneling, you can skip this if u have direct access to the internet. You also need `ngrok` installed.
Since the server we created listens on port `3310` we will run `ngrok http 3310`. Output on which url to use will come up, best to use the
https one. We'll call this `ngrok_ip`

Finally our `ts.Url` will show ur local IP which would be wrong we need to replace the IP address of this URL with either `ngrok_ip` or your public IP for those connected directly to the internet. This will be your `CallBackURL` for transactions that need this.

## WIP
- Check issues for next feature updates to be done.

## Who maintains and can contributes.
At the moment am the only maintainer and contributer open to help, contact me at floydqaranja@gmail.com 
