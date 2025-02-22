package helper

import (
	"fmt"

	"github.com/mdsauce/schelper/logger"
)

func singleOutput(prob KnownProblem, logline []byte) {
	fmt.Printf("=======PROBLEM=======\n")
	logger.Disklog.Infof(`Problem: %s

Where: %s

Suggested Next Steps: 
%s

General Steps for this type of Disruption: 
%s

`, prob.Name, logline, prob.NextSteps, prob.Disruption.GeneralSteps)
}

func metaOutput(metadata map[string]int) {
	logger.Disklog.Info("Metadata of All Problems Encountered")
	logger.Disklog.Info("------------------------------------")
	for key, val := range metadata {
		logger.Disklog.Infof("%s: %d", key, val)
	}
	fmt.Println()
}

func lifecycleOutput(cycle [6]scLifecycle) {
	logger.Disklog.Info("Sauce Connect Lifecycle")
	logger.Disklog.Info("------------------------------------")
	for i := range cycle {
		if cycle[i].reached {
			logger.Disklog.Infof("Lifecycle Stage: %s", cycle[i].stage)
			logger.Disklog.Infof("Found logline: '%s'", cycle[i].target)
		} else {
			logger.Disklog.Infof("------> Lifecycle Stage: %s not reached <------", cycle[i].stage)
			logger.Disklog.Infof("Did not find logline: '%s'", cycle[i].target)
		}
		fmt.Println()
	}
}
