/*
 * Copyright (c) 2020 wellwell.work, LLC by Zoe
 *
 * Licensed under the Apache License 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"go.zoe.im/x/cli"
	"go.zoe.im/x/version"
)

var (
	// global cmd contains all sub command
	cmd = cli.New(
		cli.Name("mart"),
		cli.Short("🍦 A universal online marketplace that can deliver anything."),
		version.NewOption(true),
		cli.Run(func(c *cli.Command, args ...string) {
			c.Help()
		}),
	)
)

// Register create a sub command
func Register(cmds ...*cli.Command) error {
	return cmd.Register(cmds...)
}

// Run execute command
func Run(opts ...cli.Option) error {
	return cmd.Run(opts...)
}
