package base

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const CmdName = "leetcode"

var Commands []*Command

type Command struct {
	Run       func(cmd *Command, args []string)
	UsageLine string
	Short     string
	Long      string
}

func (c *Command) Name() string {
	name := c.UsageLine
	if i := strings.Index(name, " "); i > 0 {
		name = name[0:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Printf("usage: %s %s\n", CmdName, c.UsageLine)
	fmt.Printf("Run '%s help %s' for details.\n", CmdName, c.Name())
	Exit()
}

func (c *Command) UsageHelp() {
	fmt.Printf("usage: %s %s\n\n", CmdName, c.UsageLine)
	fmt.Println(c.Long)
	Exit()
}

func Usage() {
	fmt.Printf("%s is a tool for managing leetcode source code.\n\n", CmdName)
	fmt.Println("Usage:")
	fmt.Printf("\t%s <command> [arguments]\n", CmdName)
	fmt.Println("The commands are:")
	for _, cmd := range Commands {
		fmt.Printf("\t%-11s \t%s\n", cmd.Name(), cmd.Short)
	}
	fmt.Printf("\nUse \"%s help <command>\" for more information about a command.", CmdName)
	fmt.Println()
	Exit()
}

func FilePutContents(filename string, data []byte) []byte {
	if len(data) > 0 {
		filename = getFilePath(filename)
		err := ioutil.WriteFile(filename, data, 0644)
		CheckErr(err)
	}
	return data
}

func getFilePath(filename string) string {
	if dir := path.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			CheckErr(err)
		}
	}
	return filename
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Exit() {
	os.Exit(0)
}

func AuthInfo(cmd string) string {
	format := "<!--|This file generated by command(leetcode %s); DO NOT EDIT.%s|-->\n"
	format += "<!--+----------------------------------------------------------------------+-->\n"
	format += "<!--|@author    Openset <openset.wang@gmail.com>                           |-->\n"
	format += "<!--|@link      https://github.com/openset                                 |-->\n"
	format += "<!--|@home      https://github.com/openset/leetcode                        |-->\n"
	format += "<!--+----------------------------------------------------------------------+-->\n"
	return fmt.Sprintf(format, cmd, strings.Repeat(" ", 15-len(cmd)))
}
