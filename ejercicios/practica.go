package main

// file types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

const (
	exe    = ".exe"
	debian = ".deb"
	zip    = ".zip"
	gz     = ".gz"
	tar    = ".tar"
	rar    = ".rar"
	png    = ".png"
	jpg    = ".jpg"
	gif    = ".gif"
)

type file struct {
	name              string
	fileType          int
	isDir             bool
	isHidden          bool
	userName          string
	groupName         string
	size              int64
	modidifcationTime string
	mode              string
}

type styleFileType struct {
	icon   string
	color  string
	symbol string
}

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📄"},
	fileDirectory:  {icon: "📂", color: "BLUE", symbol: "/"},
	fileExecutable: {icon: "🚀", color: "GREEN", symbol: "*"},
	fileCompress:   {icon: "📦", color: "RED"},
	fileImage:      {icon: "📸", color: "YELLOW"},
	fileLink:       {icon: "🔗", color: "CYAN"},
}
