package constants

import (
	"errors"
	"fmt"
)

const (
	GaussianGravitational                 = "gaussian_gravitational_constant"
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
		name:  "Light Year",
		value: "9.46073044725808 * 10^15",
		units: "meters",
	},
	GaussianGravitational: {
		name: "Gaussian Gravitational",
		value: "0.01720209895",
		units: "rad/day",
	},
	SpeedOfLight: {
		name: "Speed of light",
		value: "299792458",
		units: "meters per second",
	},
	Parsec: {
		name: "Parsec",
		value: "3.26",
		units: "light years",
	},
}

func Get(constant string) (*AstroConstant, error) {
	if val, ok := astronomicalConstants[constant]; ok {
		return &val, nil
	}

	return nil, errors.New(fmt.Sprintf("Sorry, info for constant '%s' is not configured.", constant))
}
