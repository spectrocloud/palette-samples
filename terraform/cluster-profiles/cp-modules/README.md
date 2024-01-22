## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5 |
| <a name="requirement_local"></a> [local](#requirement\_local) | >= 2.4.0 |
| <a name="requirement_spectrocloud"></a> [spectrocloud](#requirement\_spectrocloud) | >= 0.17.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_spectrocloud"></a> [spectrocloud](#provider\_spectrocloud) | 0.17.3 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [spectrocloud_cluster_profile.profile](https://registry.terraform.io/providers/spectrocloud/spectrocloud/latest/docs/resources/cluster_profile) | resource |
| [spectrocloud_pack.generic](https://registry.terraform.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) | data source |
| [spectrocloud_registry.public_registry](https://registry.terraform.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/registry) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cluster_profile_type"></a> [cluster\_profile\_type](#input\_cluster\_profile\_type) | The type of cluster profile. Default value is 'cluster'. | `string` | `"cluster"` | no |
| <a name="input_context"></a> [context](#input\_context) | The Palette scope to create the cluster profile in. | `string` | `"project"` | no |
| <a name="input_description"></a> [description](#input\_description) | The description of the cluster profile. | `string` | `""` | no |
| <a name="input_infrastructure_provider"></a> [infrastructure\_provider](#input\_infrastructure\_provider) | The infrastructure provider the cluster profile is for. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | The name of the cluster profile. The version is appended. | `string` | n/a | yes |
| <a name="input_pack_order"></a> [pack\_order](#input\_pack\_order) | The Ordered list of pack names. The order must match with the expected layer of a cluster profile. | `list(string)` | n/a | yes |
| <a name="input_packs"></a> [packs](#input\_packs) | A list | `map(string)` | <pre>{<br>  "cni-calico": "3.26.1",<br>  "csi-aws-ebs": "1.22.0",<br>  "kubernetes": "1.27.5",<br>  "ubuntu-aws": "22.04"<br>}</pre> | no |
| <a name="input_profile_version"></a> [profile\_version](#input\_profile\_version) | The version for the profile | `string` | `"1.0.0"` | no |
| <a name="input_registry_name"></a> [registry\_name](#input\_registry\_name) | n/a | `string` | `"Public Repo"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | The default tags to apply to Palette resources | `list(string)` | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_export"></a> [export](#output\_export) | n/a |
| <a name="output_id"></a> [id](#output\_id) | n/a |
