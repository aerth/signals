// aerth wuz here

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"flag"

	"gitlab.com/aerth/signals"
)

var verbose = flag.Bool("v", false, "")
var fatalerr = flag.Bool("x", false, "")

func main() {
	flag.Parse()
	log.SetFlags(0)
	if flag.NArg() < 2 {
		fmt.Printf("Usage: %s signal process[es]\n", os.Args[0])
		fmt.Printf("Example: %s kill %s\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
		os.Exit(111)
	}
	procs := flag.Args()[1:]
	sig := signals.Signal(flag.Arg(0))
	s := "none"
	for sigstr := range signals.Signals {
		if strings.Contains(sigstr, strings.ToUpper(flag.Arg(0))) {
			s = sigstr
		}
	}
	for _, proc := range procs {
		if *verbose {
			fmt.Printf("sending process %q the signal %q, it will be %q\n", proc, s, sig)
		}
		if err := filepath.Walk("/proc", searchAndDestroy(proc, sig)); err != nil {
			if err == io.EOF {
				err = nil
			} else {
				log.Fatalln(err)
			}
		}
	}
}

func searchAndDestroy(procname string, sig os.Signal) func(string, os.FileInfo, error) error {
	return func(path string, stat os.FileInfo, err error) error {
		if err != nil {
			if *verbose {
				log.Println(err)
			}
			if *fatalerr {
				os.Exit(111)
			}
		}

		if strings.Count(path, string(filepath.Separator)) != 3 || !strings.HasSuffix(path, string(filepath.Separator)+"status") {
			return nil
		}

		if split := strings.Split(path, string(filepath.Separator)); len(split) == 4 {
			pid, err := strconv.Atoi(split[2])
			if err != nil {
				if *verbose {
					log.Println(err)
				}
				if *fatalerr {
					os.Exit(111)
				}
			}
			f, err := ioutil.ReadFile(path)
			if err != nil {
				if *verbose {
					log.Println(err)
				}
				if *fatalerr {
					os.Exit(111)
				}
			}
			// TODO: leave it case insensitive?
			name := strings.ToLower(string(f[6:bytes.IndexByte(f, '\n')]))
			if *verbose {
				fmt.Println(name, pid)
			}
			name = strings.ToLower(name)

			// matched
			if strings.Contains(name, strings.ToLower(procname)) {
				proc, err := os.FindProcess(pid)
				if err != nil {
					log.Println(err)
					return nil
				}

				if sig.String() == "signal 0" {

					log.Println(name, pid)
				}
				// send sig
				if err := proc.Signal(sig); err != nil {
					log.Println(err)
				} else {
					log.Println(name, sig)
				}

			}
		}
		return nil
	}
}
