![vgy](https://vgy.me/static/img/public/logo/navbar.png)
# vgy

Simple Post image to vgy.me

## Important

We are not affilated to vgy.me

## vgy API doc

API doc is here https://vgy.me/api

[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

![](https://img.shields.io/badge/style-plastic-green.svg?style=plastic)
![](https://img.shields.io/badge/Developed%20by-Frontware-yellowgreen.svg)

## Doc

Doc online is on [godoc](https://godoc.org/github.com/Frontware/vgy)

## Usage

```golang

import "github.com/Frontware/vgy"
import "fmt"

func main () {

    // Get your user key here: https://vgy.me/account/details#userkeys
    vgy.UserKey = "PUT YOUR VGY USER KEY HERE"
    // You must have the file "image.png" in current folder
    info, err := vgy.UploadImageFile("image.png")
    if err!=nil {
        fmt.Fatal("Upload failed ",err)
    }
    fmt.Println("Upload to "+info.URL)
}

```