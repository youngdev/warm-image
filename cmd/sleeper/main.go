/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/golang/glog"
)

var (
	mode = flag.String("mode", "sleep", "One of: sleep or copy")
	to   = flag.String("to", "", "Where to copy this binary")
)

func main() {
	flag.Parse()

	switch *mode {
	case "sleep":
		// Sleep for 30 years.
		time.Sleep(30 * 365 * 24 * time.Hour)
		glog.Fatalf("Time to restart, goodbye cruel world.")
	case "copy":
		if *to == "" {
			glog.Fatalf("-to must be specified with -mode=copy")
		}
		from, err := os.Open(os.Args[0])
		if err != nil {
			glog.Fatal(err)
		}
		defer from.Close()

		to, err := os.OpenFile(*to, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			glog.Fatal(err)
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			glog.Fatal(err)
		}
	default:
		glog.Fatalf("Unsupported mode: %s", *mode)
	}
}
