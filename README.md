
# Hexagonal

but keeping things flat and simple.

## Domain

root package is domain or it can be severall packages. This mostly our focus and keeping it 
flat as possible (less packages) will have more benefit like testing and import cycle malware (Creating extra package to avoid import cycle).

No `internal` as we have dedicated `proto` package.

## Adapters package

adapters package is where are all ports from domain package implemented. 
We always have to talk more outside world than other way arround. 

AVOIDING subpackage would be better as normally this should be small interfaces and IO logic.
This also forces you not to put too much logic and end up with too much files.

It also avoids file duplication like `repository.go`

## Ports package

DOES NOT EXISTS!

it is coupled with domain code (AS DEPENDECNY INJECTION) and therefore it should
be close to code where it is used.

It can be seperate file like `ports.go` to define structures or all ports that are used within domain
but they should always be in same package and closer to where are used. 

So explicitily naming thing like ports or files should be avoided.

## Proto package

API protos (client, server). SDK for your service.
- openapi yaml
- grps protobuffer

## CMD

`main.go` binaries.
