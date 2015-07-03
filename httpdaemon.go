//keeping in mind the architechture for "github.com/Gouthamve/goplay" I am implementing a single server daemon that will handle all functionality.

//NODE daemon listening for any kind of request/response.
//action will be taken accordingly. 
//Tracker daemon will be implemented separately in this architechture.

//logging struct for a transaction
//type transac struct {
//this can wait.
package go_DHT

import (
	"net"
	"net/http"
	"log"
	"encoding/json"
	"bytes"
	"io/ioutil"
	
	"github.com/annanay25/go-DHT/routing_table"
	"github.com/annanay25/go-DHT/dht.go"
	"github.com/gorilla/mux"
)

//call to a node to send its metadata
func getMetadata(c contact) err error {
	var buffer bytes.Buffer
	addr := net.JoinHostPort(c.ipaddress, c.port)
	buffer.WriteString(addr)
	buffer.WriteString("/getMeta")
	resp, err := http.Get(buffer)
	if err != nil {
		log.Fatal(err)
	}
	//Read the body of the response.
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	meta := NewMeta()
	err = json.Unmarshal([]byte(jsonDataFromHttp), &meta)
	
	//TODO now what?
}

func dbAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	//first query to check if the same k/v pair exists. If it does, discard this operation.
	cont, err := queryKV (key)
	

func metaHandler(w http.ResponseWriter, r *http.Request) {
	//serve all the metadata - all the data that is stored that the user wants to share.
	

func main() {
	router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", pingHandler)
  router.HandleFunc("/getMeta", metaHandler)
  router.HandleFunc("/add/{key}", dbAddHandler)
  router.HandleFunc("/findvalue/{key}", queryHandler)
	log.Fatal(http.ListenAndServe(":6789", nil)
}

