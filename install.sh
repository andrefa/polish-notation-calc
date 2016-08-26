#!/bin/sh

# install ruby dependencies
gem install quartz

# install go dependencies
go get github.com/DavidHuie/quartz/go/quartz
go get github.com/irlndts/go-rpn
