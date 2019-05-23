#!/bin/bash
set -e
set -x
go get github.com/vektra/mockery/cmd/mockery

mockery -dir=gen/pb-go/flyteidl/service/ -name=AdminServiceClient -output=clients/go/admin/mocks
mockery -dir=gen/pb-go/flyteidl/legacydiscovery/ -name=ArtifactsClient -output=clients/go/legacydiscovery/mocks