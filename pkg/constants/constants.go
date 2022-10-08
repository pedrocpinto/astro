package constants

import (
	"errors"
	"fmt"
)

const (
	GaussianGravitationalConstant         = "gaussian_gravitational_constant"
	LighYear                              = "light_year"
	MeanRatioOfTheTTSecondToTheTCGSecond  = "mean_ratio_of_the_tt_second_to_the_tcg_second"
	MeanRatioOfTheTCBSecondToTheTDBSecond = "mean_ratio_of_the_tcb_second_to_the_tdb_second"
	Parsec                                = "parsec"
	SpeedOfLight                          = "speed_of_light"
)

type AstroConstant struct {
	name  string
	value string
	units string
}

var astronomicalConstants = map[string]AstroConstant{
	LighYear: {
		name: "Light Year",
		value: "9.46073044725808 * 10^15",
		units: "meters",
	},
}

func Get(constant string) (*AstroConstant, error) {
	if val, ok := astronomicalConstants[constant]; ok {
		return &val, nil
	}

	return nil, errors.New(fmt.Sprintf("Sorry, info for constant '%s' is not configured.", constant))
}
