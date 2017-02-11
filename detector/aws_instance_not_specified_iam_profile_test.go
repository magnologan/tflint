package detector

import (
	"reflect"
	"testing"

	"github.com/wata727/tflint/config"
	"github.com/wata727/tflint/issue"
)

func TestDetectAwsInstanceNotSpecifiedIAMProfile(t *testing.T) {
	cases := []struct {
		Name   string
		Src    string
		Issues []*issue.Issue
	}{
		{
			Name: "iam_instance_profile is not specified",
			Src: `
resource "aws_instance" "web" {
    instance_type = "t2.2xlarge"
}`,
			Issues: []*issue.Issue{
				&issue.Issue{
					Type:    "NOTICE",
					Message: "\"iam_instance_profile\" is not specified. If you want to change it, you need to recreate it",
					Line:    2,
					File:    "test.tf",
				},
			},
		},
		{
			Name: "iam_instance_profile is specified",
			Src: `
resource "aws_instance" "web" {
    instance_type = "t2.micro"
    iam_instance_profile = "test"
}`,
			Issues: []*issue.Issue{},
		},
	}

	for _, tc := range cases {
		var issues = []*issue.Issue{}
		TestDetectByCreatorName(
			"CreateAwsInstanceNotSpecifiedIAMProfileDetector",
			tc.Src,
			"",
			config.Init(),
			config.Init().NewAwsClient(),
			&issues,
		)

		if !reflect.DeepEqual(issues, tc.Issues) {
			t.Fatalf("Bad: %s\nExpected: %s\n\ntestcase: %s", issues, tc.Issues, tc.Name)
		}
	}
}