# Raylib Golang "GameFramework"

A very basic setup for making games with raylib. You are expected
to modify this codebase to suit your needs.


Featuring Seperate "Scenes" with their own state


### Getting started

Just create a repository using this as the template.

this template includes a spinning cube demo. Try running it.
See [raylib-go](https://github.com/gen2brain/raylib-go/) for instructions.


put your scenes inside of its own seperate package:
`scenes/<scenename>` for example:
```
scenes/
├── cube
│   └── cube.go // spinning cube
├── register.go // register your scenes in here
└── start // start scene (loaded first)
    ├── start.go
    └── systems.go
```


Scenes can be registered inside of `scenes/register.go`
```go
// register all your scenes in here
var Registered = engine.Scenes{
	"start": &start.Scene{},
	"cube": &cube.Scene{},
}
```


engine.Scenes is just a map from a "Scene Id" (string) to engine.scene:
```go
// a scene must implement the engine.scene interface
type scene interface {
	Load(Context)                        // called when this Scene is switched to
	Update(Context) (unload bool)        // called every frame
	Unload(Context) (nextSceneID string) // called after Update returns true. Switches to nextSceneID
}
```


here is a blank scene that will never quit
```go
package myscene

import "GameFrameworkTM/engine"

type Scene struct {
}
// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
}
// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	return false // if true is returned, Unload is called
}
// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "someOtherSceneId" 
}
```

### Code explanation

`main.go`

This file isn't important, you can only change the window title here.

`engine/engine.go`

The `engine.Run` function initializes a window and centers it. Its only 100 lines.
looking at the code comments is recommended.

By the way, engine.Context is supposed to be used by you to implement a feature
```go
// info to pass to scenes
// eg. a camera, game map, or save file
type Context struct {
	SomeData any
}
```


`scenes/start/start.go`
this is the main menu scene, press enter to select an option.



`scenes/cube/cube.go`
this scene features a spinning cube (wow!)


### What next?
Add what you need. Here are some features you could implement:

- [Letterboxing (black bars)](https://www.raylib.com/examples/core/loader.html?name=core_window_letterbox)
- Asset management
- Save Files
- anything else your game may need

