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
 * genimport.go
 *
 *  Created on May 26, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"bytes"
	"fmt"
	"go/constant"
	"go/types"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/base/untyped"
)

type Output = output.Output

type genimport struct {
	output      *Output
	mode        ImportMode
	gpkg        *types.Package
	scope       *types.Scope
	names       []string
	pkgrenames  map[string]string // map[path]name of packages to import, where name:s are guaranteed to be unique
	out         *bytes.Buffer
	path        string
	name, name_ string
	proxyprefix string
	reflect     string
}

var inceptionFileDeclarations = []byte(`

/**
 * Declare and fill a global variable Packages, whose type is compatible
 * with the global variable github.com/cosmos72/gomacro/imports.Packages
 *
 * If you want to automatically register this package's declarations into
 *   github.com/cosmos72/gomacro/imports.Packages
 * to let gomacro know about this package, and allow importing it without compiling
 * a plugin, you can add the following to some _other_ file in this directory:
 *
 * import "github.com/cosmos72/gomacro/imports"
 *
 * func init() {
 *     for k, v := range Packages {
 *         imports.Packages[k] = v
 *     }
 * }
 *
 * Such code is _not_ automatically added to this file, because it would introduce
 * a dependency on gomacro packages, which may be undesiderable.
 */

type Package = struct {
	Name     string
	Binds    map[string]r.Value
	Types    map[string]r.Type
	Proxies  map[string]r.Type
	Untypeds map[string]string
	Wrappers map[string][]string
}

var Packages = make(map[string]Package)
`)

var pluginMainFileContent = []byte(`
package main

import . "reflect"

type Package = struct {
	Name     string
	Binds    map[string]Value
	Types    map[string]Type
	Proxies  map[string]Type
	Untypeds map[string]string
	Wrappers map[string][]string
}

var Packages = make(map[string]Package)

func main() {
}
`)

func createPluginMainFile(dir string) (string, error) {
	filepath := paths.Subdir(dir, "main.go")
	return filepath, ioutil.WriteFile(filepath, pluginMainFileContent, os.FileMode(0644))
}

func createImportFile(o *Output, out *bytes.Buffer, path string, gpkg *types.Package, mode ImportMode) (isEmpty bool) {

	gen := newGenImport(o, out, path, gpkg, mode)
	if gen == nil {
		return true
	}
	gen.write()
	return false
}

func newGenImport(o *Output, out *bytes.Buffer, path string, gpkg *types.Package, mode ImportMode) *genimport {
	scope := gpkg.Scope()
	names := scope.Names()

	isEmpty := true
	for _, name := range names {
		if obj := scope.Lookup(name); obj.Exported() {
			switch obj.(type) {
			case *types.Const, *types.Var, *types.Func, *types.TypeName:
				isEmpty = false
				break
			}
		}
	}
	if isEmpty {
		return nil
	}

	gen := &genimport{output: o, mode: mode, gpkg: gpkg, scope: scope, names: names, out: out, path: path}

	if mode == ImInception {
		gen.reflect = "r."
		gen.name = gpkg.Name()
	}
	if mode == ImPlugin {
		gen.proxyprefix = "P_"
	} else {
		gen.proxyprefix = fmt.Sprintf("P_%s_", sanitizeIdent(path))
	}
	return gen
}

func (gen *genimport) write() {

	gen.writePreamble()

	gen.writeBinds()
	gen.writeTypes()
	gen.writeProxies()
	gen.writeUntypeds()
	gen.writeWrappers()

	gen.out.WriteString("\n\t}\n}\n")
	gen.writeInterfaceProxies()
}

type mapdecl struct {
	out  *bytes.Buffer
	head string
	foot string
}

func (gen *genimport) mapdecl(head string) mapdecl {
	if strings.IndexByte(head, '%') >= 0 {
		head = fmt.Sprintf(head, gen.reflect)
	}
	return mapdecl{gen.out, head, ""}
}

func (d *mapdecl) header() {
	if len(d.head) != 0 {
		d.out.WriteString(d.head)
		d.out.WriteByte('{')
		d.head = ""
		d.foot = "\n\t\t}"
	}
}

func (d *mapdecl) footer() {
	if len(d.foot) != 0 {
		d.out.WriteString(d.foot)
		d.out.WriteString(", ")
	}
}

