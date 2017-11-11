package main

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v1"
)

type CloudShellConfig struct {
	Interval int `yaml:"interval"`
	Padding  struct {
		Top  int `yaml:"top"`
		Left int `yaml:"left"`
	} `yaml:"padding"`
	Outputs struct {
		Host bool `yaml:"host,omitempty"`
		Load bool `yaml:"load,omitempty"`
		Cpu  bool `yaml:"cpu,omitempty"`
		Ram  bool `yaml:"ram,omitempty"`
		Swap bool `yaml:"swap,omitempty"`
		Net  []struct {
			Name    string `yaml:"name"`
			Enabled bool   `yaml:"enabled"`
		} `yaml:"net,omitempty"`
		Disk []struct {
			Name    string `yaml:"name"`
			Enabled bool   `yaml:"enabled"`
			Mount   string `yaml:"mount"`
			Space   bool   `yaml:"space"`
		} `yaml:"disk,omitempty"`
	} `yaml:"outputs"`
	outputs []string
}

func ReadConfig(p string) *CloudShellConfig {
	// declare our return value
	retv := &CloudShellConfig{}

	// open the config file
	fp, er := os.Open(p)
	pcheck(er)
	defer fp.Close()

	// read the file
	data := bytes.NewBuffer(nil)
	io.Copy(data, fp)

	// and parse the data
	pcheck(yaml.Unmarshal([]byte(data.Bytes()), retv))

	// build our order list for outputters
	if retv.Outputs.Host {
		retv.outputs = append(retv.outputs, "host")
	}
	if retv.Outputs.Load {
		retv.outputs = append(retv.outputs, "load")
	}
	if retv.Outputs.Cpu {
		retv.outputs = append(retv.outputs, "cpu")
	}
	if retv.Outputs.Ram {
		retv.outputs = append(retv.outputs, "ram")
	}
	if retv.Outputs.Swap {
		retv.outputs = append(retv.outputs, "swap")
	}

	// and hand it back
	return retv
}
