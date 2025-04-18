# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG GO_VERSION_TOOLING
ARG OS_CODENAME

# The Release Tooling can be compiled with a different version of Go than the
# the Go version used in the kubernetes/kubernetes active development branch.
FROM golang:${GO_VERSION_TOOLING}-${OS_CODENAME} AS builder

WORKDIR /go/src/k8s.io/release

COPY ./ ./

RUN ./compile-release-tools

### Production image

# Includes tools used for building Kubernetes in CI
FROM debian:${OS_CODENAME}

# arg that specifies the image name (for debugging)
ARG IMAGE_ARG

# The Golang version for the builder image should always be explicitly set to
# the Golang version of the kubernetes/kubernetes active development branch
ARG GO_VERSION

# arg that specifies docker-buildx version to install
ARG DOCKER_BUILDX_VERSION=0.10.2

# add envs:
# - so we can debug with the image name:tag
# - adding gsutil etc. to path (where we will install them)
# - disabling prompts when installing gsutil etc.
# - hinting that we are in a docker container
ENV IMAGE=${IMAGE_ARG} \
    GOPATH=/home/prow/go \
    PATH=/home/prow/go/bin:/usr/local/go/bin:/google-cloud-sdk/bin:${PATH} \
    CLOUDSDK_CORE_DISABLE_PROMPTS=1 \
    CONTAINER=docker

# copy in image utility scripts
COPY images/releng/k8s-ci-builder/wrapper.sh /usr/local/bin/

# Install tools needed to:
# - install docker
# - build kubernetes (dockerized)
#
# TODO: the `sed` is a bit of a hack, look into alternatives.
# Why this exists: `docker service start` on debian runs a `cgroupfs_mount` method,
# We're already inside docker though so we can be sure these are already mounted.
# Trying to remount these makes for a very noisy error block in the beginning of
# the pod logs, so we just comment out the call to it... :shrug:
RUN echo "Installing Packages ..." \
        && apt-get update \
        && apt-get install -y --no-install-recommends \
            apt-transport-https \
            build-essential \
            ca-certificates \
            curl \
            file \
            git \
            gnupg2 \
            jq \
            kmod \
            lsb-release \
            mercurial \
            net-tools \
            pkg-config \
            procps \
            python3 \
            python3-dev \
            python3-pip \
            rsync \
            software-properties-common \
            unzip \
        && rm -rf /var/lib/apt/lists/* \
    && echo "Installing Go ..." \
        && export GO_TARBALL="go${GO_VERSION}.linux-amd64.tar.gz"\
        && curl -fsSL "https://storage.googleapis.com/golang/${GO_TARBALL}" --output "${GO_TARBALL}" \
        && tar xzf "${GO_TARBALL}" -C /usr/local \
        && rm "${GO_TARBALL}"\
        && mkdir -p "${GOPATH}/bin" \
    && echo "Installing gcloud SDK, kubectl ..." \
        && curl -fsSL https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.tar.gz --output google-cloud-sdk.tar.gz \
        && tar xzf google-cloud-sdk.tar.gz -C / \
        && rm google-cloud-sdk.tar.gz \
        && /google-cloud-sdk/install.sh \
            --disable-installation-options \
            --bash-completion=false \
            --path-update=false \
            --usage-reporting=false \
        && gcloud components install kubectl \
    && echo "Installing Docker ..." \
        && install -m 0755 -d /etc/apt/keyrings && curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg \
        && echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null \
        && apt-get update \
        && apt-get install -y --no-install-recommends docker-ce=5:24.0.* \
        && rm -rf /var/lib/apt/lists/* \
        && sed -i 's/cgroupfs_mount$/#cgroupfs_mount\n/' /etc/init.d/docker \
    && echo "Installing Docker buildx ..." \
        && mkdir -p /usr/libexec/docker/cli-plugins \
        && curl -fsSLo /usr/libexec/docker/cli-plugins/docker-buildx https://github.com/docker/buildx/releases/download/v${DOCKER_BUILDX_VERSION}/buildx-v${DOCKER_BUILDX_VERSION}.linux-amd64 \
        && chmod a+x /usr/libexec/docker/cli-plugins/docker-buildx \
    && echo "Ensuring Legacy Iptables ..." \
        && update-alternatives --set iptables /usr/sbin/iptables-legacy \
        && update-alternatives --set ip6tables /usr/sbin/ip6tables-legacy

# Copy in release tools from kubernetes/release
WORKDIR /
COPY --from=builder /go/bin/* ./

# entrypoint is our wrapper script, in Prow you will need to explicitly re-specify this
ENTRYPOINT ["wrapper.sh"]
# volume for docker in docker, use an emptyDir in Prow
VOLUME ["/var/lib/docker"]
