package visitor

import "github.com/PicoTools/plan/pkg/engine/object"

// hold return value in process of evaluation
var retValue object.Object

// ClearRet clears return value
func ClearRet() {
	retValue = nil
}
