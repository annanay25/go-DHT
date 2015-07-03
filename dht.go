//setting up the database and querying it.
package go_DHT

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

func addKV(key string, value contact) error {
	err := db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("Files"))
    if err != nil {
        return err
    }
    if err := b.Put([]byte(key), []byte(value)); err != nil {
        return err
    }
    return nil
	})
	return err
}

//buckets have been created. Now, query them.
func queryKV(key string) ([]contact, error) {
	cont, err := db.View(func(tx *bolt.Tx) ([]contact, error) {
        retValue := tx.Bucket([]byte("Files")).Get([]byte(key))
        //fmt.Printf("The value of '%s' is: %v\n", key, retValue)
        return (retValue,nil)
  })
  if err != nil {
  	fmt.Printf("Internal DB error - %v", err)
		return (nil, err)
	}
	else {
		return (cont, nil)
	}
}

