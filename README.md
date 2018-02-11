# vgy
Simple Post image to vgy.me


## vgy API doc

API doc is here https://vgy.me/api

## Doc

Doc online is on [godoc](https://godoc.org/github.com/Frontware/vgy)

## Usage

```golang

import "github.com/Frontware/vgy"
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