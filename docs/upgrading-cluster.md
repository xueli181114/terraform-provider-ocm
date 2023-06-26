# Updating or Upgrading your ROSA cluster

You can update or upgrade your cluster using Terraform.

## Prerequisites

1. You created your [account roles using Terraform](../examples/create_rosa_cluster/create_rosa_sts_cluster/classic_sts/account_roles/README.md).
1. You created your cluster using Terraform. This cluster can either have [a managed OIDC configuration](../examples/create_rosa_cluster/create_rosa_sts_cluster/oidc_configuration/cluster_with_managed_oidc_config/README.md) or [an unmanaged OIDC configuration](../examples/create_rosa_cluster/create_rosa_sts_cluster/oidc_configuration/cluster_with_unmanaged_oidc_config/README.md).

## Upgrading your cluster

To upgrade your ROSA cluster to another version, export the following variable then run `terraform apply`.

1. Export the `TF_VAR_openshift_version` with the intended version. Your value must be prepended with `openshift-v` to succeed.
    ```
    export TF_VAR_openshift_version=<version_number>
    ```
1. If you choose to upgrade to a new version that necessitates approval, particularly when transitioning between major Y-Streams, you may be asked to provide administrative confirmation regarding significant modifications for your cluster. In such a scenario, during the initial upgrade attempt, you will receive an error message that provides guidance on the necessary changes. It is crucial to diligently follow those instructions and only acknowledge completion of the requirements by adding the "upgrade_acknowledgements_for" attribute to your resource, specifying the target version. For instance, if you are upgrading from 4.11.43 to 4.12.21, you should use '4.12' as the value for this variable.
    ```
    upgrade_acknowledgements_for = <version_acknowledgement>
    ```
1. Run `terraform apply` to upgrade your cluster.

## OpenShift documentation

 - [Upgrading ROSA clusters with STS](hhttps://docs.openshift.com/rosa/upgrading/rosa-upgrading-sts.html)