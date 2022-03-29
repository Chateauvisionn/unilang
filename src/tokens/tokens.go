package tokens

import (
	"strconv"

	"scythe.com/uni/src/util"
)

type Token int64
type Tokens struct {
	op       Token
	value    int64
	priority int8
}

// Enum of all operators
const (
	OP_PLUS Token = iota // addition
	OP_MIN               // minus
	OP_MUL               // multiplication
	OP_DIV               // division
	OP_PUSH
	OP_DUMP
	OP_DEFAULT
	L_PAREN
	R_PAREN
)

// Return the Operator of the "Token" struct
func (t *Tokens) GetOp() Token {
	return t.op
}

// Return the int value of the "Token" struct
func (t *Tokens) GetValue() int64 {
	return t.value
}

func (t *Tokens) GetPriority() int8 {
	return t.priority
}

// Return a Token with plus operator and value code 0
func Plus() Tokens {
	return Tokens{OP_PLUS, 0, 1}
}

func Default() Tokens {
	return Tokens{OP_DEFAULT, 0, -1}
}

// Return a Token with minus operator and value code 0
func Min() Tokens {
	return Tokens{OP_MIN, 0, 1}
}

// Return a Token with multiplication operator and value code 0
func Mul() Tokens {
	return Tokens{OP_MUL, 0, 2}
}

// Return a Token with division operator and value code 0
func Div() Tokens {
	return Tokens{OP_DIV, 0, 2}
}

// Return a Token with push operator and value passed in parameter
func Push(value int64) Tokens {
	return Tokens{OP_PUSH, value, 0}
}

// Return a Token with DUMP operator and value code 0
func Dump() Tokens {
	return Tokens{OP_DUMP, 0, 4}
}

func L_Paren() Tokens {
	return Tokens{L_PAREN, 0, 0}
}

func R_Paren() Tokens {
	return Tokens{R_PAREN, 0, 0}
}

func InfixToPostfix(arr []Tokens) []Tokens {
	operandStack := make([]Tokens, len(arr))
	postFixTerms := make([]Tokens, len(arr))

	termsIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i].op != OP_PUSH {
			operandStack = append(operandStack, arr[i])
		} else {
			postFixTerms[termsIndex] = arr[i]
			termsIndex++
		}
	}

	for j := 0; j < len(operandStack); j++ {
		postFixTerms[termsIndex] = util.Pop(&operandStack)
		termsIndex++
	}

	return postFixTerms
}

// Parse a string array and append for each chars an operator to an array. Then, return the array
func ParseTokenAsOperator(word string) Tokens {
	if word == "+" {
		return Plus()
	} else if word == "-" {
		return Min()
	} else if word == "*" {
		return Mul()
	} else if word == "/" {
		return Div()
	} else if word == "dmp" {
		return Dump()
	} else if word == "(" {
		return L_Paren()
	} else if word == ")" {
		return R_Paren()
	} else {
		num, _ := strconv.Atoi(word)
		return Push(int64(num))
	}
}
