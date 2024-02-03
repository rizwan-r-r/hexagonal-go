
# Hexagonal

but keeping things flat and stupid! 

## Domain

`root` package is domain. Normally you go with **main** package but if you are building something small and simple this organization is probably not needed.

Goal is mostly that this is our focus and keeping it flat and simple will have more benefit like testing and avoiding import cycle malware (Creating extra package to avoid import cycle).
Seperating things into more packages can come later (if there is really some bigger idea that can be isolated) but remember goal is to be stupid, as you mostly have defined requirements.

No `internal` as we have dedicated `proto` package for sharing external API interface. Still keeping single package and using lower case letters is mostly enough to privatize your code. 

## Adapters package

adapters package is where are all ports from domain package implemented.
We always have to talk more outside world than other way arround, or in other words your 
service is probably trying to minize complexity of speaking and scaling buisness needs.

Avoiding subpackage (BTW subpackage does not exist in GO) would be better for start as normally this should be small interfaces and IO logic but see more under [adapters](/adapters/)


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
