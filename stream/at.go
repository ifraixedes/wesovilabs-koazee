package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

const OperationAtIdentifier = ":at"

type at struct {
	items interface{}
	index int
}

func (op *at) name() string {
	return OperationAtIdentifier
}

func (op *at) run() output {
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	return output{itemsValue.Index(op.index).Interface(), nil}
}

func (op *at) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "You can not take an element for a nil stream")
	}
	itemsValue := reflect.ValueOf(op.items)
	len := itemsValue.Len()
	if len == 0 {
		return errors.ItemsNil(op.name(), "You can not take an element for an empty stream")
	}
	if op.index < 0 || len-1 < op.index {
		return errors.InvalidStreamIndex(op.name(), "Invalid index %d; A valid index must be between 0 and %d", op.index, len-1)
	}
	return nil
}

// At returns the element in the stream in the given position
func (s *stream) At(index int) output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return at{current.items, index}.run()
}