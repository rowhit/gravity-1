/*
Copyright 2020 Gravitational, Inc.

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

package utils

import (
	"log/syslog"

	"github.com/gravitational/trace"
)

// SyslogWrite writes the message to the system log with the specified priority
// and tag.
func SyslogWrite(priority syslog.Priority, message, tag string) error {
	w, err := syslog.New(priority, tag)
	if err != nil {
		return trace.Wrap(err)
	}
	defer w.Close()
	if _, err := w.Write([]byte(message)); err != nil {
		return trace.Wrap(err)
	}
	return nil
}
