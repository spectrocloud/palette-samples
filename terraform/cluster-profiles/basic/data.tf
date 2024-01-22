# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

data "spectrocloud_registry" "public_registry" {
  name = "Public Repo"
}


data "spectrocloud_pack" "csi-aws-ebs" {
  name         = "csi-aws-ebs"
  version      = "1.22.0"
  registry_uid = data.spectrocloud_registry.public_registry.id
}

data "spectrocloud_pack" "csi-aws-ebs-1-24" {
  name         = "csi-aws-ebs"
  version      = "1.24.0"
  registry_uid = data.spectrocloud_registry.public_registry.id
}

data "spectrocloud_pack" "cni-calico" {
  name         = "cni-calico"
  version      = "3.26.1"
  registry_uid = data.spectrocloud_registry.public_registry.id
}

data "spectrocloud_pack" "kubernetes" {
  name         = "kubernetes"
  version      = "1.27.5"
  registry_uid = data.spectrocloud_registry.public_registry.id
}

data "spectrocloud_pack" "ubuntu-aws" {
  name         = "ubuntu-aws"
  version      = "22.04"
  registry_uid = data.spectrocloud_registry.public_registry.id
}