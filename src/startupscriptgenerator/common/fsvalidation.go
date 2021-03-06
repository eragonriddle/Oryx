// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.
// --------------------------------------------------------------------------------------------

package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

// Gets the full path from a relative path, and ensure the path exists.
func GetValidatedFullPath(filePath string) string {
	fullAppPath, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(fullAppPath); os.IsNotExist(err) {
		panic("Path '" + fullAppPath + "' does not exist.")
	}
	return fullAppPath
}

// Writes the entrypoint command to an executable file
func WriteScript(filePath string, command string) {
	fmt.Println("Writing output script to '" + filePath + "'")
	ioutil.WriteFile(filePath, []byte(command), 0755)
}

// Add given permission to a file
func AddPermission(filePath string, permission int) bool {
	err := os.Chmod(filePath, os.FileMode(permission))
	if err != nil {
		return false
	}
	return true
}

// Check if the command is a file in app's repository and add execution permission to it
func ParseCommandAndAddExecutionPermission(commandString string, sourcePath string) bool {
	absoluteFilePath, err := filepath.Abs(filepath.Join(sourcePath, commandString))
	if err != nil {
		panic(err)
	} else {
		if FileExists(absoluteFilePath) {
			return AddPermission(absoluteFilePath, 0755)
		}
		return false
	}
}
