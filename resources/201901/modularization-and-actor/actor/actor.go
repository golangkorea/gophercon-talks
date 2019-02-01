package actor

import (
	"fmt"
	"reflect"
	"container/list"
)

type AID struct {
	ActorID  uint64
	NodeName string
}

func (a *AID) String() string {
	return fmt.Sprintf("<%d, %s>", a.ActorID, a.NodeName)
}


type IActorReceiver interface {

	GetNodeName() string
}

type iActor interface {
	IActorReceiver
	getReceiver() reflect.Value
	getAID() *AID

	setAID(aid *AID)
	call(function interface{}, args ...interface{}) ([]interface{}, error)
}

var _ IActorReceiver = &Actor{}
var _ iActor = &Actor{}

type Actor struct {
	aid       		*AID
	receiver 	    reflect.Value
	queue		    *list.List
	inChan    		chan *ActorCall
}

func (a *Actor) start(receiver IActorReceiver, aid *AID) bool {
	a.aid = aid
	a.receiver = reflect.ValueOf(receiver)
	a.queue = list.New()
	a.inChan = make(chan *ActorCall, 1)
	go a.loop()
	return true
}

func (a *Actor) loop() {
	for {
		if a.queue.Len() == 0 {
			select {
			case actorCall := <-a.inChan:
				a.queue.PushBack(actorCall)
			}
		} else {
			select {
			case actorCall := <-a.inChan:
				a.queue.PushBack(actorCall)
			default:
				actorCall := a.queue.Front().Value.(*ActorCall)
				a.queue.Remove(a.queue.Front())
				a.process(actorCall)
			}
		}

	}
}


func (a *Actor) GetNodeName() string {
	return a.aid.NodeName
}

func (a *Actor) getAID() *AID {
	return a.aid
}

func (a *Actor) getReceiver() reflect.Value {
	return a.receiver
}

func (a *Actor) setAID(aid *AID) {
	a.aid = aid
}

func (a *Actor) process(actorCall *ActorCall) {
	actorCall.Results = actorCall.Function.Call(actorCall.Args)
	if actorCall.Done != nil {
		actorCall.Done <- actorCall
	}
}


func (a *Actor) makeActorCall(done chan *ActorCall , function interface{}, args ...interface{}) *ActorCall {

	v := reflect.ValueOf(function)
	if v.Kind() != reflect.Func {
		fn, _ := defaultActorSystem.findFunc(a, v.String())
		v = reflect.ValueOf(fn)

	}

	valuedArgs := make([]reflect.Value, len(args)+1)
	valuedArgs[0] = a.receiver
	for i, x := range args {
		valuedArgs[i+1] = reflect.ValueOf(x)
	}

	return &ActorCall{Function:v, Args:valuedArgs, Done:done}

}

func (a *Actor) call(function interface{}, args ...interface{}) ([]interface{}, error) {
	done := make(chan *ActorCall, 0)
	a.inChan <- a.makeActorCall(done, function, args...)
	actorCall, _ := <-done
	return actorCall.GetResults()
}