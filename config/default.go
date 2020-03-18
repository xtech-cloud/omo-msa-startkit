package config

const defaultYAML string = `
service:
    address: :8080
    ttl: 15
    interval: 10
logger:
    level: info
    dir: /var/log/msa/
`
