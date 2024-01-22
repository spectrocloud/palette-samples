# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

module "primary-cp-1-0-0" {
  source = "../cp-modules"

  name                    = "md-test"
  infrastructure_provider = "aws"
  cluster_profile_type    = "cluster"
  registry_name           = "Public Repo"
  profile_version         = "1.0.0"
  pack_order              = ["ubuntu-aws", "kubernetes", "cni-calico", "csi-aws-ebs"]
  packs = {
    "csi-aws-ebs" = "1.22.0"
    "cni-calico"  = "3.26.1"
    "kubernetes"  = "1.27.5"
    "ubuntu-aws"  = "22.04"
  }
}

module "primary-cp-1-0-1" {
  source = "../cp-modules"

  name                    = "md-test"
  infrastructure_provider = "aws"
  cluster_profile_type    = "cluster"
  registry_name           = "Public Repo"
  profile_version         = "1.0.1"
  pack_order              = ["ubuntu-aws", "kubernetes", "cni-calico", "csi-aws-ebs"]
  packs = {
    "csi-aws-ebs" = "1.22.0"
    "cni-calico"  = "3.26.1"
    "kubernetes"  = "1.27.5"
    "ubuntu-aws"  = "22.04"
  }
}