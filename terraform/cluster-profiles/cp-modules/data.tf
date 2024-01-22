# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

data "spectrocloud_registry" "public_registry" {
  name = var.registry_name
}

data "spectrocloud_pack" "generic" {
  count        = length(keys(var.packs))
  name         = keys(var.packs)[count.index]
  version      = var.packs[keys(var.packs)[count.index]]
  registry_uid = data.spectrocloud_registry.public_registry.id
}