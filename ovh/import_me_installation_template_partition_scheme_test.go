package ovh

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMeInstallationTemplatePartitionScheme_importBasic(t *testing.T) {
	installationTemplate := acctest.RandomWithPrefix(test_prefix)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCredentials(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccMeInstallationTemplatePartitionSchemeResourceConfig_Basic, installationTemplate),
			},
			{
				ResourceName:      "ovh_me_installation_template_partition_scheme.scheme",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/myscheme", installationTemplate),
			},
		},
	})
}
