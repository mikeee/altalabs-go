//go:build e2e

/*
Copyright 2025 Mike Nguyen (mikeee) <hey@mike.ee>
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

package e2e

import (
	"os/exec"
	"testing"
)

func setupDocker(t *testing.T) {
	setupCmd := exec.Command("docker-compose", "-f", "../../../test/infra/docker-compose-local-amd64.yaml", "up", "-d")
	if err := setupCmd.Run(); err != nil {
		t.Fatalf("Failed to start docker compose: %v", err)
	}
	t.Cleanup(func() {
		teardownCmd := exec.Command("docker-compose", "-f", "../../../test/infra/docker-compose-local-amd64.yaml", "down")
		teardownCmd.Run()
	})
}

func TestLocalE2E(t *testing.T) {
	setupDocker(t)

	// TODO: Steps required:
	// 1. Wait for services to be healthy
	// 2. Activate local hosted Altalabs instance
	// 3. Check status
	t.Skip("Skipping local E2E tests (not implemented yet)")
}