func (d *mapdecl) footer1(comma bool) {
	if len(d.foot) != 0 {
		d.out.WriteString(d.foot)
		if comma {
			d.out.WriteString(", ")
		}
	}
}

func (gen *genimport) collectPackageImportsWithRename(requireAllInterfaceMethodsExported bool) {
	gen.pkgrenames = collectPackageImportsWithRename(gen.output, gen.gpkg, requireAllInterfaceMethodsExported)
	gen.name = gen.pkgrenames[gen.path]
	if gen.name == "" {
		gen.name = packageSanitizedName(gen.path)
	}
	if gen.mode != ImInception {
		gen.name_ = gen.name + "."
	}
}

func (gen *genimport) writePreamble() {
	mode := gen.mode
	out := gen.out

	var alias, filepkg string
	switch mode {
	case ImBuiltin:
		alias = "_b "
		filepkg = "imports"
	case ImThirdParty:
		filepkg = "thirdparty"
	case ImPlugin:
		filepkg = "main"
	case ImInception:
		alias = "_i "
		filepkg = gen.name
	}

	fmt.Fprintf(gen.out, `// this file was generated by gomacro command: import %s%q
// DO NOT EDIT! Any change will be lost when the file is re-generated

package %s

import (`, alias, gen.path, filepkg)

	if mode == ImInception {
		fmt.Fprint(out, "\n\tr \"reflect\"")
	} else {
		fmt.Fprint(out, "\n\t. \"reflect\"")
	}
	gen.collectPackageImportsWithRename(true)
	for path, name := range gen.pkgrenames {
		if mode == ImInception && path == gen.gpkg.Path() {
			continue // writing inside the package: it should not import itself
		} else {
			// always name the imported package: its name may differ from paths.FileName(path)
			fmt.Fprintf(out, "\n\t%s %q", name, path)
		}
	}
	fmt.Fprintf(out, "\n)\n")

	pathToRegister := gen.path
	if mode == ImInception {
		pathToRegister = gen.gpkg.Path()
		gen.pkgrenames[gen.path] = "" // writing inside the package: remove the package prefix
		out.Write(inceptionFileDeclarations)
	}

	fmt.Fprintf(out, `
// reflection: allow interpreted code to import %q
func init() {
	Packages[%q] = Package{
		Name: %q,
		`, pathToRegister, pathToRegister, gen.gpkg.Name())
}

func (gen *genimport) writeBinds() {
	d := gen.mapdecl("Binds: map[string]%sValue")

	for _, name := range gen.names {
		if obj := gen.scope.Lookup(name); obj.Exported() {
			switch obj := obj.(type) {
			case *types.Const:
				val := obj.Val()
				var conv1, conv2 string
				if t, ok := obj.Type().(*types.Basic); ok && t.Info()&types.IsUntyped != 0 {
					// untyped constants have arbitrary precision... they may overflow integers.
					// this is just an approximation, use Package.Untypeds for exact value
					if val.Kind() == constant.Int {
						str := val.ExactString()
						conv1, conv2 = detectIntKind(gen.output, gen.path, name, str)
					}
				}
				d.header()
				fmt.Fprintf(gen.out, "\n\t\t\t%q:\t%sValueOf(%s%s%s%s),",
					name, gen.reflect, conv1, gen.name_, name, conv2)
			case *types.Var:
				d.header()
				fmt.Fprintf(gen.out, "\n\t\t\t%q:\t%sValueOf(&%s%s).Elem(),", name, gen.reflect, gen.name_, name)
			case *types.Func:
				if isGenericFunc(obj) {
					gen.output.Warnf("skipping import of %v:\timporting generic functions is not supported yet", name)
				} else {
					d.header()
					fmt.Fprintf(gen.out, "\n\t\t\t%q:\t%sValueOf(%s%s),",
						name, gen.reflect, gen.name_, name)
				}
			}
		}
	}
	d.footer()
}

func (gen *genimport) writeTypes() {
	d := gen.mapdecl("Types: map[string]%sType")

	for _, name := range gen.names {
		if obj := gen.scope.Lookup(name); obj.Exported() {
			switch obj := obj.(type) {
			case *types.TypeName:
				if isGenericType(obj.Type()) {
					gen.output.Warnf("skipping import of %s.%s:\timporting generic types is not supported yet",
						obj.Pkg().Path(), obj.Name())
				} else {
					d.header()
					fmt.Fprintf(gen.out, "\n\t\t\t%q:\t%sTypeOf((*%s%s)(nil)).Elem(),",
						name, gen.reflect, gen.name_, name)
				}
			}
		}
	}
	d.footer()
}

