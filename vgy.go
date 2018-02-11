package vgy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Response returned information from vgy.me after posting an image

type Response struct {
	Error     bool
	URL       string // direct URL of the image
	Image     string
	Extension string
	Filesize  int64  // File size in bytes
	DeleteURL string `json:"delete"` // URL to delete the image from vgy.me hosting
}

// Print returns string formated Response
func (r Response) Print() string {
	result := "URL: " + r.URL + "\n"
	result += "Delete URL: " + r.DeleteURL + "\n"
	result += "Image: " + r.Image + "." + r.Extension + "\n"
	result += fmt.Sprintf("Filesize: %d", r.Filesize)
	return result
}

// UserKey must be created at https://vgy.me/account/details
// It's not mandatory. But if you don't set it you won't get the uploaded image attached to your vgy.me account
var UserKey string

const urlUpload = "https://vgy.me/upload"

// UploadImageFile uploads an image (.png, .jpg) to vgy.me
func UploadImageFile(fileName string) (response Response, err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	// Add User Key if provided
	if UserKey != "" {
		w.CreateFormField("userkey")
		fw.Write([]byte(UserKey))
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", urlUpload, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	defer res.Body.Close()
	if respBody, err := ioutil.ReadAll(res.Body); err == nil {
		if err = json.Unmarshal(respBody, &response); response.Error {
			err = errors.New("error")
		}
	}
	return
}
