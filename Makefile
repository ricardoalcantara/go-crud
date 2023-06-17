server:
	go run cmd/idp/idp.go
watch:
	CompileDaemon -build="go build cmd/idp/idp.go" -command="go run cmd/idp/idp.go" --color
migration:
	go run cmd/migration/migration.go