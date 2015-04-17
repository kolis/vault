package appId

import (
	"testing"

	"github.com/hashicorp/vault/logical"
	logicaltest "github.com/hashicorp/vault/logical/testing"
)

func TestBackend_basic(t *testing.T) {
	logicaltest.Test(t, logicaltest.TestCase{
		Backend: Backend(),
		Steps: []logicaltest.TestStep{
			testAccStepMapAppId(t),
			testAccStepMapUserId(t),
			testAccLogin(t, ""),
			testAccLoginInvalid(t),
		},
	})
}

func TestBackend_displayName(t *testing.T) {
	logicaltest.Test(t, logicaltest.TestCase{
		Backend: Backend(),
		Steps: []logicaltest.TestStep{
			testAccStepMapAppIdDisplayName(t),
			testAccStepMapUserId(t),
			testAccLogin(t, "tubbin"),
			testAccLoginInvalid(t),
		},
	})
}

func testAccStepMapAppId(t *testing.T) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.WriteOperation,
		Path:      "map/app-id/foo",
		Data: map[string]interface{}{
			"value": "foo,bar",
		},
	}
}

func testAccStepMapAppIdDisplayName(t *testing.T) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.WriteOperation,
		Path:      "map/app-id/foo",
		Data: map[string]interface{}{
			"display_name": "tubbin",
			"value":        "foo,bar",
		},
	}
}

func testAccStepMapUserId(t *testing.T) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.WriteOperation,
		Path:      "map/user-id/42",
		Data: map[string]interface{}{
			"value": "foo",
		},
	}
}

func testAccLogin(t *testing.T, display string) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.WriteOperation,
		Path:      "login",
		Data: map[string]interface{}{
			"app_id":  "foo",
			"user_id": "42",
		},
		Unauthenticated: true,

		Check: logicaltest.TestCheckMulti(
			logicaltest.TestCheckAuth([]string{"bar", "foo"}),
			logicaltest.TestCheckAuthDisplayName(display),
		),
	}
}

func testAccLoginInvalid(t *testing.T) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.WriteOperation,
		Path:      "login",
		Data: map[string]interface{}{
			"app_id":  "foo",
			"user_id": "48",
		},
		ErrorOk:         true,
		Unauthenticated: true,

		Check: logicaltest.TestCheckError(),
	}
}
