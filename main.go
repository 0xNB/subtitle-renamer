package main

import (
	"flag"
	"fmt"
)

/*
* Name of the subtile file that we want to find
 */
var subtitleName string
var vidName string
var auto bool
var resizeFoundSubs bool
var subSize int

func initFlags() {

	const (
		usageSubs    = "Use this option to define the common subtitle name to rename"
		usageVidName = "Specifies the name of the show to use, choose a name that most of the file names contain!"
		usageAuto    = "Specifies whether the script should automatically deduce show and subtitle name from file formats, asks before renaming! Default is [false]"
		usageResize  = "Use this flag to resize all found subtitles if they are in a recognized format"
		usageSubSize = "Specifies the size to which the subitltes should be resized to, default value is [40]"
	)

	flag.StringVar(&subtitleName, "subname", "", usageSubs)
	flag.StringVar(&subtitleName, "s", "", usageSubs+" (shorthand)")
	flag.StringVar(&vidName, "vidname", "", usageVidName)
	flag.StringVar(&vidName, "v", "", usageVidName+" (shorthand)")
	flag.BoolVar(&auto, "auto", false, usageAuto)
	flag.BoolVar(&auto, "a", false, usageAuto+" (shorthand)")
	flag.BoolVar(&resizeFoundSubs, "resize", false, usageResize)
	flag.BoolVar(&resizeFoundSubs, "r", false, usageResize+" (shorthand)")
	flag.IntVar(&subSize, "subsize", 40, usageSubSize)
	flag.IntVar(&subSize, "z", 40, usageSubSize+" (shorthand)")

}

func main() {
	// create flags and parse them
	initFlags()
	flag.Parse()

	// print status in console
	fmt.Println("Given value for subtitle name is", subtitleName)
}
