package env

// Env is the environment an application runs in.
type Env string

var (
	// Development environment
	Development = Env("development")

	// Testing environment
	Testing = Env("testing")

	// Staging environment
	Staging = Env("staging")

	// Production environment
	Production = Env("production")
)

// IsDevelopment is whether the environment is development (or empty).
func (env Env) IsDevelopment() bool {
	return string(env) == "" || env == Development
}

// IsTesting is whether the environment is testing.
func (env Env) IsTesting() bool {
	return env == Testing
}

// IsStaging is whether the environment is staging.
func (env Env) IsStaging() bool {
	return env == Staging
}

// IsProd is whether the environment is production.
func (env Env) IsProduction() bool {
	return env == Production
}
