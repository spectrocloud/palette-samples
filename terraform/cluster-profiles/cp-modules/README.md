## Requirements

| Name                                                                              | Version   |
| --------------------------------------------------------------------------------- | --------- |
| <a name="requirement_terraform"></a> [terraform](#requirement_terraform)          | >= 1.5    |
| <a name="requirement_local"></a> [local](#requirement_local)                      | >= 2.4.0  |
| <a name="requirement_spectrocloud"></a> [spectrocloud](#requirement_spectrocloud) | >= 0.17.2 |

## Providers

| Name                                                                        | Version |
| --------------------------------------------------------------------------- | ------- |
| <a name="provider_spectrocloud"></a> [spectrocloud](#provider_spectrocloud) | 0.17.3  |

## Modules

No modules.

## Resources

| Name                                                                                                                                            | Type        |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [spectrocloud_cluster_profile.profile](https://registry.terraform.io/providers/spectrocloud/spectrocloud/latest/docs/resources/cluster_profile) | resource    |
| [spectrocloud_pack.generic](https://registry.terraform.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack)                    | data source |
| [spectrocloud_registry.public_registry](https://registry.terraform.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/registry)    | data source |

## Inputs

| Name                                                                                                   | Description                                                                                                                                                                                                                    | Type           | Default         | Required |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------- | --------------- | :------: |
| <a name="input_cluster_profile_type"></a> [cluster_profile_type](#input_cluster_profile_type)          | The type of cluster profile. Default value is 'cluster'.                                                                                                                                                                       | `string`       | `"cluster"`     |    no    |
| <a name="input_context"></a> [context](#input_context)                                                 | The Palette scope to create the cluster profile in.                                                                                                                                                                            | `string`       | `"project"`     |    no    |
| <a name="input_custom_yaml_files"></a> [custom_yaml_files](#input_custom_yaml_files)                   | The file path to the custom YAML file that matches the pack name its values are overriding. Use the function 'file() to read the file. Example: 'kubernetes: 'file(path/to/file)'                                              | `map(string)`  | `{}`            |    no    |
| <a name="input_description"></a> [description](#input_description)                                     | The description of the cluster profile.                                                                                                                                                                                        | `string`       | `""`            |    no    |
| <a name="input_infrastructure_provider"></a> [infrastructure_provider](#input_infrastructure_provider) | The infrastructure provider the cluster profile is for.                                                                                                                                                                        | `string`       | n/a             |   yes    |
| <a name="input_name"></a> [name](#input_name)                                                          | The name of the cluster profile. The version is appended.                                                                                                                                                                      | `string`       | n/a             |   yes    |
| <a name="input_pack_order"></a> [pack_order](#input_pack_order)                                        | The Ordered list of pack names. The order must match with the expected layer of a cluster profile. The order goes from highest to lowest. For example. the first item has the highest priority order value assigned and so on. | `list(string)` | n/a             |   yes    |
| <a name="input_packs"></a> [packs](#input_packs)                                                       | A list                                                                                                                                                                                                                         | `map(string)`  | `{}`            |    no    |
| <a name="input_profile_version"></a> [profile_version](#input_profile_version)                         | The version for the profile                                                                                                                                                                                                    | `string`       | `"1.0.0"`       |    no    |
| <a name="input_registry_name"></a> [registry_name](#input_registry_name)                               | n/a                                                                                                                                                                                                                            | `string`       | `"Public Repo"` |    no    |
| <a name="input_tags"></a> [tags](#input_tags)                                                          | The default tags to apply to Palette resources                                                                                                                                                                                 | `list(string)` | `[]`            |    no    |

## Outputs

| Name                                                  | Description |
| ----------------------------------------------------- | ----------- |
| <a name="output_export"></a> [export](#output_export) | n/a         |
| <a name="output_id"></a> [id](#output_id)             | n/a         |
