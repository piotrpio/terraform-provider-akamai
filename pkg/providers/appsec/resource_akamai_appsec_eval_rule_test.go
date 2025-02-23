package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
)

func TestAccAkamaiEvalRule_res_basic(t *testing.T) {
	t.Run("match by EvalRule ID", func(t *testing.T) {
		client := &mockappsec{}

		cu := appsec.UpdateEvalRuleResponse{}
		expectJSU := compactJSON(loadFixtureBytes("testdata/TestResEvalRule/EvalRuleUpdated.json"))
		json.Unmarshal([]byte(expectJSU), &cu)

		cr := appsec.GetEvalRuleResponse{}
		expectJS := compactJSON(loadFixtureBytes("testdata/TestResEvalRule/EvalRule.json"))
		json.Unmarshal([]byte(expectJS), &cr)

		cd := appsec.UpdateEvalRuleResponse{}
		expectJSD := compactJSON(loadFixtureBytes("testdata/TestResEvalRule/EvalRule.json"))
		json.Unmarshal([]byte(expectJSD), &cd)

		config := appsec.GetConfigurationResponse{}
		expectConfigs := compactJSON(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"))
		json.Unmarshal([]byte(expectConfigs), &config)

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&config, nil)

		client.On("GetEvalRule",
			mock.Anything, // ctx is irrelevant for this test
			appsec.GetEvalRuleRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", RuleID: 12345},
		).Return(&cr, nil)

		conditionExceptionJSON := loadFixtureBytes("testdata/TestResEvalRule/ConditionException.json")
		client.On("UpdateEvalRule", mock.Anything, // ctx is irrelevant for this test
			appsec.UpdateEvalRuleRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", Action: "alert", RuleID: 12345, JsonPayloadRaw: conditionExceptionJSON},
		).Return(&cu, nil)

		client.On("UpdateEvalRule",
			mock.Anything, // ctx is irrelevant for this test
			appsec.UpdateEvalRuleRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", RuleID: 12345, Action: "none"},
		).Return(&cd, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestResEvalRule/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_appsec_eval_rule.test", "id", "43253:AAAA_81230:12345"),
						),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
