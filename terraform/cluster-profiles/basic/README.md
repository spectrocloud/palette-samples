# Basic Pattern

The basic pattern of managing and maintain cluster profiles. Each cluster profile version is defined through the `spectrocloud_cluster_profile` resource. 


| Pros | 
| ---- |
| Easy to understand | 
| Easy to implement |
| Mimumum complexity |
| Profiles are clearly defined |
| Low chances of accidental changes |
| YAML customization supported through dedicated `pack {}` block |

| Cons |
| ---- |
| Requires code duplications |
| Tedious to maintain |


## Usage

1. Create a data resource for each pack. Replace the name of the pack with the name of the pack you are adding. For example, if you are adding the `csi-aws-ebs` pack, the data resource would look like the following:

```hcl
data "spectrocloud_pack" "csi-aws-ebs" {
  name         = "csi-aws-ebs"
  version      = "1.22.0"
  registry_uid = data.spectrocloud_registry.public_registry.id
}
```

2. If adding a new pack version, add a new data resource for the pack. 

```hcl
data "spectrocloud_pack" "csi-aws-ebs-1-24" {
  name         = "csi-aws-ebs"
  version      = "1.24.0"
  registry_uid = data.spectrocloud_registry.public_registry.id
}
```


3. In the cluster profile resource `spectrocloud_cluster_profile`, add a new `pack {}` block for each pack version. Assign a name and version to the cluster profile.  

```hcl
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
```

4. Repeat the above steps for each cluster profile version.


5. Reference the desired cluster profile in the `spectrocloud_cluster` resource. 