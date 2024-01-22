# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

resource "spectrocloud_cluster_profile" "aws-profile-1-0-0" {

  name        = "tf-aws-profile"
  description = "A basic cluster profile for AWS"
  tags        = concat(var.tags, ["env:aws", "version:1.0.0"])
  cloud       = "aws"
  type        = "cluster"
  version     = "1.0.0"

  pack {
    name   = data.spectrocloud_pack.ubuntu-aws.name
    tag    = data.spectrocloud_pack.ubuntu-aws.version
    uid    = data.spectrocloud_pack.ubuntu-aws.id
    values = data.spectrocloud_pack.ubuntu-aws.values
  }

  pack {
    name   = data.spectrocloud_pack.kubernetes.name
    tag    = data.spectrocloud_pack.kubernetes.version
    uid    = data.spectrocloud_pack.kubernetes.id
    values = data.spectrocloud_pack.kubernetes.values
  }

  pack {
    name   = data.spectrocloud_pack.cni-calico.name
    tag    = data.spectrocloud_pack.cni-calico.version
    uid    = data.spectrocloud_pack.cni-calico.id
    values = data.spectrocloud_pack.cni-calico.values
  }

  pack {
    name   = data.spectrocloud_pack.csi-aws-ebs.name
    tag    = data.spectrocloud_pack.csi-aws-ebs.version
    uid    = data.spectrocloud_pack.csi-aws-ebs.id
    values = data.spectrocloud_pack.csi-aws-ebs.values
  }
}

resource "spectrocloud_cluster_profile" "aws-profile-1-0-1" {

  name        = "tf-aws-profile"
  description = "A basic cluster profile for AWS"
  tags        = concat(var.tags, ["env:aws", "version:1.0.1"])
  cloud       = "aws"
  type        = "cluster"
  version     = "1.0.1"

  pack {
    name   = data.spectrocloud_pack.ubuntu-aws.name
    tag    = data.spectrocloud_pack.ubuntu-aws.version
    uid    = data.spectrocloud_pack.ubuntu-aws.id
    values = data.spectrocloud_pack.ubuntu-aws.values
  }

  pack {
    name   = data.spectrocloud_pack.kubernetes.name
    tag    = data.spectrocloud_pack.kubernetes.version
    uid    = data.spectrocloud_pack.kubernetes.id
    values = data.spectrocloud_pack.kubernetes.values
  }

  pack {
    name   = data.spectrocloud_pack.cni-calico.name
    tag    = data.spectrocloud_pack.cni-calico.version
    uid    = data.spectrocloud_pack.cni-calico.id
    values = data.spectrocloud_pack.cni-calico.values
  }

  pack {
    name   = data.spectrocloud_pack.csi-aws-ebs-1-24.name
    tag    = data.spectrocloud_pack.csi-aws-ebs-1-24.version
    uid    = data.spectrocloud_pack.csi-aws-ebs-1-24.id
    values = data.spectrocloud_pack.csi-aws-ebs-1-24.values
  }
}