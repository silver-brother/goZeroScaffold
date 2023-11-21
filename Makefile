.PHONY:api
api:
	@goctl1.6 api format -dir=.
	@goctl1.6 api go -api desc.api -dir . -style go_zero  -home  ./goctl

.PHONY:docs
docs:
	@goctl1.6 api plugin -plugin goctl-swagger="swagger -filename docs/docs.json" -api desc.api -dir .

.PHONY:model
model:
	@goctl1.6 --db="mysql" --dsn="root:Ktza_1234@tcp(39.97.101.114:3306)/ois?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" --tables="user" --outPath="./internal/model/query" --modelPkgName="entity" --fieldNullable=true --fieldWithIndexTag=true --fieldWithTypeTag=true