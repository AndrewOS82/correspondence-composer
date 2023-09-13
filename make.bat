@ECHO off
setlocal

set GOCMD=go
set GOBUILD=%GOCMD% build
set GORUN=%GOCMD% run
set NAME=correspondence-composer
set ENTRY_PATH=cmd/%NAME%/main.go

IF %1.==. GOTO NoArgs

GOTO %1

:run
	%GORUN% %ENTRY_PATH%
  GOTO End

:build
	%GOBUILD% -o bin/%NAME% %ENTRY_PATH%
  GOTO End

:kafka-start
	docker-compose -f kafka.yml up -d --remove-orphans
  GOTO End

:kafka-stop
	docker-compose -f kafka.yml down
  GOTO End

:docker-build
	docker-compose -f docker-compose-local.yml build correspondence-composer
  GOTO End

:docker-run
	docker-compose -f docker-compose-local.yml up
  GOTO End

:test
	%GOCMD% test -v ./... -p 1
  GOTO End

:lint
	golangci-lint run -c .golangci.yml --fix
  GOTO End

:generate-xsd-types
	xgen -i "./xsds/$(xsd).xsd" -o "./models/generated/$(output).go" -l Go
  GOTO End


:NoArgs
  ECHO No arguments passed

:End
  ECHO Exiting

endlocal