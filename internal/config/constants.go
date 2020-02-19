package config

// Global constants
const (
	// InlinePolicyName defines user managed inline policy
	InlinePolicyName = "custom"

	// IamManagerNamespaceName is the namespace name where iam-manager controllers are running
	IamManagerNamespaceName = "iam-manager-system"

	// IamManagerConfigMapName is the config map name for iam-manager namespace
	IamManagerConfigMapName = "iam-manager-iamroles-v1alpha1-configmap"
)

const (
	// iam policy action prefix
	propertyIamPolicyWhitelist = "iam.policy.action.prefix.whitelist"

	// iam policy to blacklist resource
	propertyIamPolicyBlacklist = "iam.policy.resource.blacklist"

	// iam policy for restricting s3 resources
	propertyIamPolicyS3Restricted = "iam.policy.s3.restricted.resource"

	// aws region
	propertyAwsRegion = "aws.region"

	// aws master role
	propertyAwsMasterRole = "aws.MasterRole"

	// user managed policies
	propertyManagedPolicies = "iam.managed.policies"

	// user managed permission boundary policy
	propertyPermissionBoundary = "iam.managed.permission.boundary.policy"
)

const (
	separator = ","
)