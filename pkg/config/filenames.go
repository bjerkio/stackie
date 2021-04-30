package config

type FileType int64

const (
	YAML FileType = iota
	JSON
)

type FileName struct {
	FileName string
	Type     FileType
}

// GetProjectFileNames represents project-specific configuration file names
func GetProjectFileNames() []FileName {
	return []FileName{
		{FileName: ".stackie.yml", Type: YAML},
		{FileName: ".stackie.json", Type: JSON},
		{FileName: "stackie.yml", Type: YAML},
		{FileName: "stackie.json", Type: JSON},
	}
}

// GetUserFileNames represents developer and project-specific configuration file names
func GetUserFileNames() []FileName {
	return []FileName{
		{FileName: ".stackie.local.yml", Type: YAML},
		{FileName: ".stackie.local.json", Type: JSON},
		{FileName: ".stackie.user.yml", Type: YAML},
		{FileName: ".stackie.user.json", Type: JSON},
		{FileName: "stackie.local.yml", Type: YAML},
		{FileName: "stackie.local.json", Type: JSON},
		{FileName: "stackie.user.yml", Type: YAML},
		{FileName: "stackie.user.json", Type: JSON},
	}
}
