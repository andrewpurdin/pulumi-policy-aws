// Copyright 2016-2019, Pulumi Corporation.
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

package integrationtests

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	runPolicyPackIntegrationTest(
		t, "network",
		awsGuardSettings{},
		map[string]string{
			"aws:region": "us-west-2",
		},
		[]policyTestScenario{
			// Test scenario 1 - ALB Listener is using HTTP and not redirecting to HTTPS.
			{
				[]string{
					"aws:elasticloadbalancingv2:Listener (httpListener):",
					"Default action for HTTP listener must be a redirect using HTTPS.",
				},
			},
			// Test scenario 2 - AOK.
			{
				WantErrors: nil,
			},
		})
}
