.PHONY: apigen

apigen:
	oapi-codegen -generate "server" -package apigen docs/openapi.yaml \
		> apigen/server.gen.go
	oapi-codegen -generate "types" -package apigen docs/openapi.yaml \
		> apigen/types.gen.go
	oapi-codegen -generate "spec" -package apigen docs/openapi.yaml \
		> apigen/spec.gen.go
