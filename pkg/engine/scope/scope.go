package scope

import (
	"github.com/PicoTools/plan/pkg/engine/object"
)

// global scope
var GlobalScope *Scope

// current scope
var CurrentScope *Scope

type Scope struct {
	// parent's scope
	parent *Scope
	// curent deptch
	depth int
	// scope inside function
	isFunc bool
	// scope inside loop
	isLoop bool
	// break necessity
	isLoopBreak bool
	// continue necessity
	isLoopContinue bool
	// variables storage
	objects map[string]object.Object
}

// NewScope creates new scope
func NewScope(
	parent *Scope,
	depth int,
	isFunc bool,
	isLoop bool,
	objects map[string]object.Object,
) *Scope {
	return &Scope{
		parent:  parent,
		depth:   depth,
		isFunc:  isFunc,
		isLoop:  isLoop,
		objects: objects,
	}
}

// IsGLobal returns true if scope is global
func (s *Scope) IsGlobal() bool {
	return s.parent == nil
}

// Parent returns parent scope
func (s *Scope) Parent() *Scope {
	return s.parent
}

// IsFunc returns true if scope is function
func (s *Scope) IsFunc() bool {
	return s.isFunc
}

// IsInFunc returns true if scope inside of function
func (s *Scope) IsInFunc() bool {
	t := s
	for {
		if t == nil {
			return false
		}
		if t.isFunc {
			return true
		}
		t = t.parent
	}
}

// IsInLoop returns true if scope inside of loop
func (s *Scope) IsInLoop() bool {
	t := s
	for {
		if t == nil {
			return false
		}
		if t.isLoop {
			return true
		}
		t = t.parent
	}
}

// IsBreak returns true if break necessary
func (s *Scope) IsBreak() bool {
	return s.isLoopBreak
}

// IsContinue returns true if continue necessary
func (s *Scope) IsContinue() bool {
	return s.isLoopContinue
}

// SetLoopBreak sets flag for break necessity
func (s *Scope) SetLoopBreak() {
	t := s
	for {
		if t.isLoop {
			// set break on first loop context
			t.isLoopBreak = true
			break
		}
		t = t.parent
	}
}

// UnsetLoopBreak unsets flag for break necessity
func (s *Scope) UnsetLoopBreak() {
	t := s
	for {
		if t.isLoop {
			// set break on first loop context
			t.isLoopBreak = false
			break
		}
		t = t.parent
	}
}

// SetLoopContinue sets flag for continue necessity
func (s *Scope) SetLoopContinue() {
	t := s
	for {
		if t.isLoop {
			// set continue on first loop context
			t.isLoopContinue = true
			break
		}
		t = t.parent
	}
}

// UnsetLoopContinue unsets flag for continue necessity
func (s *Scope) UnsetLoopContinue() {
	t := s
	for {
		if t.isLoop {
			// set continue on first loop context
			t.isLoopContinue = false
			break
		}
		t = t.parent
	}
}

// Depth returns current depth of scopes
func (s *Scope) Depth() int {
	return s.depth
}

// Get returns value of variable from storage (with search inside of parent scope)
func (s *Scope) Get(n string, p bool) object.Object {
	val, ok := s.objects[n]
	if ok {
		return val
	}
	if p && !s.IsGlobal() {
		// search in parent scope
		return s.Parent().Get(n, p)
	} else {
		// nothing founded
		return nil
	}
}

// Put saves variable in scope
func (s *Scope) Put(n string, o object.Object) {
	val := s.Get(n, !s.IsFunc())
	if val == nil {
		s.objects[n] = o
	} else {
		s.reAssign(n, o)
	}
}

// Reassign used for reassignment variable inside scopes
func (s *Scope) reAssign(n string, o object.Object) {
	val := s.objects[n]
	if val == nil {
		s.Parent().reAssign(n, o)
	} else {
		s.objects[n] = o
	}
}
