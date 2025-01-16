package data

var (
	configFilePath string = "~/.config/configgus/config.json"
)

// Config defines the structure of the configgus config file
type Config struct {
	MainDirectoryPath string
	FilesData         []FileData
}

// Load tries to read the configgus config file and init the struct
func (c *Config) Load() {

}

// Save write to the configgus config file the current state of the struct
func (c *Config) Save() {

}
