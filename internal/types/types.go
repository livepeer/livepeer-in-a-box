package types

type TagInformation struct {
	Name       string `json:"name"`
	Id         uint   `json:"id"`
	PreRelease bool   `json:"prerelease"`
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
}

type BuildFlags struct {
	Version string
}

type CliFlags struct {
	SkipDownloaded bool
	Cleanup        bool
	DownloadPath   string
	Platform       string
	Architecture   string
	ManifestFile   string
	Verbosity      string
}

type Service struct {
	Name         string `yaml:"name"`
	Project      string `yaml:"project"`
	Binary       string `yaml:"binary,omitempty"`
	Release      string `yaml:"release,omitempty"`
	ArchivePath  string `yaml:"archivePath,omitempty"`
	Skip         bool   `yaml:"skip"`
	SkipGPG      bool   `yaml:"skipGpg"`
	SkipChecksum bool   `yaml:"skipChecksum"`
}

type BoxManifest struct {
	Version string    `yaml:"version"`
	Release string    `yaml:"release,omitempty"`
	Box     []Service `yaml:"box,omitempty"`
}

type ArtifactInfo struct {
	Name         string
	Binary       string
	Version      string
	Platform     string
	Architecture string
	ArchiveURL   string
	ChecksumURL  string
	SignatureURL string
	Tag          TagInformation
}
