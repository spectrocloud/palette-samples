# Copyright (c) Spectro Cloud
# SPDX-License-Identifier: Apache-2.0

variable "tags" {
  type        = list(string)
  description = "The default tags to apply to Palette resources"
  default = [
    "spectro-cloud-education",
    "repository:spectrocloud:tutorials",
    "terraform_managed:true",
  ]
}
