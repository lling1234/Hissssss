package node

type Config struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	TTL  int64  `yaml:"ttl"`
	HTTP string `yaml:"http"`
}
