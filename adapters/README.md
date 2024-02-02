# Adapters

adapters should be concrete implemetation of things. 
DO NOT MAKE libraries/wrappers here or reflection of existing!

Keeping things inside adapters as package can drive 
- smaller interfaces
- integration testing easier if more knobs (inputs, outputs) you need to verify
- less packages or single that reveals all app dependencies

