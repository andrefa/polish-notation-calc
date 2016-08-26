package main

import (
	"github.com/DavidHuie/quartz/go/quartz"
	"github.com/irlndts/go-rpn"
	"regexp"
)

type Resolver struct{}

type CalcArgs struct {
	Expression string
}

type CalcResponse struct {
	Result float64
}

func (r *Resolver) Calc(args CalcArgs, response *CalcResponse) error {
	*response = CalcResponse{}

	if ok,_ := regexp.MatchString("^[0-9/*-+. ]+$", args.Expression); !ok {
		panic("Error: invalid expression " + args.Expression + " : ")
	}

	result,err := rpn.Calc(args.Expression)

	if err != nil {
		panic(err)
	}

	response.Result = result

	return nil
}

func main() {
	resolver := &Resolver{}
	quartz.RegisterName("resolver", resolver)
	quartz.Start()
}
