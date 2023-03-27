CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.49.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
MODULE=github.com/mokan-r/day02

all: build

myXargs: build_myXargs
	@${BINDIR}/myFind -f -ext 'log' /var/log | ${BINDIR}/myXargs ${BINDIR}/myWc -l

myRotate: build_myRotate
	@${BINDIR}/myRotate

precommit: format build lint
	@echo "OK"

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

build: build_myFind build_myWc build_myXargs build_myRotate

build_myFind: bindir
	go build -o ${BINDIR}/myFind ${MODULE}/cmd/ex00

build_myWc: bindir
	@go build -o ${BINDIR}/myWc ${MODULE}/cmd/ex01

build_myXargs: bindir
	@go build -o ${BINDIR}/myXargs ${MODULE}/cmd/ex02

build_myRotate: bindir
	@go build -o ${BINDIR}/myRotate ${MODULE}/cmd/ex03

lint: install-lint
	${LINTBIN} run

run:
	go run ${PACKAGE}/cmd/ex00

demo: build
	${BINDIR}/statistic < data-samples/data1

# service

bindir:
	@mkdir -p ${BINDIR}

install-lint: bindir
	@test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	@test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

clean:
	@rm -rf bin
