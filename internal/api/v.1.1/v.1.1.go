package v11

import (
	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/api/handlers/flowid"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.1")
	{
		flowid.ApplyRoutes(v1)
		// authenticate.ApplyRoutes(v1)
		// profile.ApplyRoutes(v1)
		// delegatereviewcreation.ApplyRoutes(v1)
		// deletereview.ApplyRoutes(v1)
		// status.ApplyRoutes(v1)
		// feedback.ApplyRoutes(v1)
		// waitlist.ApplyRoutes(v1)
		// stats.ApplyRoutes(v1)
		// getreviews.ApplyRoutes(v1)
		// getreviewerdetails.ApplyRoutes(v1)
		// domain.ApplyRoutes(v1)
		// report.ApplyRoutes(v1)
		// account.ApplyRoutes(v1)
		// siteinsights.ApplyRoutes(v1)
		// subscription.ApplyRoutes(v1)
		// summary.ApplyRoutes(v1)
		// sdkauthentication.ApplyRoutes(v1)
		// leaderboard.ApplyRoutes(v1)
		// nftcontract.ApplyRoutes(v1)
		// referral.ApplyReferraAccountlRoutes(v1)

	}
}
