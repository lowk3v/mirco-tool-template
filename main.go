package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/lowk3v/micro-tool-template/internal"
	"github.com/lowk3v/micro-tool-template/utils"
	"os"
)
import global "github.com/lowk3v/micro-tool-template/config"

func __existArg(arg string) bool {
	args := os.Args[1:]
	return len(args) > 0 && args[0] == arg
}

type ArgList []string

func (a *ArgList) String() string {
	return ""
}

func (a *ArgList) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func _banner() {
	// https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=%20dumpsc
	_, _ = fmt.Fprintf(os.Stderr, "%s by %s\n%s\nCredits: https://github.com/%s/%s\nTwitter: https://twitter.com/%s\n\n",
		color.HiBlueString(`
     █████╗ ██████╗ ██████╗ 
    ██╔══██╗██╔══██╗██╔══██╗
    ███████║██████╔╝██████╔╝
    ██╔══██║██╔═══╝ ██╔═══╝ 
    ██║  ██║██║     ██║     
    ╚═╝  ╚═╝╚═╝     ╚═╝ `),
		color.BlueString("@author_name"),
		"project_description",
		"author_name",
		"project_name",
		"#",
	)
	_, _ = fmt.Fprintf(os.Stderr, "Usage of: %s <options> <args>\n", os.Args[0])
	flag.PrintDefaults()
}

func _parseFlags() (string, *internal.Options, error) {
	var configPath string
	var verbose bool
	var version bool
	var output string

	// global arguments
	flag.StringVar(&configPath, "c", "./config.yml", "optional. Path to config file")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.StringVar(&output, "o", "./output", "output directory")
	flag.BoolVar(&version, "-version", false, "print version and exit")
	flag.Usage = _banner
	flag.Parse()

	// global configurations
	options := &internal.Options{
		Verbose: verbose,
		Version: version,
		Output:  output,
	}

	// module configurations, implement if needed

	if err := utils.FileExists(configPath, false); err != nil {
		return "", nil, err
	}

	return configPath, options, nil
}

func main() {
	cfgPath, options, err := _parseFlags()
	if utils.HandleError(err, "") {
		os.Exit(0)
	}
	if global.NewConfig(cfgPath) != nil {
		os.Exit(0)
	}
	app := internal.New(options)
	app.Run()
}
