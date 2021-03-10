package config

const defaultYAML string = `
service:
    name: omo.api.msa.startkit
    address: :7079
    ttl: 15
    interval: 10
logger:
    level: info
    dir: /var/log/msa/
`
