package main

import (
	"github.com/mrkhay/creative-insider-backend/app"
)

// @title Creative Insider Backend
// @version 0.1
// @description The mad backend 🔥🔥🔥🔥🔥
// @contact.name KhAy
// @license.name MIT
// @host localhost:3000
// @BasePath /
func main() {
	// project init
	err := app.SetupAndRun()
	if err != nil {
		panic(err)
	}

}
