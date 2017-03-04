// this file was generated by gomacro command: import "archive/tar"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "archive/tar"
	. "reflect"
)

func init() {
	Binds["archive/tar"] = map[string]Value{
		"ErrFieldTooLong":	ValueOf(&pkg.ErrFieldTooLong).Elem(),
		"ErrHeader":	ValueOf(&pkg.ErrHeader).Elem(),
		"ErrWriteAfterClose":	ValueOf(&pkg.ErrWriteAfterClose).Elem(),
		"ErrWriteTooLong":	ValueOf(&pkg.ErrWriteTooLong).Elem(),
		"FileInfoHeader":	ValueOf(pkg.FileInfoHeader),
		"NewReader":	ValueOf(pkg.NewReader),
		"NewWriter":	ValueOf(pkg.NewWriter),
		"TypeBlock":	ValueOf(pkg.TypeBlock),
		"TypeChar":	ValueOf(pkg.TypeChar),
		"TypeCont":	ValueOf(pkg.TypeCont),
		"TypeDir":	ValueOf(pkg.TypeDir),
		"TypeFifo":	ValueOf(pkg.TypeFifo),
		"TypeGNULongLink":	ValueOf(pkg.TypeGNULongLink),
		"TypeGNULongName":	ValueOf(pkg.TypeGNULongName),
		"TypeGNUSparse":	ValueOf(pkg.TypeGNUSparse),
		"TypeLink":	ValueOf(pkg.TypeLink),
		"TypeReg":	ValueOf(pkg.TypeReg),
		"TypeRegA":	ValueOf(pkg.TypeRegA),
		"TypeSymlink":	ValueOf(pkg.TypeSymlink),
		"TypeXGlobalHeader":	ValueOf(pkg.TypeXGlobalHeader),
		"TypeXHeader":	ValueOf(pkg.TypeXHeader),
	}
	Types["archive/tar"] = map[string]Type{
		"Header":	TypeOf((*pkg.Header)(nil)).Elem(),
		"Reader":	TypeOf((*pkg.Reader)(nil)).Elem(),
		"Writer":	TypeOf((*pkg.Writer)(nil)).Elem(),
	}
}