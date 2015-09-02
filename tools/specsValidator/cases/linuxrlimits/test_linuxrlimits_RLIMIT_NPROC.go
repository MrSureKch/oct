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
//

package linuxrlimits

import (
	"github.com/huawei-openlab/oct/tools/specsValidator/manager"
	"github.com/opencontainers/specs"
	"strconv"
)

func TestRlimitNPROCSoft() string {
	var rlimitstest specs.Rlimit = specs.Rlimit{
		Type: 0,
		Soft: uint64(2048),
		Hard: uint64(0),
	}
	var testResult manager.TestResult
	linuxspec := setRlimits(rlimitstest)
	result, err := testRlimits(&linuxspec, "-n", strconv.FormatUint(rlimitstest.Soft, 10), true)
	testResult.Set("TestRlimitNPROCSoft", rlimitstest.Soft, err, result)
	return testResult.Marshal()
}

func TestRlimitNPROCHard() string {
	var rlimitstest specs.Rlimit = specs.Rlimit{
		Type: 0,
		Soft: uint64(0),
		Hard: uint64(20480),
	}
	var testResult manager.TestResult
	linuxspec := setRlimits(rlimitstest)
	result, err := testRlimits(&linuxspec, "-n", strconv.FormatUint(rlimitstest.Hard, 10), false)
	testResult.Set("TestRlimitNPROCHard", rlimitstest.Hard, err, result)
	return testResult.Marshal()
}