# kasscom

kasscom is a package for calling the kass-api's most important endpoint "payment". 

See the api for more details: https://kass.github.io/api

## Usage

```golang
import (
    "log"
"github.com/Skjaldbaka17/kasscomm"
)

var base_Request kasscomm.Request = kasscomm.Request{
    Amount:      2199,
    Description: "Kass bolur",
    Image_Url:   "https://photos.kassapi.is/kass/kass-bolur.jpg",
    Order:       "ABC123",
    Recipient:   "1001000",
    Terminal:    1,
    Expires_In:  90,
    Notify_Url:  "https://example.com/api/callback",
}

var my_auth_token string = "MY_AUTH_TOKEN"

//kasscomm.SetDev() //for test env
kasscomm.SetProd() //for real env
kasscomm.SetAuthToken(my_auth_token)
resp, err := kasscomm.InitiatePayment(&base_Request)

if err != nil {
    log.Printf("Errors happen... %s\n", err)
    return err
}

log.Println(resp)
```

## License
[MIT](https://choosealicense.com/licenses/mit/)