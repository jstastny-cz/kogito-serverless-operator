// Copyright 2023 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"os"
	"path"
	"strings"
)

func RemoveFileExtension(fileName string) string {
	if i := strings.LastIndex(fileName, "."); i >= 0 {
		return fileName[:i]
	}
	return fileName
}

func RemoveKnownExtension(fileName, extension string) string {
	if i := strings.LastIndex(fileName, extension); i >= 0 {
		return fileName[:i]
	}
	return fileName
}

// PathToString replaces the PathSeparator from a given path.
func PathToString(pathRef string) string {
	return strings.ReplaceAll(path.Clean(pathRef), string(os.PathSeparator), "")
}
