package main

import (
	"errors"
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

	if ok,_ := regexp.MatchString("^[0-9/*+. -]+$", args.Expression); !ok {
		return errors.New("Error: '" + args.Expression + "' is an invalid expression ")
	}

	result,err := rpn.Calc(args.Expression)

	if err != nil {
		return err
	}

	response.Result = result

	return nil
}

func main() {
	resolver := &Resolver{}
	quartz.RegisterName("resolver", resolver)
	quartz.Start()
}
