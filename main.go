package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	 "github.com/rs/cors"
)
type Response struct{
	Filename string `json:"filename"`
	Base64format  string `json:"base64"`

}
func main()  {
	corshandle := cors.Default().Handler(http.DefaultServeMux)
http.HandleFunc("/upload",uplaodhandler)
http.ListenAndServe(":8080",corshandle)
}
func uplaodhandler(w http.ResponseWriter ,  r *http.Request)  {
// validate Method
if r.Method != http.MethodPost {
	http.Error(w,"Method not allow",401)
}

// check size of image
r.Body = http.MaxBytesReader(w,r.Body,5<<20)
	
err := r.ParseMultipartForm(5 << 20)
if err != nil {
	http.Error(w,"File to large",401)
}

// get file
file,handler,err := r.FormFile("photo")
if err != nil {
	http.Error(w,"Method not allow",401)
}
defer file.Close()
// read file image===>>>
filebytes,err := io.ReadAll(file)
if err != nil {
	http.Error(w,"Method not allow",401)
}

// encoding that bytes
encoded := base64.StdEncoding.EncodeToString(filebytes)

// response send to client
resp := Response{
	Filename: handler.Filename,
	Base64format: encoded,

}
w.Header().Set("Content-Type","application/json")
json.NewEncoder(w).Encode(resp)
	
}