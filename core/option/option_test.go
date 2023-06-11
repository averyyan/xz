package xzoption

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

type car struct {
	Name string
}

func TestOption(t *testing.T) {
	_car := &car{}
	var opts []Option[car]
	opts = append(opts, WithName[car]("test"))
	for _, opt := range opts {
		opt(_car)
	}
	then.AssertThat(t, _car.Name, is.EqualTo("test"))
}
