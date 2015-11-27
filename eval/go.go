// Copyright 2015 The Neugram Authors. All rights reserved.
// See the LICENSE file for rights to use this source code.

package eval

import (
	"fmt"
	gotypes "go/types"
	"reflect"

	"neugram.io/eval/gowrap"
	"neugram.io/lang/tipe"
)

type GoPkg struct {
	Type  *tipe.Package
	GoPkg *gotypes.Package
	Wrap  *gowrap.Pkg
}

type GoFunc struct {
	Type *tipe.Func
	Func interface{}
}

func (f GoFunc) call(args ...interface{}) (res []interface{}, err error) {
	var vres []reflect.Value
	v := reflect.ValueOf(f.Func)
	if f.Type.Variadic {
		nonVarLen := len(f.Type.Params.Elems) - 1
		var vargs []reflect.Value
		for i := 0; i < nonVarLen; i++ {
			vargs = append(vargs, reflect.ValueOf(args[i]))
		}
		if len(args) > nonVarLen {
			vargs = append(vargs, reflect.ValueOf(args[nonVarLen:]))
		} else {
			vargs = append(vargs, reflect.ValueOf([]interface{}{}))
		}
		vres = v.CallSlice(vargs)
	} else {
		return nil, fmt.Errorf("Call GoFunc TODO")
	}

	res = make([]interface{}, len(vres))
	for i, v := range vres {
		res[i] = v.Interface()
	}
	return res, nil
}
