/*
Copyright 2020 The Kubernetes Authors.

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

package build

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewCIInstance can be used to create a new build `Instance` for use in CI.
func NewCIInstance(opts *Options) *Instance {
	instance := NewInstance(opts)
	instance.opts.CI = true

	return instance
}

// Push pushes the build by taking the internal options into account.
func (bi *Instance) Build() error {
	/*
		def main(args):
				# pylint: disable=too-many-branches
				"""Build and push kubernetes.

				This is a python port of the kubernetes/hack/jenkins/build.sh script.
				"""
				if os.path.split(os.getcwd())[-1] != 'kubernetes':
						print >>sys.stderr, (
								'Scenario should only run from either kubernetes directory!')
						sys.exit(1)
	*/

	buildExists, buildExistsErr := bi.checkBuildExists()
	if buildExistsErr != nil {
		return errors.Wrapf(buildExistsErr, "checking if build exists")
	}

	if buildExists {
		logrus.Infof("Build already exists. Exiting...")
		return nil
	}

	/*
		env = {
				# Skip gcloud update checking; do we still need this?
				'CLOUDSDK_COMPONENT_MANAGER_DISABLE_UPDATE_CHECK': 'true',
				# Don't run any unit/integration tests when building
				'KUBE_RELEASE_RUN_TESTS': 'n',
		}
	*/

	/*
		if args.register_gcloud_helper:
				# Configure docker client for gcr.io authentication to allow communication
				# with non-public registries.
				check_no_stdout('gcloud', 'auth', 'configure-docker')
	*/

	/*
		for key, value in env.items():
				os.environ[key] = value
	*/

	/*
		check('make', 'clean')
	*/

	// Create a Kubernetes build
	/*
		if args.fast:
				check('make', 'quick-release')
		else:
				check('make', 'release')
	*/

	// Pushing the build
	return bi.Push()
}

// checkBuildExists check if the target build exists in the specified GCS
// bucket. This allows us to speed up build jobs by not duplicating builds.
func (bi *Instance) checkBuildExists() (bool, error) {
	/*
		def check_build_exists(gcs, suffix, fast):
			""" check if a k8s build with same version
					already exists in remote path
			"""
			if not os.path.exists('hack/print-workspace-status.sh'):
					print >>sys.stderr, 'hack/print-workspace-status.sh not found, continue'
					return False
	*/

	/*
		version = ''
		try:
				match = re.search(
						r'gitVersion ([^\n]+)',
						check_output('hack/print-workspace-status.sh')
				)
				if match:
						version = match.group(1)
		except subprocess.CalledProcessError as exc:
				# fallback with doing a real build
				print >>sys.stderr, 'Failed to get k8s version, continue: %s' % exc
				return False

		if version:
				if not gcs:
						gcs = 'kubernetes-release-dev'
				gcs = 'gs://' + gcs
				mode = 'ci'
				if fast:
						mode += '/fast'
				if suffix:
						mode += suffix
				gcs = os.path.join(gcs, mode, version)
				try:
						check_no_stdout('gsutil', 'ls', gcs)
						check_no_stdout('gsutil', 'ls', gcs + "/kubernetes.tar.gz")
						check_no_stdout('gsutil', 'ls', gcs + "/bin")
						return True
				except subprocess.CalledProcessError as exc:
						print >>sys.stderr, (
								'gcs path %s (or some files under it) does not exist yet, continue' % gcs)
		return False
	*/

	return true, nil
}
