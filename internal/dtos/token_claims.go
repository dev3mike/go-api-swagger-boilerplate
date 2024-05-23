package dtos

// ContextKey is a custom type to avoid key collisions in context
type ContextKey string

const (
	TokenClaimsKey ContextKey = "TokenClaims"
)

type UserClaims struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	CreatedAt     uint32 `json:"created_at"`
	UpdatedAt     uint32 `json:"updated_at"`
	EmailVerified bool   `json:"email_verified"`
}
