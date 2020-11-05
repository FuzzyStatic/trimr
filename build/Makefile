PROGRAM_DIR=${GOPATH}/src/github.com/FuzzyStatic/trimr/cmd
BINARY_LINUX_AMD64=trimr-linux
BINARY_DARWIN_AMD64=trimr-darwin

BUILD_TIME=`date +%FT%T%z`
BUILD_HOST=`hostname`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.BuildHost=${BUILD_HOST} -s -w"
BUILD_LINUX=env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_LINUX_AMD64} -v ${PROGRAM_DIR}/*.go
BUILD_DARWIN=env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_DARWIN_AMD64} -v ${PROGRAM_DIR}/*.go
UPX=upx --brute -q

.DEFAULT_GOAL:= all

all:
	${BUILD_LINUX}
	${BUILD_DARWIN}

all-upx:
	${BUILD_LINUX}
	${UPX} ${BINARY_LINUX_AMD64}
	${BUILD_DARWIN}
	${UPX} ${BINARY_DARWIN_AMD64}

linux:
	${BUILD_LINUX}

darwin:
	${BUILD_DARWIN}

linux-upx:
	${BUILD_LINUX}
	${UPX} ${BINARY_LINUX_AMD64}

darwin-upx:
	${BUILD_DARWIN}
	${UPX} ${BINARY_DARWIN_AMD64}

clean:
	if [ -f ${BINARY_LINUX_AMD64} ] ; then rm ${BINARY_LINUX_AMD64} ; fi
	if [ -f ${BINARY_DARWIN_AMD64} ] ; then rm ${BINARY_DARWIN_AMD64} ; fi
