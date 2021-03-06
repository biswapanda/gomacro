// this file was generated by gomacro command: import _b "html/template"
// DO NOT EDIT! Any change will be lost when the file is re-generated

// +build go1.11

package go1_11

import (
	. "reflect"
	template "html/template"
)

// reflection: allow interpreted code to import "html/template"
func init() {
	Packages["html/template"] = Package{
	Binds: map[string]Value{
		"ErrAmbigContext":	ValueOf(template.ErrAmbigContext),
		"ErrBadHTML":	ValueOf(template.ErrBadHTML),
		"ErrBranchEnd":	ValueOf(template.ErrBranchEnd),
		"ErrEndContext":	ValueOf(template.ErrEndContext),
		"ErrNoSuchTemplate":	ValueOf(template.ErrNoSuchTemplate),
		"ErrOutputContext":	ValueOf(template.ErrOutputContext),
		"ErrPartialCharset":	ValueOf(template.ErrPartialCharset),
		"ErrPartialEscape":	ValueOf(template.ErrPartialEscape),
		"ErrPredefinedEscaper":	ValueOf(template.ErrPredefinedEscaper),
		"ErrRangeLoopReentry":	ValueOf(template.ErrRangeLoopReentry),
		"ErrSlashAmbig":	ValueOf(template.ErrSlashAmbig),
		"HTMLEscape":	ValueOf(template.HTMLEscape),
		"HTMLEscapeString":	ValueOf(template.HTMLEscapeString),
		"HTMLEscaper":	ValueOf(template.HTMLEscaper),
		"IsTrue":	ValueOf(template.IsTrue),
		"JSEscape":	ValueOf(template.JSEscape),
		"JSEscapeString":	ValueOf(template.JSEscapeString),
		"JSEscaper":	ValueOf(template.JSEscaper),
		"Must":	ValueOf(template.Must),
		"New":	ValueOf(template.New),
		"OK":	ValueOf(template.OK),
		"ParseFiles":	ValueOf(template.ParseFiles),
		"ParseGlob":	ValueOf(template.ParseGlob),
		"URLQueryEscaper":	ValueOf(template.URLQueryEscaper),
	}, Types: map[string]Type{
		"CSS":	TypeOf((*template.CSS)(nil)).Elem(),
		"Error":	TypeOf((*template.Error)(nil)).Elem(),
		"ErrorCode":	TypeOf((*template.ErrorCode)(nil)).Elem(),
		"FuncMap":	TypeOf((*template.FuncMap)(nil)).Elem(),
		"HTML":	TypeOf((*template.HTML)(nil)).Elem(),
		"HTMLAttr":	TypeOf((*template.HTMLAttr)(nil)).Elem(),
		"JS":	TypeOf((*template.JS)(nil)).Elem(),
		"JSStr":	TypeOf((*template.JSStr)(nil)).Elem(),
		"Srcset":	TypeOf((*template.Srcset)(nil)).Elem(),
		"Template":	TypeOf((*template.Template)(nil)).Elem(),
		"URL":	TypeOf((*template.URL)(nil)).Elem(),
	}, 
	}
}
