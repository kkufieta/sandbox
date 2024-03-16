// Package main is a module with a fake SLAM service model.
package main

import (
	"context"
	"strings"

	bunnySlam "github.com/kkufieta/sandbox/bunnyslam"
	fakeSlam "github.com/kkufieta/sandbox/fakeslam"
	robotSlam "github.com/kkufieta/sandbox/robotslam"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/services/slam"
	"go.viam.com/utils"
)

// Versioning variables which are replaced by LD flags.
var (
	Version     = "development"
	GitRevision = ""
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("fakeSlam"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	var versionFields []interface{}
	if Version != "" {
		versionFields = append(versionFields, "version", Version)
	}
	if GitRevision != "" {
		versionFields = append(versionFields, "git_rev", GitRevision)
	}
	if len(versionFields) != 0 {
		logger.Infow(fakeSlam.Model.String(), versionFields...)
	} else {
		logger.Info(fakeSlam.Model.String() + " built from source; version unknown")
	}

	if len(args) == 2 && strings.HasSuffix(args[1], "-version") {
		return nil
	}

	// Instantiate the module
	fakeModule, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Add the fakeSlam model to the module
	if err = fakeModule.AddModelFromRegistry(ctx, slam.API, fakeSlam.Model); err != nil {
		return err
	}

	// Add the bunnySlam model to the module
	if err = fakeModule.AddModelFromRegistry(ctx, slam.API, bunnySlam.Model); err != nil {
		return err
	}

	// Add the robotSlam model to the module
	if err = fakeModule.AddModelFromRegistry(ctx, slam.API, robotSlam.Model); err != nil {
		return err
	}

	// Start the module
	err = fakeModule.Start(ctx)
	defer fakeModule.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
