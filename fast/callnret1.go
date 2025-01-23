// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * callnret1.go
 *
 *  Created on Apr 20, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	xr "github.com/WilliamNHarvey/gomacro/xreflect"
)

func (c *Comp) callnret1(call *Call, maxdepth int) I {
	expr := call.Fun
	exprfun := expr.AsX1()
	if expr.Sym != nil && expr.Sym.Desc.Index() == NoIndex {
		c.Errorf("internal error: callnret1() invoked for constant function %#v. use call_builtin() instead", expr)
	}

	kret := expr.Type.Out(0).Kind()
	argfuns := call.MakeArgfunsX1()
	var ret I
	switch kret {
	case xr.Bool:
		ret = func(env *Env) bool {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0.Bool()
		}
	case xr.Int:
		ret = func(env *Env) int {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return int(ret0.Int())
		}
	case xr.Int8:
		ret = func(env *Env) int8 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return int8(ret0.Int())
		}
	case xr.Int16:
		ret = func(env *Env) int16 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return int16(ret0.Int())
		}
	case xr.Int32:
		ret = func(env *Env) int32 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return int32(ret0.Int())
		}
	case xr.Int64:
		ret = func(env *Env) int64 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0.Int()
		}

	case xr.Uint:
		ret = func(env *Env) uint {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return uint(ret0.Uint())
		}

	case xr.Uint8:
		ret = func(env *Env) uint8 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return uint8(ret0.Uint())
		}

	case xr.Uint16:
		ret = func(env *Env) uint16 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return uint16(ret0.Uint())
		}

	case xr.Uint32:
		ret = func(env *Env) uint32 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return uint32(ret0.Uint())
		}

	case xr.Uint64:
		ret = func(env *Env) uint64 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0.Uint()
		}

	case xr.Uintptr:
		ret = func(env *Env) uintptr {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return uintptr(ret0.Uint())
		}

	case xr.Float32:
		ret = func(env *Env) float32 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return float32(ret0.Float())
		}

	case xr.Float64:
		ret = func(env *Env) float64 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0.Float()
		}

	case xr.Complex64:
		ret = func(env *Env) complex64 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return complex64(ret0.Complex())
		}

	case xr.Complex128:
		ret = func(env *Env) complex128 {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0.Complex()
		}

	case xr.String:
		ret = func(env *Env) string {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0.String()
		}

	default:
		ret = func(env *Env) xr.Value {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := callxr(funv, argv)[0]
			return ret0

		}

	}
	return ret
}
