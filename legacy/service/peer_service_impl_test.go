package service

import (
	"testing"
	"fmt"
)

func (mh *MockHandler) handle(ms interface{}) (interface{},error){

	fmt.Print(ms)

	return "success", nil
}


func TestPeerServiceImpl_GetPeerByPeerID(t *testing.T) {
	//var peerServiceImpl := PeerServiceImpl{}
}
