# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

resource "spectrocloud_cluster_profile" "profile" {

  name        = "${var.name}-${replace(var.profile_version, ".", "-")}"
  description = var.description
  tags        = concat(var.tags, ["version:${var.profile_version}"])
  cloud       = var.infrastructure_provider
  type        = var.cluster_profile_type
  version     = var.profile_version

  dynamic "pack" {
    for_each = { for idx, cp in local.combined_packs : idx => cp }

    content {
      name   = pack.value.name
      tag    = pack.value.pack_data.version
      uid    = pack.value.pack_data.id
      values = lookup(var.custom_yaml_files, pack.value.name, null) != null ? var.custom_yaml_files[pack.value.name] : pack.value.pack_data.values
    }
  }


  depends_on = [
    data.spectrocloud_pack.generic
  ]
}
