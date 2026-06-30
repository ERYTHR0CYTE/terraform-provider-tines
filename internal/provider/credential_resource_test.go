package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccTinesCredential_Text(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create the Tines Credential.
				Config: providerConfig + testAccCreateTinesCredentialText(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectUnknownValue(
							"tines_credential.test_example_text",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"tines_credential.test_example_text",
							tfjsonpath.New("team_id"),
							knownvalue.Int64Exact(30906),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"tines_credential.test_example_text",
						tfjsonpath.New("id"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						"tines_credential.test_example_text",
						tfjsonpath.New("name"),
						knownvalue.StringExact("Terraform Test Text Credential"),
					),
					statecheck.ExpectKnownValue(
						"tines_credential.test_example_text",
						tfjsonpath.New("mode"),
						knownvalue.StringExact("TEXT"),
					),
				},
			},
			{
				// Rotate the Tines Credential value.
				Config: providerConfig + testAccUpdateTinesCredentialText(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction("tines_credential.test_example_text", plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue("tines_credential.test_example_text", tfjsonpath.New("id"), knownvalue.NotNull()),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"tines_credential.test_example_text",
						tfjsonpath.New("value"),
						knownvalue.StringExact("rotated_secret_value"),
					),
				},
			},
			{
				// Import the existing Tines Credential. The secret value is never
				// returned by the API, so it cannot be verified on import.
				ResourceName:            "tines_credential.test_example_text",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})
}

func testAccCreateTinesCredentialText() string {
	return `
resource "tines_credential" "test_example_text" {
	team_id = 30906
	name = "Terraform Test Text Credential"
	value = "initial_secret_value"
}
	`
}

func testAccUpdateTinesCredentialText() string {
	return `
resource "tines_credential" "test_example_text" {
	team_id = 30906
	name = "Terraform Test Text Credential"
	value = "rotated_secret_value"
}
	`
}
