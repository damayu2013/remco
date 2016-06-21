// Copyright © 2016 The Remco Authors
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

package poll

import "github.com/spf13/cobra"

// pollCmd represents the watch command
var Cmd = &cobra.Command{
	Use:   "poll",
	Short: "poll a backend for changes and render the template accordingly",
}

func init() {
	Cmd.PersistentFlags().Bool("onetime", false, "run once and exit")
}
