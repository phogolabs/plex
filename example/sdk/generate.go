package sdk

// we should generate the sdk
//go:generate clang-format -i proto/example.proto
//go:generate buf generate --path proto/example.proto
//go:generate buf lint

// we need to modify the swagger file
//go:generate strava apply
