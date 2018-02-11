# vgy

![vgy](https://vgy.me/static/img/public/logo/navbar.png)
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
    // You must have the file "image.png" in current folder
    info, err := vgy.UploadImageFile("image.png")
    if err!=nil {
        fmt.Fatal("Upload failed ",err)
    }
    fmt.Println("Upload to "+info.URL)
}

```