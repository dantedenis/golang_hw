package main
import "fmt"

import "net/http"

type LicenseResponse struct {
    Id      int         `json:"id"`
    OrgName string      `json:"orgname"`
    GUID    string      `json:"guid"`
}

func main() {
    http.HandleFunc("/getinfo", LicenseServer)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Println(err)
    }
}

func LicenseServer(w http.ResponseWriter, r *http.Request) {
    //license := LicenseResponse{1, "TestName", "123"}
    //response, err := json.Marshall(license)
    response := "Hello world"

    fmt.Fprintf(w, string(response))
