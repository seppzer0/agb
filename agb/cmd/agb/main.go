package main

import (
	"fmt"
	"strings"

	"agb/command"

	"github.com/spf13/cobra"
)

// newBuildCmd registers the "build" command.
func newBuildCmd() *cobra.Command {
	var (
		linuxKernelVersion string
		androidVersion     int
		patchVersion       string
		defconfigPath      string
		modulesPath        string
		clangUrl           string
		sourceLocation     string
		kernelSu           bool
	)

	command := &cobra.Command{
		Use:   "build",
		Short: "Build the Android kernel.",
		ValidArgs: []string{
			"linux-kernel-version",
			"android-version",
			"patch-version",
			"defconfig-path",
			"modules-path",
			"clang-url",
			"source-location",
			"kernelsu",
		},
		Args: cobra.MatchAll(cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			command_obj := command.NewBuildCommand(
				linuxKernelVersion,
				androidVersion,
				patchVersion,
				defconfigPath,
				clangUrl,
				sourceLocation,
				modulesPath,
				kernelSu,
			)

			res := command_obj.Execute()
			return res
		},
	}

	flags := command.Flags()
	flags.stringVarP(
		&linuxKernelVersion,
		"linux-kernel-version",
		"l",
		0,
		"Linux kernel version number (required)",
	)
	flags.IntVarP(
		&androidVersion,
		"android-version",
		"a",
		0,
		"Android version number (required)",
	)
	flags.StringVarP(
		&patchVersion,
		"patch-version",
		"p",
		"",
		"patch version number",
	)
	flags.StringVarP(
		&defconfigPath,
		"defconfig-path",
		"d",
		"",
		"custom path to defconfig",
	)
	flags.StringVarP(
		&modulesPath,
		"modules-path",
		"m",
		"",
		"custom path to kernel modules",
	)
	flags.StringVarP(
		&clangUrl,
		"clang-url",
		"c",
		"https://android.googlesource.com/platform/prebuilts/clang/host/linux-x86/+archive/eed2fff8b93ce059eea7ccd8fc5eee37f8adb432/clang-r458507.tar.gz",
		"path to a Clang pre-build zip",
	)
	flags.StringVarP(
		&sourceLocation,
		"source-location",
		"s",
		"",
		"custom location to GKI kernel sources (local path or URL to git repo)",
	)
	flags.BoolVarP(
		&kernelSu,
		"kernelsu",
		"k",
		false,
		"optional KernelSU support (for GKI mode)",
	)

	command.MarkFlagRequired("linux-kernel-version")
	command.MarkFlagRequired("android-version")

	return command
}

// newCleanCmd registers the "clean" command.
func newCleanCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Clean the build environment.",
		Run: func(cmd *cobra.Command, args []string) {
			// check if anything at all was provided for the query
			if len(args) < 1 {
				fmt.Println("WIP")
			}

			fmted_args := strings.Join(args[:], " ")
			fmt.Println(fmted_args)
		},
	}
}

// newVersionCmd registers the "version" command.
func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information.",
		RunE: func(cmd *cobra.Command, args []string) error {
			command_obj := command.NewVersionCommand()

			res := command_obj.Execute()
			return res
		},
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "agb",
		Short: "Android GKI Builder.",
	}

	rootCmd.AddCommand(newBuildCmd())
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newCleanCmd())

	rootCmd.Execute()
}
