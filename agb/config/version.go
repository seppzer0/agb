package config

var (
	// this data should be overriden during compilation
	appVersion = "v0.0.0"
	goVersion  = "v0.0.0"
)

// VersionConfig is a version info holder updated dynamically during compilation.
type VersionConfig struct {
	AppVersion string
	GoVersion  string
}

// NewVersionConfig returns new instance of VersionConfig.
func NewVersionConfig() *VersionConfig {
	return &VersionConfig{
		AppVersion: appVersion,
		GoVersion:  goVersion,
	}
}
