package failover

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
