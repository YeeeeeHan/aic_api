proto_path = $(GOPATH)/src/github.com/aic/http_idl/idl
idl_path = $(proto_path)/aic/serv/aic_api.proto
http_idl_gen_repo = github.com/aic/http_idl_gen
model_file = biz/model/aic_api.pb.go
http_idl_gen_branch ?= master

init:
	hz new -I=$(proto_path) -idl=$(idl_path) -force --handler_by_method=true
	go mod init


# By default, run `make update-http-idl` which will go get http_idl_gen from remote master
# If you want go get http_idl_gen from a different remote branch, run:
# `run make update-http-idl http_idl_gen_branch="{branch-name}"
# e.g. `make update-http-idl http_idl_gen_branch="feature/lucas"`
update-http-idl:
	hz update -I=$(proto_path) -idl=$(idl_path) --handler_by_method=true
	rm -rf biz/vendor


