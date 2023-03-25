package config

type Config struct {
	MoveSourceFiles 	bool `json:"moveSourceFiles"`
	DestinationFolder 	string `json:"destinationFolder"`
	SourceFolder		string `json:"sourceFolder"`
	ResultFolderPattern string `json:"resultFolderPattern"`
}