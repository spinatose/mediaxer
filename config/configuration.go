package config

type Config struct {
	MoveSourceFiles 	bool `json:"moveSourceFiles"`
	TargetFolder 		string `json:"targetFolder"`
	SourceFolder		string `json:"sourceFolder"`
	ResultFolderPattern string `json:"resultFolderPattern"`
}