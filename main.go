// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package main

import "github.com/mauriciozanettisalomao/backseat-driver-kid/cmd"

var (
	// BuildStamp is a timestamp (injected by go) of the build time
	BuildStamp = "None"
	// GitHash is the tag for current hash the build represents
	GitHash = "None"
)

func main() {
	cmd.Execute()
}
