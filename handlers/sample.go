package handlers
import "net/http"
//Home renders a Hello message
func Home(rw http.ResponseWriter, req *http.Request){
  rw.Write([]byte("Hello From Webo!"))
}