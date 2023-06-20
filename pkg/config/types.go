package config

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Backend    BackendConfig    `mapstructure:"backend"`
	Kubernetes KubernetesConfig `mapstructure:"kubernetes"`
}

/* SERVER */
type ServerConfig struct {
	Address string `mapstructure:"address"`
}

/* BACKEND */
type BackendConfig struct {
	Storage string               `mapstructure:"storage"`
	Dqlite  *DqliteBackendConfig `mapstructure:"dqlite"`
}

type DqliteBackendConfig struct {
	Port    int    `mapstructure:"port"`
	Datadir string `mapstructure:"datadir"`
}

/* KUBERNETES */
type KubernetesConfig struct {
	Server               string `mapstructure:"server"`
	CACert               string `mapstructure:"caCert"`
	FrontProxyClientCert string `mapstructure:"frontProxyClientCert"`
	FrontProxyClientKey  string `mapstructure:"frontProxyClientKey"`
}
