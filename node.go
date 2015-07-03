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

//function that returns the number of leading zeros in the XOR of two NodeIDs. Which in our implementation is also the bucket number of that respective node.
func (node NodeID) PrefixLen() (ret int) {
  for i := 0; i < IdLength; i++ {
    for j := 0; j < 8; j++ {
      if (node[i] >> uint8(7 - j)) & 0x1 != 0 {
        return i * 8 + j;
      }
    }
  }
  return IdLength * 8 - 1;
}


