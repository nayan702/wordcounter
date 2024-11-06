package constants

const (
	CooldownDurationSeconds     = 300 // Cooldown time in seconds for responses with status codes 999/429/503
	MaxConcurrentRequestsPerSec = 15  // Maximum concurrent requests allowed per second by rate limiter
	RateLimiterBurstSize        = 5   // Burst size for rate limiting
	SuccessThresholdForIncrease = 100 // Number of successful requests to trigger an increase in rate limit
	RateLimiterAdjustmentPct    = 10  // Percentage adjustment for rate limiter upon success
)
