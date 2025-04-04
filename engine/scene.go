package engine

type Context struct {
	VirtualWidth, VirtualHeight int32
}

type scene interface {
	Load(Context)                // called when this Scene is switched to
	Update(Context) (quit bool)  // called every frame
	Unload() (nextSceneID string) // called after Quit returns true
}
