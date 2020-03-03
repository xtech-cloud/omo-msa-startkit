package config

type Service_ struct {
	TTL      int64  `yaml:ttl`
	Interval int64  `yaml:interval`
	Address  string `yaml:address`
}

type ConfigSchema_ struct {
	Service Service_ `yaml:service`
}
