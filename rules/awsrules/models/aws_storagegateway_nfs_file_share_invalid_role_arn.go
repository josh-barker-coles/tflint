// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayNfsFileShareInvalidRoleArnRule checks the pattern is valid
type AwsStoragegatewayNfsFileShareInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayNfsFileShareInvalidRoleArnRule returns new rule with default attributes
func NewAwsStoragegatewayNfsFileShareInvalidRoleArnRule() *AwsStoragegatewayNfsFileShareInvalidRoleArnRule {
	return &AwsStoragegatewayNfsFileShareInvalidRoleArnRule{
		resourceType:  "aws_storagegateway_nfs_file_share",
		attributeName: "role_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayNfsFileShareInvalidRoleArnRule) Name() string {
	return "aws_storagegateway_nfs_file_share_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayNfsFileShareInvalidRoleArnRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsStoragegatewayNfsFileShareInvalidRoleArnRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayNfsFileShareInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayNfsFileShareInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}