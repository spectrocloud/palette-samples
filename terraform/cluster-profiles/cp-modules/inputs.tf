# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0



variable "name" {
  type        = string
  description = "The name of the cluster profile. The version is appended."
}

variable "description" {
  type        = string
  description = "The description of the cluster profile."
  default     = ""
}

variable "infrastructure_provider" {
  type        = string
  description = "The infrastructure provider the cluster profile is for."
}

variable "cluster_profile_type" {
  type        = string
  description = "The type of cluster profile. Default value is 'cluster'."
  default     = "cluster"
}

variable "profile_version" {
  type        = string
  description = "The version for the profile"
  default     = "1.0.0"
}

variable "context" {
  type        = string
  description = "The Palette scope to create the cluster profile in."
  default     = "project"
}


variable "registry_name" {
  type    = string
  default = "Public Repo"
}

variable "pack_order" {
  type        = list(string)
  description = "The Ordered list of pack names. The order must match with the expected layer of a cluster profile. The order goes from highest to lowest. For example. the first item has the highest priority order value assigned and so on."
}

variable "packs" {
  type        = map(string)
  description = "A list "
  default     = {}
}


variable "custom_yaml_files" {
  type        = map(string)
  description = "The file path to the custom YAML file that matches the pack name its values are overriding. Use the function 'file() to read the file. Example:  'kubernetes: 'file(path/to/file)'"
  default     = {}
}

variable "tags" {
  type        = list(string)
  description = "The default tags to apply to Palette resources"
  default     = []
}


locals {
  combined_packs = [for pack_name in var.pack_order : {
    name      = pack_name
    version   = var.packs[pack_name]
    pack_data = [for pack in data.spectrocloud_pack.generic : pack if pack.name == pack_name][0]
  }]
}
