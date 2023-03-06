package hypolashlckping

import (
	"strconv"

	helpers "github.com/hypolas/hypolashlckhelpers"
)

// hypolas healthcheck implementation
var (
	log    = helpers.NewLogger()
	result = helpers.NewResult()
)

// Your variables
var (
	host              = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_PING_HOST", "")
	pingCount         = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_PING_COUNT", "4")
	maxPercentLost, _ = strconv.Atoi(helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_PING_MAX_PERCENT_LOST", "0"))
)
