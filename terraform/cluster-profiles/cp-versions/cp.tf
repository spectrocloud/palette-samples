# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

resource "spectrocloud_cluster_profile" "aws-profile" {
  count = length(local.cp-versions)

  name        = "tf-aws-profile-${local.cp-versions[count.index]}"
  description = "A basic cluster profile for AWS"
  tags        = concat(var.tags, ["env:aws", "version:${local.cp-versions[count.index]}"])
  cloud       = "aws"
  type        = "cluster"
  version     = local.cp-versions[count.index]


  dynamic "pack" {
    for_each = local.packs[local.cp-versions[count.index]]

    content {
      name = pack.key
      tag  = pack.value

      uid = local.pack_data[pack.key].data_source[
        index(
          [for v, p in local.packs : p[pack.key] if contains(keys(p), pack.key)],
          pack.value
        )
      ].id

      values = local.pack_data[pack.key].data_source[
        index(
          [for v, p in local.packs : p[pack.key] if contains(keys(p), pack.key)],
          pack.value
        )
      ].values
    }
  }


  depends_on = [
    data.spectrocloud_pack.cni-calico,
    data.spectrocloud_pack.csi-aws-ebs,
    data.spectrocloud_pack.kubernetes,
    data.spectrocloud_pack.ubuntu-aws,
    data.spectrocloud_pack.cni-cilium-oss
  ]
}
