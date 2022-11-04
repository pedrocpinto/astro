package main

import (
	"fmt"

	"github.com/pedrocpinto/astro/pkg/constants"
)

func main() {
	astroConst, err := constants.GetConstant(constants.LighYear)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(astroConst)
}
