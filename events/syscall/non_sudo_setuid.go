//go:build linux
// +build linux

// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.
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

package syscall

import (
	"github.com/falcosecurity/event-generator/events"
)

var _ = events.Register(NonSudoSetuid)

func NonSudoSetuid(h events.Helper) error {
	if h.Spawned() {
		h.Log().Debug("first setuid to something non-root, then try to setuid back to root")
		if err := becameUser(h, "daemon"); err != nil {
			return err
		}
		err := becameUser(h, "root")
		h.Log().WithError(err).Debug("ignore root setuid error")
		return nil
	} else {
		return h.SpawnAs("child", "syscall.NonSudoSetuid")
	}
}
