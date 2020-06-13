package actor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"reflect"
)

var _ IActorReceiver = &Actor{}
var _ iActor = &Actor{}

type WebActor struct {
	aid        *AID
	receiver   reflect.Value
	shutdownCh chan bool

	address string
	client  *http.Client
}

func (a *WebActor) start(receiver IActorReceiver, aid *AID) bool {
	a.aid = aid
	a.receiver = reflect.ValueOf(receiver)
	return true
}

func (a *WebActor) GetNodeName() string {
	return a.aid.NodeName
}

func (a *WebActor) getAID() *AID {
	return a.aid
}

func (a *WebActor) getReceiver() reflect.Value {
	return a.receiver
}

func (a *WebActor) setAID(aid *AID) {
	a.aid = aid
}

func (a *WebActor) call(function interface{}, args ...interface{}) ([]interface{}, error) {
	req := defaultActorSystem.createRequest(a, function, args...)

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(req)

	httpReq, err := http.NewRequest("POST", a.address+"/Call", &buffer)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := gob.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		return nil, fmt.Errorf("%v", resp.Status)
	}

	response := new(Response)
	if err := dec.Decode(response); err != nil {
		return response.Results, err
	}

	return response.Results, nil
}
