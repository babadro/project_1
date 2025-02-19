package main

import (
	sdk "github.com/Picovoice/picovoice/sdk/go/v2"
	rhino "github.com/Picovoice/rhino/binding/go/v2"
)

func main() {

	wakeWordCallback := func() {
		// wake word detected
	}

	inferenceCallback := func(inference rhino.RhinoInference) {
		if inference.IsUnderstood {
			intent := inference.Intent
			slots := inference.Slots
			// take action based on inferred intent and slot values
		} else {
			// handle unsupported commands
		}
	}

	picovoice := sdk.NewPicovoice(
		"${ACCESS_KEY}",
		"${KEYWORD_FILE_PATH}",
		wakeWordCallback,
		"${CONTEXT_FILE_PATH}",
		inferenceCallback)
	err := picovoice.Init()
}
