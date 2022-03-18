NAME = object
REPO = https://github.com/the-go-tool/${NAME}
AUTHOR = Kulagin Yuri (jkulvich)
EMAIL = jkulvichi@gmail.com
GITTER = https://gitter.im/the-go-tool-object/community

DATE = $(shell date +%d.%m.%y)
TIME = $(shell date +%H:%0M)
DATE_TIME_TAG = $(shell date +%d%m%y%H%0M)

COVERAGE_OUT = coverage.out

help:
	@echo "==========[ the go tool :: $(NAME) ]=========="
	@echo ""
	@echo "Date Time:  $(DATE) $(TIME) (Tag: $(DATE_TIME_TAG))"
	@echo "Author:     $(AUTHOR)"
	@echo "Email:      $(EMAIL)"
	@echo "Repository: $(REPO)"
	@echo "Gitter:     $(GITTER)"
	@echo ""
	@echo "Tool to work with unspecified-scheme objects"
	@echo "(like unmarshalled json, yaml, etc.)"
	@echo ""
	@echo "make help ---- Show this help & info"
	@echo "make check --- Run vet, tests & show coverage info"
.PHONY: hrelp

check:
	go vet
	go test -coverprofile $(COVERAGE_OUT)
	go tool cover -func $(COVERAGE_OUT)
	rm $(COVERAGE_OUT)
.PHONY: test