# Adapters

adapters should be concrete implemetation of things.

- DO NOT MAKE libraries/wrappers here or reflection of existing!
- Try to AVOID concurency or any hidden complexity which your `caller` can not control.


Keeping things inside adapters simple and as single package can drive 
- higher abstraction
- integration testing easier if more knobs (inputs, outputs) you need to verify
- easier to understand all app dependencies

This also forces you not to put too much logic and end up with too much files, or 
having shared interface implementations with same file name. ex `file.go`. 