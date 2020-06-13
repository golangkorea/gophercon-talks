package actor

import "reflect"

type ActorCall struct {
	Function reflect.Value
	Args     []reflect.Value
	Results  []reflect.Value
	Done     chan *ActorCall
	Error    error
}

func (c ActorCall) GetResults() ([]interface{}, error) {
	if c.Error != nil {
		return nil, c.Error
	}
	values := c.Results

	var rtnResults []interface{}
	var err error
	for _, x := range values {
		itf := x.Interface()
		errorType := reflect.TypeOf((*error)(nil)).Elem()

		switch x.Type() {
		case errorType:
			if castingErr, ok := itf.(error); ok && castingErr != nil {
				err = x.Interface().(error)
			}
		default:
			rtnResults = append(rtnResults, itf)
		}
	}
	return rtnResults, err
}
