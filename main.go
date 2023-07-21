package main

import (
	"github.com/mrkhay/creative-quill-backend/app"
)

// @title Creatives Quill Backend
// @version 0.1
// @description The mad backend 🔥🔥🔥🔥🔥
// @contact.name KhAy
// @license.name MIT
// @host localhost:3400
// @BasePath /
func main() {
	// project init
	err := app.SetupAndRun()
	if err != nil {
		panic(err)
	}

}
