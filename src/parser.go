package main

import (
	"fmt"
	"strings"
)

var opers = map[string]ops{
	"STAR":  {6, "LEFT"},
	"PLUS":  {5, "LEFT"},
	"MIN":   {5, "LEFT"},
	"SLASH": {6, "LEFT"},
}

func NewPrsr() *Parser {
	return &Parser{
		pos: 0,
	}
}

func (p *Parser) parse(rl string) []cookie {
	//shunting yard algorrithm
	stack := []cookie{}  //symbols
	output := []cookie{} //numbers
	count := 0
	fmt.Println("rl : ", rl)
	for i := 0; i < len(p.tokens); i++ {
		fmt.Println("................ token is : ", p.tokens[count])
		switch {
		case p.tokens[count].isNumber():
			fmt.Println("----entered case : isNumber----")
			output = append(output, p.tokens[count])
			count++
		case p.tokens[count].isOperator():
			fmt.Println("----entered case : isOperator----")
			fmt.Println(" length of stack : ", len(stack))
			if len(stack) == 0 {
				stack = append(stack, p.tokens[count])
				count++
				break
			}
			if len(stack) >= 1 {
				fmt.Println("----entered : len(Stack >= 1)----")
				o1 := p.tokens[count]
				o2 := stack[len(stack)-1]

				//move the top of the stack to the ouput
				if o1.isHigherPrec(o2) {
					fmt.Println("----entered : o1.isHigherPrec(o2)----")
					fmt.Println("----stack before moval : ", stack, " ----")
					output = append(output, stack[0])
					fmt.Println("---- moveing stack item : ", stack[0], " in stack : ", stack, "----")
					stack = append(stack[:0], stack[0+1:]...)
					fmt.Println("----BEGIN | move top of stack to output | stack after  ----")
					fmt.Println(stack)
					fmt.Println("----END   | move top of stack to output | stack after  ----")
				}

			}
			stack = append(stack, p.tokens[count])

			count++
		case p.tokens[count].isComma():
			//loop this untill the top of the stack is a comma
			for {
				output = append(output, stack[len(stack)-1])
				stack = append(stack[:len(stack)-1], stack[len(stack)+1:]...)
			}
			count++
		case p.tokens[count].isLPAREN():
			stack = append(stack, p.tokens[count])
			count++
		case p.tokens[count].isRPAREN():
			holder := []cookie{}
			//loop this untill the top of the stack is a left parentheses
			for {
				//if it finds a left parentheses
				if stack[len(stack)-1].isLPAREN() {
					holder = append(holder, stack[len(stack)-1])
					stack = append(stack[:len(stack)-1], stack[len(stack)+1:]...)
					//if the top of the stack is a function operator
					if stack[len(stack)-1].isOperator() {
						output = append(output, holder[0])
						holder = append(holder[:0], holder[1:]...)
					}
				}
				output = append(output, stack[len(stack)-1])
				stack = append(stack[:len(stack)-1], stack[len(stack)+1:]...)

			}
			count++
		}

	}
	fmt.Println("stack at the end : ", stack)
	if len(stack) != 0 {
		fmt.Println("stack was not 0. it was : ", len(stack))
		for l, j := 0, len(stack)-1; l < j; l, j = l+1, j-1 {
			stack[l], stack[j] = stack[j], stack[l]
		}
		for m := 0; m <= len(stack)-1; m++ {
			fmt.Println("finally adding : ", stack[m], "to the output")
			output = append(output, stack[m])
		}
	}
	return output
}

func (p *Parser) Run(ln string) []cookie {
	tokens := lexicallyAnalize(ln)

	p.flush()
	p.tokens = tokens

	return p.parse(ln)
}

func (c *cookie) isRPAREN() bool {
	cc := strings.Split(c.t_sort, "_")[1]
	return cc == "RGT"
}

func (c *cookie) isLPAREN() bool {
	cc := strings.Split(c.t_sort, "_")[1]
	return cc == "LFT"
}

func (c *cookie) isNumber() bool {
	return c.t_sort == "INTERGER"
}

func (c *cookie) isOperator() bool {
	tok := strings.Split(c.t_sort, "_")[1]

	var ok bool
	if _, ok := opers[tok]; ok {

		return ok
	}
	return ok
}

func (c *cookie) isComma() bool {
	return strings.Split(c.t_sort, "_")[1] == "COMMA"
}

func (p *Parser) flush() {
	p.pos = 0
	p.tokens = nil
}

func (o1 cookie) isHigherPrec(o2 cookie) bool {
	o1o := strings.Split(o1.t_sort, "_")[1]
	o2o := strings.Split(o2.t_sort, "_")[1]

	if opers[o1o].assoc == "LEFT" && opers[o2o].prec >= opers[o1o].prec ||
		opers[o1o].assoc == "RIGHT" && opers[o2o].prec > opers[o1o].prec {
		return true
	}
	return false
}
