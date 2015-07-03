package go_DHT

import (
	"encoding/hex"
	"fmt"
	"bytes"
	"math/rand"
	)
	
const IdLength = 20
type NodeID [IdLength]byte

//make a random node id generator function.
func NewNodeID(data string) (ret NodeID) {
  decoded, _ := hex.DecodeString(data);
  for i := 0; i < IdLength; i++ {
    ret[i] = decoded[i];
  }
  return;
}

func NewRandomNodeID() (ret NodeID) {
  for i := 0; i < IdLength; i++ {
    ret[i] = uint8(rand.Intn(256));
  }
  return;
}

//distance metric is the XOR of the nodeIDs:
func (first NodeID) NodeDist(second NodeID) (dist NodeID) {
	for i:= 0; i < IdLength; i++ {
    dist[i] = first[i] ^ second [i];
  }
  return dist;
}


