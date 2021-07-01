package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"plugin"
	"strings"
)

type CAHPlugin interface {
	Init()
}

type PluginManager struct {
}

func (p *PluginManager) LoadPlugin(mod string) {

	plug, err := plugin.Open(mod)
	if err != nil {
		panic("Error loading " + mod)
		return
	}
	symPlugin, err := plug.Lookup("CAHPlugin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var plugin CAHPlugin
	plugin, ok := symPlugin.(CAHPlugin)
	if !ok {
		panic("Error parsing symbols of " + mod)
		return
	}
	plugin.Init()

}

func (p *PluginManager) LoadAllPlugins(basePath string) {

	wd, err := os.Getwd()

	lookUp := path.Join(wd, basePath, "plugins")

	files, err := ioutil.ReadDir(lookUp)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".so") {
			p.LoadPlugin(path.Join(lookUp, f.Name()))
		}
	}
}
