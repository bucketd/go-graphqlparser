package parser

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"unsafe"

	"github.com/bucketd/go-graphqlparser/ast"
	"github.com/bucketd/go-graphqlparser/lexer"
	"github.com/bucketd/go-graphqlparser/token"
)

type Parser struct {
	lexer *lexer.Lexer
	token lexer.Token
}

func New(input []byte) *Parser {
	return &Parser{
		lexer: lexer.New(input),
	}
}

func (p *Parser) Parse() (ast.Document, error) {
	var document ast.Document

	p.scan()

	for {
		definition, err := p.parseDefinition()
		if err != nil {
			return document, err
		}

		document.Definitions = append(document.Definitions, definition)

		if p.peek(token.Illegal) {
			return document, p.unexpected(p.token, p.expected(token.EOF))
		}

		if p.peek(token.EOF) {
			return document, nil
		}
	}
}

func (p *Parser) parseDefinition() (ast.Definition, error) {
	var definition ast.Definition
	var err error

	// ExecutableDefinition...
	if p.peek(token.Name, "query", "mutation", "subscription") || p.peek(token.Punctuator, "{") {
		definition.ExecutableDefinition, err = p.parseOperationDefinition(p.token.Literal == "{")
		return definition, err
	}

	// ExecutableDefinition...
	if p.peek(token.Name, "fragment") {
		// TODO(seeruk): Implement.
	}

	return definition, p.unexpected(p.token,
		p.expected(token.Name, "query", "mutation", "fragment"),
		p.expected(token.Punctuator, "{"),
	)

}

func (p *Parser) parseOperationDefinition(isQuery bool) (ast.ExecutableDefinition, error) {
	var definition ast.ExecutableDefinition
	var variableDefinitions []ast.VariableDefinition
	var directives []ast.Directive

	var name string
	var err error

	opType := ast.OperationTypeQuery

	if !isQuery {
		opType, err = p.parseOperationType()
		if err != nil {
			return definition, err
		}

		if tok, ok := p.consume(token.Name); ok {
			name = tok.Literal
		}

		variableDefinitions, err = p.parseVariableDefinitions()
		if err != nil {
			return definition, err
		}

		directives, err = p.parseDirectives()
		if err != nil {
			return definition, err
		}
	}

	selectionSet, err := p.parseSelectionSet(false)
	if err != nil {
		return definition, err
	}

	return ast.ExecutableDefinition{
		Kind:                ast.DefinitionKindOperation,
		OperationType:       opType,
		Name:                name,
		VariableDefinitions: variableDefinitions,
		Directives:          directives,
		SelectionSet:        selectionSet,
	}, nil
}

func (p *Parser) parseOperationType() (ast.OperationType, error) {
	tok, err := p.mustConsume(token.Name, "query", "mutation")
	if err != nil {
		return -1, err
	}

	if tok.Literal == "query" {
		return ast.OperationTypeQuery, nil
	}

	// Only other thing it can be at this point...
	return ast.OperationTypeMutation, nil
}

func (p *Parser) parseVariableDefinitions() ([]ast.VariableDefinition, error) {
	var definitions []ast.VariableDefinition

	if _, err := p.mustConsume(token.Punctuator, "("); err != nil {
		return definitions, err
	}

	for {
		if _, err := p.mustConsume(token.Punctuator, "$"); err != nil {
			return definitions, err
		}

		tok, err := p.mustConsume(token.Name)
		if err != nil {
			return definitions, err
		}

		definition := ast.VariableDefinition{}
		definition.Name = tok.Literal

		if _, err := p.mustConsume(token.Punctuator, ":"); err != nil {
			return definitions, err
		}

		definition.Type, err = p.parseType()
		if err != nil {
			return definitions, err
		}

		definition.DefaultValue, err = p.parseDefaultValue()
		if err != nil {
			return definitions, err
		}

		definitions = append(definitions, definition)

		if p.peek(token.Punctuator, ")") {
			break
		}
	}

	if _, err := p.mustConsume(token.Punctuator, ")"); err != nil {
		return definitions, err
	}

	return definitions, nil
}

func (p *Parser) parseDirectives() ([]ast.Directive, error) {
	var directives []ast.Directive

	for p.peek(token.Punctuator, "@") {
		directive, err := p.parseDirective()
		if err != nil {
			return directives, err
		}

		directives = append(directives, directive)
	}

	return directives, nil
}

func (p *Parser) parseDirective() (ast.Directive, error) {
	var directive ast.Directive

	_, err := p.mustConsume(token.Punctuator, "@")
	if err != nil {
		return directive, err
	}

	name, err := p.mustConsume(token.Name)
	if err != nil {
		return directive, err
	}

	args, err := p.parseArguments()
	if err != nil {
		return directive, err
	}

	directive.Name = name.Literal
	directive.Arguments = args

	return directive, nil
}

