//setting up the database and querying it.

import (
	"net"
	"fmt"
	"log"
	"os"
	"github.com/boltdb/bolt"
	"github.com/annanay25/go-DHT/routing_table"
)

func setupDB() {
	db, err := bolt.Open("myHashes.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer os.Remove(db.Path())
    defer db.Close()
}

e := (func addKV() error {
	err := db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucket([]byte(reqObj.params.bucketName))
    if err != nil {
        return err
    }
    if err := b.Put([]byte(reqObj.params.key), []byte(reqObj.params.value)); err != nil {
        return err
    }
    return nil
	})
	return err
})

if e!= nil {
	//write error handling later
}

//buckets have been created. Now, query them.
func queryKV() []contact {
	db.View(func(tx *bolt.Tx) error {
        value := tx.Bucket([]byte(reqObj.params.bucketName)).Get([]byte(reqObj.params.key))
        fmt.Printf("The value of '%s' is: %s\n", req.Obj.params.key, value)
        //pass this k/v pair to the http response object now.
        return nil
  })
}

