package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

type typeInfo struct {
	realTypeName string
	isPointer    bool
}

var (
	packageName string

	ti = map[string]typeInfo{
		"Arguments":                   {"", true},
		"ArgumentsDefinition":         {"InputValueDefinitions", true},
		"Definitions":                 {"", true},
		"DirectiveDefinition":         {"", true},
		"DirectiveLocations":          {"", true},
		"Directives":                  {"", true},
		"EnumTypeExtension":           {"TypeExtension", true},
		"EnumValuesDefinition":        {"EnumValueDefinitions", true},
		"ExecutableDefinition":        {"", true},
		"FieldsDefinition":            {"FieldDefinitions", true},
		"FragmentDefinition":          {"ExecutableDefinition", true},
		"ImplementsInterfaces":        {"Types", true},
		"InputFieldsDefinition":       {"InputValueDefinitions", true},
		"InputObjectTypeExtension":    {"TypeExtension", true},
		"InterfaceTypeExtension":      {"TypeExtension", true},
		"ObjectTypeExtension":         {"TypeExtension", true},
		"OperationDefinition":         {"ExecutableDefinition", true},
		"ScalarTypeExtension":         {"TypeExtension", true},
		"SchemaDefinition":            {"", true},
		"SchemaExtension":             {"", true},
		"Selections":                  {"", true},
		"TypeDefinition":              {"", true},
		"TypeDefinitionEnum":          {"TypeDefinition", true},
		"TypeDefinitionInputObject":   {"TypeDefinition", true},
		"TypeDefinitionInterface":     {"TypeDefinition", true},
		"TypeDefinitionObject":        {"TypeDefinition", true},
		"TypeDefinitionScalar":        {"TypeDefinition", true},
		"TypeDefinitionUnion":         {"TypeDefinition", true},
		"TypeExtension":               {"TypeExtension", true},
		"TypeSystemDefinition":        {"", true},
		"TypeSystemExtension":         {"", true},
		"UnionMemberTypes":            {"Types", true},
		"UnionTypeExtension":          {"TypeExtension", true},
		"VariableDefinitions":         {"", true},
		"Argument":                    {"", false},
		"Definition":                  {"", false},
		"Description":                 {"string", false},
		"Directive":                   {"", false},
		"Document":                    {"", false},
		"EnumValueDefinition":         {"", false},
		"FieldDefinition":             {"", false},
		"FieldSelection":              {"Selection", false},
		"FragmentSpread":              {"Selection", false},
		"InlineFragment":              {"Selection", false},
		"InputValueDefinition":        {"", false},
		"OperationTypeDefinition":     {"", false},
		"RootOperationTypeDefinition": {"", false},
		"Selection":                   {"", false},
		"Type":                        {"", false},
		"Value":                       {"", false},
		"VariableDefinition":          {"", false},
	}
)

var header = `
// Code generated by lab/walkergen
// DO NOT EDIT!
`

var walkerHead = `
// Walker holds event handlers for entering and leaving AST nodes.
type Walker struct {
`

var walkerBottom = `
}

// NewWalker returns a *Walker.
func NewWalker() *Walker {
	return &Walker{}
}
`

func main() {
	flag.StringVar(&packageName, "package", "", "The package name to use in the generated code.")
	flag.Parse()

	tns := make([]string, 0, len(ti))
	for tn := range ti {
		tns = append(tns, tn)
	}
	sort.Strings(tns)

	// Header and package name
	fmt.Fprintf(os.Stdout, strings.TrimSpace(header))
	fmt.Fprintf(os.Stdout, "\npackage %s\n", packageName)

	// Walker head
	fmt.Fprintf(os.Stdout, walkerHead+"\t")

	// Walker middle
	for i, tn := range tns {
		handlers.Execute(os.Stdout, map[string]string{
			"TypeNameLCF": lcfirst(tn),
			"TypeName":    tn,
		})
		if i < len(tns) {
			fmt.Fprintf(os.Stdout, "\n\t")
		}
	}

	// Walker bottom
	fmt.Fprintf(os.Stdout, walkerBottom)

	// Event handers
	for _, tn := range tns {
		eventHandlers.Execute(os.Stdout, map[string]string{
			"TypeNameLCF":  lcfirst(tn),
			"TypeName":     tn,
			"RealTypeName": realTypeName(tn),
			"AbridgedTN":   sane(strings.Map(abridger, tn)),
			"Pointer":      ptrIf(ti[tn].isPointer),
		})
	}
}

var handlers = template.Must(template.New("handlers").Parse(strings.TrimSpace(`
	{{.TypeNameLCF}}EventHandlers {{.TypeName}}EventHandlers
`)))

var eventHandlers = template.Must(template.New("eventHandlers").Parse(`
// {{.TypeName}}EventHandler function can handle enter/leave events for {{.TypeName}}.
type {{.TypeName}}EventHandler func({{.Pointer}}{{.RealTypeName}})

// {{.TypeName}}EventHandlers stores the enter and leave events handlers.
type {{.TypeName}}EventHandlers struct {
	enter []{{.TypeName}}EventHandler
	leave []{{.TypeName}}EventHandler
}

// Add{{.TypeName}}EnterEventHandler adds an event handler to be called when entering a {{.TypeName}} node.
func (w *Walker) Add{{.TypeName}}EnterEventHandler(handler {{.TypeName}}EventHandler) {
	w.{{.TypeNameLCF}}EventHandlers.enter = append(w.{{.TypeNameLCF}}EventHandlers.enter, handler)
}

// Add{{.TypeName}}EnterEventHandler adds an event handler to be called when leaving a {{.TypeName}} node.
func (w *Walker) Add{{.TypeName}}LeaveEventHandler(handler {{.TypeName}}EventHandler) {
	w.{{.TypeNameLCF}}EventHandlers.leave = append(w.{{.TypeNameLCF}}EventHandlers.leave, handler)
}

// On{{.TypeName}}Enter calls the enter event handlers for this node type.
func (w *Walker) On{{.TypeName}}Enter({{.AbridgedTN}} {{.Pointer}}{{.RealTypeName}}) {
	for _, handler := range w.{{.TypeNameLCF}}EventHandlers.enter {
		handler({{.AbridgedTN}})
	}
}

// On{{.TypeName}}Leave calls the leave event handlers for this node type.
func (w *Walker) On{{.TypeName}}Leave({{.AbridgedTN}} {{.Pointer}}{{.RealTypeName}}) {
	for _, handler := range w.{{.TypeNameLCF}}EventHandlers.leave {
		handler({{.AbridgedTN}})
	}
}
`))

func lcfirst(in string) string {
	if len(in) == 0 {
		return in
	}

	if len(in) == 1 {
		return strings.ToLower(in)
	}

	fr, w := utf8.DecodeRuneInString(in)

	return strings.ToLower(string(fr)) + in[w:]
}

func abridger(r rune) rune {
	if unicode.IsUpper(r) {
		return unicode.ToLower(r)
	}
	return -1
}

func in(item string, list []string) bool {
	for _, thing := range list {
		if thing == item {
			return true
		}
	}
	return false
}

func ptrIf(isPtr bool) string {
	if isPtr {
		return "*"
	}
	return ""
}

func sane(typeName string) string {
	switch typeName {
	case "type":
		return "t"
	case "error":
		return "err"
	case "if":
		return "ilf"
	}
	return typeName
}

func realTypeName(tn string) string {
	if typeInfo := ti[tn]; len(typeInfo.realTypeName) > 0 {
		return typeInfo.realTypeName
	}
	return tn
}
