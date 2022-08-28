package main

import (
	"bufio"
	"fmt"
	"github.com/netscotte/goSample/practise/calculator/stack"
	"os"
	"strconv"
)

func getInput() string {
	fmt.Println("input a express: ")
	buf := bufio.NewReader(os.Stdin)
	a, _ := buf.ReadString('\n')
	return a
}

func transPostExpress(express string) (postExpress []string) {
	// express = 2*(9+6/3-5)+4
	operationPriv := map[byte]int{
		'(': 0,
		')': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	var opStack stack.Stack

	// 对字符串进行遍历时，每一个元素为byte类型，双引号的对象为string类型，单引号的对象为byte类型
	for i := 0; i < len(express); {
		fmt.Printf("get element：%c, ", express[i])
		topElement, _ := opStack.Top()
		fmt.Printf("postExpress: %v, stackTop, %s\n", postExpress, topElement)
		//time.Sleep(1*time.Second)
		switch {
		// 遇到数字直接添加到到后缀表达式
		case express[i] >= '0' && express[i] <= '9':
			var number []byte
			for ; i < len(express); i++ {
				// 如果不是数字，证明数字扫描完毕，继续下一个扫描
				if express[i] < '0' || express[i] > '9' {
					break
				}
				number = append(number, express[i])
			}
			postExpress = append(postExpress, string(number))
		// 遇到左括号直接压栈
		case express[i] == '(':
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		// 遇到右括号，将栈中的所有操作符全都压到后缀表达式，直到遇到左括号
		case express[i] == ')':
			for !opStack.Empty() {
				op, err := opStack.Pop()
				if err != nil {
					fmt.Println("error get pop")
				}
				if op[0] != '(' {
					postExpress = append(postExpress, op)
				} else {
					break
				}
			}
			i++
		// 若遇到运算符，则比较
		case express[i] == '+' || express[i] == '-' || express[i] == '*' || express[i] == '/':
			// 栈顶为空，直接压入
			if opStack.Empty() {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue
			}
			// 与栈顶元素做比较
			topElement, err := opStack.Top()
			// topElement是一个string, 需要与byte进行比较，所以取其第一位
			if err != nil {
				fmt.Printf("error get top: %v", err)
			}
			// 若栈顶元素优先级高，则直接压入, 否则将栈顶元素弹出到后缀表达式，继续与新的栈顶元素进行比较
			if operationPriv[express[i]] > operationPriv[topElement[0]] {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue
			} else {
				data, err := opStack.Pop()
				if err != nil {
					fmt.Println("error pop ", err)
				}
				postExpress = append(postExpress, data)
			}
		default:
			fmt.Printf("ignore element %v of express\n", express[i])
			i++
		}
	}
	for !opStack.Empty() {
		value, err := opStack.Pop()
		if err != nil {
			fmt.Println("error, can't pop")
			break
		}
		postExpress = append(postExpress, value)
	}
	return postExpress
}

func calculate(postExpress []string) (result float64) {
	var s stack.Stack
	for _, value := range postExpress {
		if value == "+" || value == "-" || value == "*" || value == "/" {
			n1, err := s.Pop()
			if err != nil {
				fmt.Println("error no any thing in stack")
				return
			}
			n2, err := s.Pop()
			if err != nil {
				fmt.Println("error no anything in stack")
				return
			}
			num2, err := strconv.ParseFloat(n1, 64)
			num1, err := strconv.ParseFloat(n2, 64)
			var r1 float64
			switch value {
			case "+":
				r1 = num1 + num2
			case "-":
				r1 = num1 - num2
			case "*":
				r1 = num1 * num2
			case "/":
				r1 = num1 / num2
			default:
				fmt.Printf("unknown operation %v", value)
				break
			}

			s.Push(fmt.Sprintf("%f", r1))
		} else {
			s.Push(value)
		}
	}

	resultString, err := s.Pop()
	if err != nil {
		fmt.Printf("error get final result: %v", err)
	}
	result, err = strconv.ParseFloat(resultString, 64)
	return

}

func main() {

	a := getInput()
	fmt.Println("get input from console: ", a)
	postExpress := transPostExpress(a)
	fmt.Println("get postExpress: ", postExpress)
	result := calculate(postExpress)
	fmt.Println("get calculate result: ", result)
}
