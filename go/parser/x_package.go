// this file was generated by gomacro command: import _i "github.com/WilliamNHarvey/gomacro/go/parser"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package parser

import (
	r "reflect"

	"github.com/WilliamNHarvey/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/WilliamNHarvey/gomacro/go/parser"
func init() {
	imports.Packages["github.com/WilliamNHarvey/gomacro/go/parser"] = imports.Package{
		Binds: map[string]r.Value{
			"AllErrors":         r.ValueOf(AllErrors),
			"DeclarationErrors": r.ValueOf(DeclarationErrors),
			"ImportsOnly":       r.ValueOf(ImportsOnly),
			"MakeQuote":         r.ValueOf(MakeQuote),
			"PackageClauseOnly": r.ValueOf(PackageClauseOnly),
			"ParseComments":     r.ValueOf(ParseComments),
			"SpuriousErrors":    r.ValueOf(SpuriousErrors),
			"Trace":             r.ValueOf(Trace),
		}, Types: map[string]r.Type{
			"Mode":   r.TypeOf((*Mode)(nil)).Elem(),
			"Parser": r.TypeOf((*Parser)(nil)).Elem(),
		}, Wrappers: map[string][]string{
			"Parser": []string{"Configure", "Init", "Parse"},
		},
	}
}
