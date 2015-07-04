package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
  
  //"github.com/annanay25/go-DHT/httpdaemon"
)


type metadata struct {
	info string
	infohash [IdLength]byte
	contact Contact
}

func NewMeta() *metadata {
	m := new(metadata)
	return m
}

meta = make([]metadata, 1)

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
  //store this information in a slice of metadata objects.
  
  return nil
} 


func main() {
  flag.Parse()
  root := flag.Arg(0)
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}
