package controller

import (
	"github.com/kiali/kiali-operator/pkg/controller/kiali"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kiali.Add)
}
