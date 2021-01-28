package main

import (
	"flag"
	"fmt"
)

var subtitleName string
var showName string

var auto bool
var resizeFoundSubs bool
var subSize int

func initFlags() {

	const (
		usageSubs = "Use this option to define the common subtitle name to rename"
	)

	flag.StringVar(&subtitleName, "subname", "", usageSubs)
	flag.StringVar(&subtitleName, "s", "", usageSubs+" (shorthand)")
	flag.StringVar(&showName, "mkvname", "", "Specifies the name of the show to use, choose a name that most of the file names contain!")

	flag.BoolVar(&auto, "auto", false, "Specifies whether the script should automatically deduce show and subtitle name from file formats, asks before renaming!")
	flag.BoolVar(&resizeFoundSubs, "resize", false, "Use this flag to resize all found subtitles if they are in a recognized format")
	flag.IntVar(&subSize, "subsize", 40, "Specifies the size to which the subitltes should be resized to, default value is [40]")
}

func main() {
	// create flags and parse them
	initFlags()
	flag.Parse()

	fmt.Println("Given value for subtitle name is", subtitleName)
}
