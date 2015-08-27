// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var cmdRun = &Command{
	UsageLine: "run",
	Short:     "show the lbgo and Go version",
	Long: `
show the lbgo and Go version

bee version
    lbgo  :0.0.1
    Go    :go version go1.5 darwin/amd64

`,
}

var (
	pidFile string
)

func init() {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	fs.StringVar(&pidFile, "p", "/run/lbgo.pid", "pid file, default is '/run/lbgo.pid'")
	cmdRun.Flag = *fs
	cmdRun.Run = runCmd
}

func runCmd(cmd *Command, args []string) int {
	fmt.Println("lbgo   :" + version)
	goversion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Go    :" + string(goversion))
	return 0
}
