package config

type Logger_ struct {
	Level int `yaml:level`
}

type ConfigScheme_ struct {
	Logger Logger_ `yaml:logger`
}
