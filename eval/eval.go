package eval

import (
	"fmt"
	"log"
	"strings"

	"github.com/exfly/newlang/ast"
	"github.com/exfly/newlang/object"
)

var (
	// TRUE is a cached Boolean object holding the `true` value
	TRUE = &object.Boolean{Value: true}

	// FALSE is a cached Boolean object holding the `false` value
	FALSE = &object.Boolean{Value: false}

	// NULL is a cached Null object
	NULL = &object.Null{}
)

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func fromNativeBoolean(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.PrefixExpression:

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.BindExpression:
		value := Eval(node.Value, env)
		if isError(value) {
			return value
		}

		if ident, ok := node.Left.(*ast.Identifier); ok {
			if immutable, ok := value.(object.Immutable); ok {
				env.Set(ident.Value, immutable.Clone())
			} else {
				env.Set(ident.Value, value)
			}

			return NULL
		}
		return newError("expected identifier on left got=%T", node.Left)

	case *ast.AssignmentExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		value := Eval(node.Value, env)
		if isError(value) {
			return value
		}

		if ident, ok := node.Left.(*ast.Identifier); ok {
			env.Set(ident.Value, value)
		} else if ie, ok := node.Left.(*ast.IndexExpression); ok {
			obj := Eval(ie.Left, env)
			if isError(obj) {
				return obj
			}

			if array, ok := obj.(*object.Array); ok {
				index := Eval(ie.Index, env)
				if isError(index) {
					return index
				}
				if idx, ok := index.(*object.Integer); ok {
					array.Elements[idx.Value] = value
				} else {
					return newError("cannot index array with %#v", index)
				}
			} else if hash, ok := obj.(*object.Hash); ok {
				key := Eval(ie.Index, env)
				if isError(key) {
					return key
				}
				if hashKey, ok := key.(object.Hashable); ok {
					hashed := hashKey.HashKey()
					hash.Pairs[hashed] = object.HashPair{Key: key, Value: value}
				} else {
					return newError("cannot index hash with %T", key)
				}
			} else {
				return newError("object type %T does not support item assignment", obj)
			}
		} else {
			return newError("expected identifier or index expression got=%T", left)
		}

		return NULL

	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Env: env, Body: body}
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.Return{Value: val}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	default:
		log.Printf("unknown node: %v", node)
	}
	return &object.Null{}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		if right.Type() == object.BOOLEAN {
			return evalBooleanPrefixOperatorExpression(operator, right)
		}
		return evalIntegerPrefixOperatorExpression(operator, right)
	case "~", "-":
		return evalIntegerPrefixOperatorExpression(operator, right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBooleanPrefixOperatorExpression(operator string, right object.Object) object.Object {
	if right.Type() != object.BOOLEAN {
		return newError("unknown operator: %s%s", operator, right.Type())
	}

	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalIntegerPrefixOperatorExpression(operator string, right object.Object) object.Object {
	if right.Type() != object.INTEGER {
		return newError("unknown operator: %s%s", operator, right.Type())
	}

	value := right.(*object.Integer).Value
	switch operator {
	case "!":
		return FALSE
	case "~":
		return &object.Integer{Value: ^value}
	case "-":
		return &object.Integer{Value: -value}
	default:
		return newError("unknown operator: %s", operator)
	}
}

func evalInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	switch {

	// {"a": 1} + {"b": 2}
	case operator == "+" && left.Type() == object.HASH && right.Type() == object.HASH:
		leftVal := left.(*object.Hash).Pairs
		rightVal := right.(*object.Hash).Pairs
		pairs := make(map[object.HashKey]object.HashPair)
		for k, v := range leftVal {
			pairs[k] = v
		}
		for k, v := range rightVal {
			pairs[k] = v
		}
		return &object.Hash{Pairs: pairs}

	// [1] + [2]
	case operator == "+" && left.Type() == object.ARRAY && right.Type() == object.ARRAY:
		leftVal := left.(*object.Array).Elements
		rightVal := right.(*object.Array).Elements
		elements := make([]object.Object, len(leftVal)+len(rightVal))
		elements = append(leftVal, rightVal...)
		return &object.Array{Elements: elements}

	// [1] * 3
	case operator == "*" && left.Type() == object.ARRAY && right.Type() == object.INTEGER:
		leftVal := left.(*object.Array).Elements
		rightVal := int(right.(*object.Integer).Value)
		elements := leftVal
		for i := rightVal; i > 1; i-- {
			elements = append(elements, leftVal...)
		}
		return &object.Array{Elements: elements}

	// 3 * [1]
	case operator == "*" && left.Type() == object.INTEGER && right.Type() == object.ARRAY:
		leftVal := int(left.(*object.Integer).Value)
		rightVal := right.(*object.Array).Elements
		elements := rightVal
		for i := leftVal; i > 1; i-- {
			elements = append(elements, rightVal...)
		}
		return &object.Array{Elements: elements}

	// " " * 4
	case operator == "*" && left.Type() == object.STRING && right.Type() == object.INTEGER:
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.Integer).Value
		return &object.String{Value: strings.Repeat(leftVal, int(rightVal))}

	// 4 * " "
	case operator == "*" && left.Type() == object.INTEGER && right.Type() == object.STRING:
		leftVal := left.(*object.Integer).Value
		rightVal := right.(*object.String).Value
		return &object.String{Value: strings.Repeat(rightVal, int(leftVal))}

	case operator == "==":
		return fromNativeBoolean(left.(object.Comparable).Compare(right) == 0)
	case operator == "!=":
		return fromNativeBoolean(left.(object.Comparable).Compare(right) != 0)
	case operator == "<=":
		return fromNativeBoolean(left.(object.Comparable).Compare(right) < 1)
	case operator == ">=":
		return fromNativeBoolean(left.(object.Comparable).Compare(right) > -1)
	case operator == "<":
		return fromNativeBoolean(left.(object.Comparable).Compare(right) == -1)
	case operator == ">":
		return fromNativeBoolean(left.(object.Comparable).Compare(right) == 1)

	case left.Type() == right.Type() && left.Type() == object.BOOLEAN:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() == right.Type() && left.Type() == object.INTEGER:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == right.Type() && left.Type() == object.STRING:
		return evalStringInfixExpression(operator, left, right)

	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value

	switch operator {
	case "&&":
		return fromNativeBoolean(leftVal && rightVal)
	case "||":
		return fromNativeBoolean(leftVal || rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "|":
		return &object.Integer{Value: leftVal | rightVal}
	case "^":
		return &object.Integer{Value: leftVal ^ rightVal}
	case "&":
		return &object.Integer{Value: leftVal & rightVal}
	case "<<":
		return &object.Integer{Value: leftVal << uint64(rightVal)}
	case ">>":
		return &object.Integer{Value: leftVal >> uint64(rightVal)}
	case "<":
		return fromNativeBoolean(leftVal < rightVal)
	case "<=":
		return fromNativeBoolean(leftVal <= rightVal)
	case ">":
		return fromNativeBoolean(leftVal > rightVal)
	case ">=":
		return fromNativeBoolean(leftVal >= rightVal)
	case "==":
		return fromNativeBoolean(leftVal == rightVal)
	case "!=":
		return fromNativeBoolean(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.ERROR {
				return result
			}
		}
	}

	return result
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	return newError("identifier not found: " + node.Value)
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {

	case *object.Function:
		env := extendFunctionEnv(fn, args)
		return unwrapReturnValue(Eval(fn.Body, env))

	case *object.Builtin:
		if result := fn.Fn(args...); result != nil {
			return result
		}
		return NULL

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := fn.Env.Clone()

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	}

	return obj
}
