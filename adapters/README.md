# Adapters

adapters should be concrete implemetation of things.

- DO NOT MAKE libraries or some big wrappers here!
- Try to AVOID concurency or any hidden complexity which your `caller` from root can not control.
- Always provide way for cancelation and return errors instead logging them


Keeping things inside adapters simple and as single package can drive 
- integration testing easier if more knobs (inputs, outputs) you need to verify
- easier to understand all app dependencies
- better abstraction 

This also forces you not to put too much logic and end up with too much files, or 
having shared interface implementations with same file name. ex `file.go`. 