package actor

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
)

func StartActor(receiver IActorReceiver) *AID {
	ac := defaultActorSystem.startActor(receiver)
	return ac.getAID()
}

func StartWebActor(receiver IActorReceiver, nodeName string, address string) *AID {
	ac := defaultActorSystem.startWebActor(receiver, nodeName, address)
	return ac.getAID()
}

func Call(aid *AID, function interface{}, args ...interface{}) ([]interface{}, error) {
	if defaultActorSystem == nil {
		panic("ActorSystem not initialized")
	}

	ac, err := defaultActorSystem.getActorByAIDWithLock(aid)
	if err != nil {
		return nil, fmt.Errorf("ActorSystem.Call err: %v, AID %v", err, aid)
	}
	return ac.call(function, args...)
}

func StartWebServer(mainCtx context.Context, port string) {
	ctx, cancel := context.WithCancel(mainCtx)

	defaultActorSystem.server = &http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/Call", httpCall)

	defaultActorSystem.server.Shutdown(ctx)
	defaultActorSystem.serverCloseFunc = cancel

	go func() {
		fmt.Printf("Actor Server Start %v \n", defaultActorSystem.server.Addr)
		fmt.Println(defaultActorSystem.server.ListenAndServe())
		fmt.Printf("Actor Server End %v \n", defaultActorSystem.server.Addr)
	}()
}

func httpCall(w http.ResponseWriter, req *http.Request) {

	dec := gob.NewDecoder(req.Body)

	request := new(Request)
	if err := dec.Decode(request); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	response := new(Response)
	if err := defaultActorSystem.handleRemoteCall(request, response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(response)
	w.Write(buffer.Bytes())
}
