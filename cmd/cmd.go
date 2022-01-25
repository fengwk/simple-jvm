package cmd

import (
	"flag"
	"fmt"
	"os"
)

// Version 当前版本号
const Version = "0.0.1"

// CommandLine 命令行
type CommandLine struct {

	// 命令行选项
	Option struct {
		Xjre      string
		Classpath string
		Cp        string
		Version   bool
		Help      bool
	}

	// 主类的名称
	MainClassName string
	// 向主类传递的参数
	Args []string
}

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-option...] mainClassName [args...]\n\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	// 指定用法输出函数
	flag.Usage = printUsage
}

// Parse 解析命令行参数
func Parse() *CommandLine {
	var cmd CommandLine
	flag.StringVar(&cmd.Option.Xjre, "Xjre", "", "Specify the directory path where jre is located.")
	flag.StringVar(&cmd.Option.Classpath, "classpath", "", "Specify the loading path of the application class loader, and use the operating system-related path list separator to separate multiple class paths.")
	flag.StringVar(&cmd.Option.Cp, "cp", "", "short for classpath.")
	flag.BoolVar(&cmd.Option.Version, "version", false, "Display the version.")
	flag.BoolVar(&cmd.Option.Help, "help", false, "Displays the usage.")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.MainClassName = args[0]
		cmd.Args = args[1:]
	}
	return &cmd
}

func runJVM(cmd *CommandLine) {
	fmt.Println("SimpleJVM run")
	fmt.Println(cmd)
}

// Execute 根据命令行内容执行命令
func (cmd *CommandLine) Execute() {
	if cmd.Option.Version {
		fmt.Printf("SimpleJVM version %s\n", Version)
	} else if cmd.Option.Help {
		flag.CommandLine.SetOutput(os.Stdout)
		printUsage()
	} else if cmd.MainClassName == "" {
		flag.CommandLine.SetOutput(os.Stderr)
		printUsage()
		os.Exit(1)
	} else {
		runJVM(cmd)
	}
}