func (p *Parser) parseSelectionSet(optional bool) ([]ast.Selection, error) {
	var selectionSet []ast.Selection

	if optional && !p.skip(token.Punctuator, "{") {
		return selectionSet, nil
	}

	if !optional && !p.skip(token.Punctuator, "{") {
		return selectionSet, p.unexpected(p.token, p.expected(token.Name))
	}

	for {
		selection, err := p.parseSelection()
		if err != nil {
			return selectionSet, err
		}

		selectionSet = append(selectionSet, selection)

		if p.peek(token.Punctuator, "}") {
			break
		}
	}

	_, err := p.mustConsume(token.Punctuator, "}")
	if err != nil {
		return selectionSet, err
	}

	return selectionSet, nil
}

func (p *Parser) parseSelection() (ast.Selection, error) {
	var selection ast.Selection

	if p.peek(token.Punctuator, "...") {
		return selection, p.unexpected(p.token, "not yet implemented")
	}

	field, err := p.parseField()
	if err != nil {
		return selection, err
	}

	selection.Field = field

	return selection, nil
}

func (p *Parser) parseField() (ast.Field, error) {
	var field ast.Field

	name, err := p.mustConsume(token.Name)
	if err != nil {
		return field, err
	}

	if p.skip(token.Punctuator, ":") {
		field.Alias = name.Literal

		name, err = p.mustConsume(token.Name)
		if err != nil {
			return field, err
		}
	}

	field.Name = name.Literal

	field.Arguments, err = p.parseArguments()
	if err != nil {
		return field, err
	}

	field.Directives, err = p.parseDirectives()
	if err != nil {
		return field, err
	}

	field.SelectionSet, err = p.parseSelectionSet(true)
	if err != nil {
		return field, err
	}

	return field, nil
}

func (p *Parser) parseArguments() ([]ast.Argument, error) {
	var arguments []ast.Argument

	if !p.skip(token.Punctuator, "(") {
		return arguments, nil
	}

	for !p.skip(token.Punctuator, ")") {
		argument, err := p.parseArgument()
		if err != nil {
			return arguments, err
		}

		arguments = append(arguments, argument)
	}

	return arguments, nil
}

func (p *Parser) parseArgument() (ast.Argument, error) {
	var argument ast.Argument

	name, err := p.mustConsume(token.Name)
	if err != nil {
		return argument, err
	}

	_, err = p.mustConsume(token.Punctuator, ":")
	if err != nil {
		return argument, err
	}

	value, err := p.parseValue()
	if err != nil {
		return argument, err
	}

	argument.Name = name.Literal
	argument.Value = value

	return argument, nil
}

func (p *Parser) parseDefaultValue() (*ast.Value, error) {
	if !p.skip(token.Punctuator, "=") {
		return nil, nil
	}

	val, err := p.parseValue()
	if err != nil {
		return nil, err
	}

	return &val, nil
}

func (p *Parser) parseValue() (ast.Value, error) {
	if p.skip(token.Punctuator, "$") {
		return p.parseVariableValue()
	}

	if tok, ok := p.consume(token.IntValue); ok {
		return p.parseIntValue(tok)
	}

	if tok, ok := p.consume(token.FloatValue); ok {
		return p.parseFloatValue(tok)
	}

	if tok, ok := p.consume(token.StringValue); ok {
		return p.parseStringValue(tok)
	}

	if tok, ok := p.consume(token.Name, "true", "false"); ok {
		return p.parseBooleanValue(tok)
	}

	if p.skip(token.Name, "null") {
		return p.parseNullValue()
	}

	if tok, ok := p.consume(token.Name); ok {
		return p.parseEnumValue(tok)
	}

	if p.skip(token.Punctuator, "[") {
		return p.parseListValue()
	}

	if p.skip(token.Punctuator, "{") {
		return p.parseObjectValue()
	}

	return ast.Value{}, errors.New("TODO: see `parseDefinition`")
}

func (p *Parser) parseVariableValue() (ast.Value, error) {
	tok, err := p.mustConsume(token.Name)
	if err != nil {
		return ast.Value{}, err
	}

	return ast.Value{
		Kind:          ast.ValueKindVariable,
		VariableValue: tok.Literal,
	}, nil
}

func (p *Parser) parseIntValue(tok lexer.Token) (ast.Value, error) {
	iv, err := strconv.Atoi(tok.Literal)
	if err != nil {
		return ast.Value{}, err
	}

	return ast.Value{
		Kind:     ast.ValueKindIntValue,
		IntValue: iv,
	}, nil
}

func (p *Parser) parseFloatValue(tok lexer.Token) (ast.Value, error) {
	fv, err := strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		return ast.Value{}, err
	}

	return ast.Value{
		Kind:       ast.ValueKindFloatValue,
		FloatValue: fv,
	}, nil
}

func (p *Parser) parseStringValue(tok lexer.Token) (ast.Value, error) {
	return ast.Value{
		Kind:        ast.ValueKindStringValue,
		StringValue: tok.Literal,
	}, nil
}

func (p *Parser) parseBooleanValue(tok lexer.Token) (ast.Value, error) {
	return ast.Value{
		Kind:         ast.ValueKindBooleanValue,
		BooleanValue: tok.Literal == "true",
	}, nil
}

