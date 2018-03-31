// This is the global webserver

import (
    "log"
    "net/http"
)

func checkError(err error) {
    if err != nil {
        log.Fatal(err)  //log if there is an error
    }
}

func main() {
    err := http.ListenAndServe(":9999", nil)    //server listens to port 9999
    checkError(err)
}
