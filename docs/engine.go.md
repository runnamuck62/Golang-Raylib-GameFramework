# Package engine

# Types

### type **Config**
```go
type Config struct{}
```
Struct used to pass configuration options to the [Run function](#func-run). eg: Window Size, config flags, etc...

**Example:**
```go
type Config struct {
	WindowTitle string
}
```

---
### type **Context**
```go
type Context struct{}
```
Struct used to pass information to individual scenes. eg: Assets, Save Files, Camera, etc...

**Example:**
```go
type Context struct {
	Assets fs.FS
	IsWeb bool
}
```

---
### type **scene**
```go
type scene interface{}
```
Interface that holds all of the common methods of all scenes. Every Scene must have the methods used in this interface.

**Example:**
```go
type scene interface{
	Load(Context)
	Update(Context) (unload bool) 
	Unload(Context) (nextSceneID string)
}
```

---
### type **Scenes**
```go
type Scenes map[string]scene
```
Map used to retrieve a scene based on the Scene ID from RegisteredScenes.go 

**Example:**
```go
//map defined in RegisteredScenes.go
var Registered = engine.Scenes{
	"start": &start.Scene{},
	"cube": &cube.Scene{},
}
```

```go
//check for Scene named start in Scenes parameter
func Run(scenes Scenes, cfg Config, Assets fs.FS) error {
	ActiveSceneId := "start"
	ActiveScene, ok := scenes[ActiveSceneId]
	...
```

---

# Functions

### func **Run**
```go
func Run(scenes Scenes, cfg Config, Assets fs.FS) error 
```
Main game loop called in Main.go. Handles loading and unloading scenes and assets, initializing configuration options, and game updating. 

---
### func centerWindow
```go
func centerWindow()
```
Open's game window in the center of the screen

---


