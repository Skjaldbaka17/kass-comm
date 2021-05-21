# kasscom

kasscom is a package for calling the kass-api's most important endpoint "payment".

## Usage

```golang
import "github.com/Skjaldbaka17/kasscom"
var base_Request Request = Request{
	Amount:      2199,
	Description: "Kass bolur",
	Image_Url:   "https://photos.kassapi.is/kass/kass-bolur.jpg",
	Order:       "ABC123",
	Recipient:   "1001000",
	Terminal:    1,
	Expires_In:  90,
	Notify_Url:  "https://example.com/api/callback",
}

kasskom.SetDev()
kasskom.SetAuthToken(test_auth_token)
resp, err := InitiatePayment(&base_Request)
```

## License
[MIT](https://choosealicense.com/licenses/mit/)