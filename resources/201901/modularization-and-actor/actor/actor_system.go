package actor

import (
	"context"
	"sync"
	"sync/atomic"
	"fmt"
	"runtime"
	"reflect"
	"strings"
	"net/http"
)

type ActorProtocolType uint16

const (
	None ActorProtocolType = iota
	Web
)

type Request struct {
	ActorID  uint64
	NodeName string
	FuncName string
	Args     []interface{}
}

type Response struct {
	Err    			error
	Results 		[]interface{}
}

var defaultActorSystem *ActorSystem

func init() {
	defaultActorSystem = newActorSystem()
}

type ActorSystem struct {
	mutex              	sync.RWMutex
	nowActorID         	uint64
	actorMapByAID       map[*AID]iActor
	actorMapByNodeName 	map[string]iActor

	server              *http.Server
	serverCloseFunc     context.CancelFunc
}

func newActorSystem() *ActorSystem {
	d := &ActorSystem{
		actorMapByAID:        make(map[*AID]iActor),
		actorMapByNodeName:   make(map[string]iActor),
		nowActorID:        	  0,
	}
	return d
}

func (as *ActorSystem) startActor(receiver IActorReceiver) iActor {
	actor := &Actor{}

	as.mutex.Lock()
	defer as.mutex.Unlock()

	aid := as.createAID(receiver.GetNodeName())
	actor.start(receiver, aid)
	as.actorMapByAID[aid] = actor
	as.actorMapByNodeName[receiver.GetNodeName()] = actor

	return actor
}

func (as *ActorSystem) startWebActor(receiver IActorReceiver, nodeName string, address string) iActor {
	actor := &WebActor{}

	as.mutex.Lock()
	defer as.mutex.Unlock()

	aid := as.createAID(nodeName)
	actor.start(receiver, aid)
	as.actorMapByAID[aid] = actor
	as.actorMapByNodeName[nodeName] = actor

	actor.address = address
	actor.client = &http.Client{}
	return actor
}

func (as *ActorSystem) createAID(nodeName string) *AID {
	id := atomic.AddUint64(&as.nowActorID, 1)
	return &AID{NodeName:nodeName, ActorID:id}
}

func (as *ActorSystem) getActorByAIDWithLock(aid *AID) (iActor, error) {
	as.mutex.RLock()
	defer as.mutex.RUnlock()

	actor, ok := as.actorMapByAID[aid]
	if !ok {
		return nil, fmt.Errorf("%s", "Not Found Actor")
	}
	return actor, nil
}

func (d *ActorSystem) getActorByNodeNameWithLock(nodeName string) (iActor, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	actor, ok := d.actorMapByNodeName[nodeName]
	if !ok {
		return nil, fmt.Errorf("%s", "Not Found Actor")
	}
	return actor, nil
}

func (as *ActorSystem) createRequest(ac iActor, function interface{}, args ...interface{}) *Request {


	req := &Request{NodeName:ac.GetNodeName()}

	v := reflect.ValueOf(function)
	if v.Kind() == reflect.Func {
		funcFullName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
		fmt.Println(funcFullName)
		tokens := strings.Split(funcFullName, ".")
		req.FuncName = tokens[len(tokens)-1]
	} else {
		req.FuncName = v.String()
	}

	req.Args = args
	return req
}

func (as *ActorSystem) handleRemoteCall(req *Request, res *Response) error {
	ac, acErr := as.getActorByNodeNameWithLock(req.NodeName)
	if acErr != nil {
		res.Err = acErr
		return acErr
	}
	fn, fnErr := as.findFunc(ac, req.FuncName)
	if fnErr != nil {
		res.Err = fnErr
		return fnErr
	}
	rtn, callErr := ac.call(fn, req.Args...)
	res.Results = rtn
	res.Err = callErr
	return callErr
}

func (as *ActorSystem) findFunc(ac iActor, funcName string) (interface{}, error) {
	t := ac.getReceiver().Type()
	method, ok := t.MethodByName(funcName)
	if !ok {
		return nil, fmt.Errorf("%s", "Not Found Func")
	}
	fn := method.Func.Interface()
	return fn, nil
}