func isGenericFunc(t *types.Func) bool {
	sig, ok := t.Type().(*types.Signature)
	return !ok || isGenericType(sig)
}

func isGenericType(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		if t.TypeParams() != nil && t.TypeArgs() == nil {
			return true
		}
		return isGenericType(t.Underlying())
	case *types.Signature:
		return t.RecvTypeParams() != nil || t.TypeParams() != nil
	case *types.Interface:
		return t.IsImplicit() || !t.IsMethodSet()
	case *types.TypeParam:
		return true
	default:
		return false
	}
}

func (gen *genimport) writeProxies() {
	d := gen.mapdecl("Proxies: map[string]%sType")

	for _, name := range gen.names {
		if obj := gen.scope.Lookup(name); obj.Exported() {
			if t := extractInterface(obj, true); t != nil {
				d.header()
				fmt.Fprintf(gen.out, "\n\t\t\t%q:\t%sTypeOf((*%s%s)(nil)).Elem(),",
					name, gen.reflect, gen.proxyprefix, name)
			}
		}
	}
	d.footer()
}

func (gen *genimport) writeUntypeds() {
	d := gen.mapdecl("Untypeds: map[string]string")

	for _, name := range gen.names {
		if obj := gen.scope.Lookup(name); obj.Exported() {
			switch obj := obj.(type) {
			case *types.Const:
				if t, ok := obj.Type().(*types.Basic); ok && t.Info()&types.IsUntyped != 0 {
					kind := untyped.GoUntypedToKind(t.Kind())
					str := untyped.Marshal(kind, obj.Val())
					if len(str) != 0 {
						d.header()
						fmt.Fprintf(gen.out, "\n\t\t\t%q:\t%q,", name, str)
					}
				}
			}
		}
	}
	d.footer()
}

// find wrapper methods and write them. needed for accurate method selection.
func (gen *genimport) writeWrappers() {
	d := gen.mapdecl("Wrappers: map[string][]string")

	for _, name := range gen.names {
		if obj := gen.scope.Lookup(name); obj.Exported() {
			switch obj.(type) {
			case *types.TypeName:
				if t, ok := obj.Type().(*types.Named); ok {
					// only structs can have embedded fields, and thus wrapper methods for embedded fields
					if _, ok := t.Underlying().(*types.Struct); ok {
						wrappers := new(analyzer).Analyze(t)
						if len(wrappers) != 0 {
							d.header()
							fmt.Fprintf(gen.out, "\n\t\t\t%q:\t[]string{", obj.Name())
							for _, wrapper := range wrappers {
								fmt.Fprintf(gen.out, "%q,", wrapper)
							}
							fmt.Fprint(gen.out, "},")
						}
					}
				}
			}
		}
	}
	d.footer()
}

// write proxies that pre-implement package's interfaces
func (gen *genimport) writeInterfaceProxies() {
	path := gen.gpkg.Path()
	for _, name := range gen.names {
		obj := gen.scope.Lookup(name)
		if t := extractInterface(obj, true); t != nil {
			gen.writeInterfaceProxy(path, name, t)
		}
	}
}

func detectIntKind(o *Output, path, name, str string) (string, string) {
	i, err := strconv.ParseInt(str, 0, 64)
	if err == nil {
		if i == int64(int32(i)) {
			// constant fits int32. We can use the default (i.e. int)
			// on both 32-bit and 64-bit platforms
			return "", ""
		} else if i == int64(uint32(i)) {
			// constant fits uint32
			return "uint32(", ")"
		} else {
			return "int64(", ")"
		}
	}
	_, err = strconv.ParseUint(str, 0, 64)
	if err == nil {
		return "uint64(", ")"
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		// nothing fits... leave the default
		return "", ""
	} else {
		prefix := "float64"
		f = math.Abs(f)
		if f == float64(float32(f)) && f <= math.MaxFloat32 && f >= math.SmallestNonzeroFloat32 {
			// float32 loses no precision vs. float64
			prefix = "float32"
		}
		o.Warnf("package %q: integer constant %s = %s overflows both int64 and uint64, converting to %s",
			path, name, str, prefix)
		return prefix + "(", ")"
	}
}
