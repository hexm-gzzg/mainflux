// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
	prettyjson "github.com/hokaccha/go-prettyjson"
)

var (
	// Limit query parameter
	Limit uint = 10
	// Offset query parameter
	Offset uint = 0
	// Name query parameter
	Name string = ""
	// UserAuthToken user auth token parameter
	UserAuthToken string = ""
	// ConfigPath config path parameter
	ConfigPath string = ""
	// RawOutput raw output mode
	RawOutput bool = false
)

func logJSON(iList ...interface{}) {
	for _, i := range iList {
		m, err := json.Marshal(i)
		if err != nil {
			logError(err)
			return
		}

		pj, err := prettyjson.Format(m)
		if err != nil {
			logError(err)
			return
		}

		fmt.Printf("\n%s\n\n", string(pj))
	}
}

func logUsage(u string) {
	fmt.Printf(color.YellowString("\nusage: %s\n\n"), u)
}

func logError(err error) {
	boldRed := color.New(color.FgRed, color.Bold)
	boldRed.Print("\nerror: ")

	fmt.Printf("%s\n\n", color.RedString(err.Error()))
}

func logOK() {
	fmt.Printf("\n%s\n\n", color.BlueString("ok"))
}

func logCreated(e string) {
	if RawOutput {
		fmt.Println(e)
	} else {
		fmt.Printf(color.BlueString("\ncreated: %s\n\n"), e)
	}
}

func getUserAuthToken() string {
	if UserAuthToken == "" {
		log.Fatal("user auth token not valid, please pass using --user-auth-token flag or config file")
	}

	return UserAuthToken
}
