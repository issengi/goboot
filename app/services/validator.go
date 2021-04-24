package services

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

type GinDefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (g *GinDefaultValidator) ValidateStruct(i interface{}) error {
	if kindOfData(i) == reflect.Struct {
		g.lazyinit()
		if err := g.validate.Struct(i); err != nil {
			return err
		}
	}
	return nil
}

func (g *GinDefaultValidator) Engine() interface{} {
	g.lazyinit()
	return g.validate
}

func (g *GinDefaultValidator) lazyinit() {
	g.once.Do(func() {
		g.validate = validator.New()
		g.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
