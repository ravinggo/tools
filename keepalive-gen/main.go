package main

import (
	"flag"

	"github.com/spf13/pflag"
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/klog/v2"

	"github.com/ravinggo/tools/keepalive-gen/args"
	"github.com/ravinggo/tools/keepalive-gen/generators"
)

func main() {
	klog.InitFlags(nil)
	args := args.New()

	args.AddFlags(pflag.CommandLine)
	flag.Set("logtostderr", "true")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := args.Validate(); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	myTargets := func(context *generator.Context) []generator.Target {
		return generators.GetTargets(context, args)
	}

	// Run it.
	if err := gengo.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		myTargets,
		generators.IgnoreTag,
		pflag.Args(),
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}
	klog.V(2).Info("Completed successfully.")
}
