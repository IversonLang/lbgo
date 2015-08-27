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
	"html/template"
	"io"
	"log"
	"os"
	"strings"
    "os/signal"
    "syscall"
)

const (
    REEXEC_FLAG = "LBGO_SYSTEMD_REEXEC"
)

var commands = []*Command{
	cmdVersion,
    cmdRun,
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

    signals := make(chan os.Signal)
    signal.Notify(signals, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

    if os.Getenv(REEXEC_FLAG) != "" {

    }

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		//        help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			os.Exit(cmd.Run(cmd, args))
			return
		}
	}

	fmt.Fprintf(os.Stderr, "lbgo: unknown subcommand %q\nRun 'lbgo help' for usage.\n", args[0])
	os.Exit(2)
}

var usageTemplate = `Lbgo is a tool for helping load balance of services.

Usage:

	lbgo command [arguments]

The commands are:
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "lbgo help [command]" for more information about a command.

Additional help topics:
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "lbgo help [topic]" for more information about that topic.

`

func usage() {
	tmpl(os.Stdout, usageTemplate, commands)
	os.Exit(2)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": func(s template.HTML) template.HTML {
		return template.HTML(strings.TrimSpace(string(s)))
	}})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
