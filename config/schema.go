package config

type Service_ struct {
    Name string  `yaml:name`
	TTL      int64  `yaml:ttl`
	Interval int64  `yaml:interval`
	Address  string `yaml:address`
}

type Logger_ struct {
	Level string `yaml:level`
}

type ConfigSchema_ struct {
	Service Service_ `yaml:service`
	Logger  Logger_  `yaml:logger`
}
