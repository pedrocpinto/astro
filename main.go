package astro

import (
	"fmt"

	"github.com/pedrocpinto/astro/pkg/constants"
)

func PrintSomething() {
	fmt.Println("Printing something as promised")
}

func GetAstroConstant(constName string) (*constants.AstroConstant, error) {
	return constants.Get(constName)
}