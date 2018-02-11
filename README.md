# vgy
Simple Post image to vgy.me


## vgy API doc

API doc is here https://vgy.me/api

## Usage

```golang

import "github.com/frontware.com/vgy"
import "fmt"

func main () {

    vgy.UserKey = "PUT YOUR VGY USER KEY HERE"
    info, err := vgy.UploadImageFile("image.png")
    if err!=nil {
        fmt.Fatal("Upload failed ",err)
    }
    fmt.Println("Upload success")
    fmt.Println(info.String())

}


```