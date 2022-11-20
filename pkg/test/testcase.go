package test

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type (
	CaseInterface interface {
		Name() string
		Compare(t *testing.T, expected interface{}, err error)
	}

	Case struct {
		err  error
		name string
		data interface{}
	}
)

func New(err error, name string, data interface{}) *Case {
	return &Case{err: err, name: name, data: data}
}

func (c *Case) Name() string { return c.name }

func (c *Case) Compare(t *testing.T, expected interface{}, err error) {
	if diff := cmp.Diff(c.data, expected); diff != "" {
		t.Errorf("Name: %s. Case.Data[%+v] != Expected[%+v]. Diff: \n%s", c.name, c.data, expected, diff)
		return
	}

	if c.err != err {
		t.Errorf("Name: %s. Case.Error[%+v] != %+v.", c.name, err, c.err)
		return
	}
}
