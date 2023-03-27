CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
MODULE=github.com/mokan-r/go-cli-tools

all: build

myXargs: build_myXargs
	@${BINDIR}/myFind -f -ext 'log' /var/log | ${BINDIR}/myXargs ${BINDIR}/myWc -l

myRotate: build_myRotate
	@${BINDIR}/myRotate


build: build_myFind build_myWc build_myXargs build_myRotate

build_myFind: bindir
	go build -o ${BINDIR}/myFind ${MODULE}/cmd/find

build_myWc: bindir
	@go build -o ${BINDIR}/myWc ${MODULE}/cmd/wc

build_myXargs: bindir
	@go build -o ${BINDIR}/myXargs ${MODULE}/cmd/xargs

build_myRotate: bindir
	@go build -o ${BINDIR}/myRotate ${MODULE}/cmd/rotate


bindir:
	@mkdir -p ${BINDIR}


clean:
	@rm -rf bin