func (p *Parser) parseNullValue() (ast.Value, error) {
	return ast.Value{
		Kind: ast.ValueKindNullValue,
	}, nil
}

func (p *Parser) parseEnumValue(tok lexer.Token) (ast.Value, error) {
	return ast.Value{
		Kind:      ast.ValueKindEnumValue,
		EnumValue: tok.Literal,
	}, nil
}

func (p *Parser) parseListValue() (ast.Value, error) {
	list := ast.Value{}
	list.Kind = ast.ValueKindListValue

	for !p.skip(token.Punctuator, "]") {
		val, err := p.parseValue()
		if err != nil {
			return list, err
		}

		list.ListValue = append(list.ListValue, val)
	}

	return list, nil
}

func (p *Parser) parseObjectValue() (ast.Value, error) {
	object := ast.Value{}
	object.Kind = ast.ValueKindObjectValue

	for !p.skip(token.Punctuator, "}") {
		field, err := p.parseObjectField()
		if err != nil {
			return object, err
		}

		object.ObjectValue = append(object.ObjectValue, field)
	}

	return object, nil
}

func (p *Parser) parseObjectField() (ast.ObjectField, error) {
	var field ast.ObjectField

	tok, err := p.mustConsume(token.Name)
	if err != nil {
		return field, err
	}

	_, err = p.mustConsume(token.Punctuator, ":")
	if err != nil {
		return field, err
	}

	value, err := p.parseValue()
	if err != nil {
		return field, err
	}

	field.Name = tok.Literal
	field.Value = value

	return field, nil
}

func (p *Parser) parseType() (ast.Type, error) {
	var astType ast.Type

	// If we hit an opening square brace, we've got a list type, time to dive in.
	if p.skip(token.Punctuator, "[") {
		astType.Kind = ast.TypeKindListType

		itemType, err := p.parseType()
		if err != nil {
			return astType, nil
		}

		astType.ListType = &itemType

		if _, err := p.mustConsume(token.Punctuator, "]"); err != nil {
			return astType, err
		}
	} else {
		astType.Kind = ast.TypeKindNamedType

		tok, err := p.mustConsume(token.Name)
		if err != nil {
			return astType, err
		}

		astType.NamedType = tok.Literal
	}

	if p.skip(token.Punctuator, "!") {
		astType.NonNullable = true
	}

	return astType, nil
}

// Parser utilities:

func (p *Parser) consume(t token.Type, ls ...string) (lexer.Token, bool) {
	tok := p.token
	if tok.Type != t {
		return tok, false
	}

	if len(ls) == 0 {
		p.scan()
		return tok, true
	}

	for _, l := range ls {
		if tok.Literal != l {
			continue
		}

		p.scan()
		return tok, true
	}

	return tok, false
}

func (p *Parser) mustConsume(t token.Type, ls ...string) (lexer.Token, error) {
	tok, ok := p.consume(t, ls...)
	if !ok {
		return tok, p.unexpected(tok, p.expected(t, ls...))
	}

	return tok, nil
}

func (p *Parser) peek(t token.Type, ls ...string) bool {
	if p.token.Type != t {
		return false
	}

	if len(ls) == 0 {
		return true
	}

	for _, l := range ls {
		if p.token.Literal == l {
			return true
		}
	}

	return false
}

func (p *Parser) skip(t token.Type, ls ...string) bool {
	match := p.peek(t, ls...)
	if !match {
		return false
	}

	p.scan()

	return true
}

func (p *Parser) scan() {
	p.token = p.lexer.Scan()
}

func (p *Parser) expected(t token.Type, ls ...string) string {
	buf := bytes.Buffer{}
	buf.WriteString(t.String())
	buf.WriteString(" '")
	buf.WriteString(strings.Join(ls, "|"))
	return buf.String()
}

// TODO(Luke-Vear): think over the readability of the punctuation and caps.
func (p *Parser) unexpected(token lexer.Token, wants ...string) error {
	if len(wants) == 0 {
		wants = []string{"N/A"}
	}

	buf := bytes.Buffer{}
	buf.WriteString("parser error: unexpected token found at ")
	buf.WriteString("line: ")
	buf.WriteString(strconv.Itoa(token.Line))
	buf.WriteString(", column: ")
	buf.WriteString(strconv.Itoa(token.Position))
	buf.WriteString(". Found: ")
	buf.WriteString(token.Type.String())
	buf.WriteString(" '")
	buf.WriteString(token.Literal)
	buf.WriteString("'. Wanted: ")
	for i, want := range wants {
		buf.WriteString(want)
		if i < len(wants)-1 {
			buf.WriteString("' or ")
		}
	}
	buf.WriteString("'.")

	return errors.New(btos(buf.Bytes()))
}

// btos takes the given bytes, and turns them into a string.
// Q: naming btos or bbtos? :D
// TODO(seeruk): Is this actually portable then?
func btos(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}