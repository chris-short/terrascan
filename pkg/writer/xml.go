/*
    Copyright (C) 2020 Accurics, Inc.

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

package writer

import (
	"encoding/xml"
	"io"

	"github.com/accurics/terrascan/pkg/policy"
)

const (
	xmlFormat supportedFormat = "xml"
)

func init() {
	RegisterWriter(xmlFormat, XMLWriter)
}

// XMLWriter prints data in XML format
func XMLWriter(data policy.EngineOutput, writer io.Writer) error {
	j, _ := xml.MarshalIndent(data, "", "  ")
	writer.Write(j)
	writer.Write([]byte{'\n'})
	return nil
}