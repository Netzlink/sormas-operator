package controller

import (
	"github.com/Netzlink/sormas-operator/pkg/controller/sormas"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sormas.Add)
}
