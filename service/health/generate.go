package health

// we should generate the sdk
//go:generate clang-format -i proto/health.proto
//go:generate buf generate --path proto/health.proto
//go:generate buf lint

// we need to modify the swagger file
//go:generate strava apply
