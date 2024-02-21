package main

import (
	"cdnetwork/internal/cli"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	rootCmd := getRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "cdnctl",
		Short: "cdnctl is a command-line tool for managing CDNetworks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Usage: cdnctl <command> [options]")
		},
	}

	addCommands(rootCmd)
	return rootCmd
}

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getCreateCommand())
	rootCmd.AddCommand(getDeleteCommand())
	rootCmd.AddCommand(getDisableCommand())
	rootCmd.AddCommand(getEditGroupCommand())
	rootCmd.AddCommand(getCertificateCommand())
	rootCmd.AddCommand(getDomainIdCommand())
	rootCmd.AddCommand(getEnableDomainCommand())
}

func getCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create <domain>",
		Short: "Create a new domain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Creating domain: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.CreateDomain()
		},
	}
	return cmd
}

func getDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete <domain>",
		Short: "Delete an existing domain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Deleting domain: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.DeleteDomain()
		},
	}
	return cmd
}

func getDisableCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "disable <domain>",
		Short: "Disable an existing domain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Disabling domain: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.DisableDomain()
		},
	}
	return cmd
}

func getEnableDomainCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "enable <domain>",
		Short: "Disable an existing domain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Disabling domain: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.EnableDomain()
		},
	}
	return cmd
}

func getEditGroupCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update <domains>",
		Short: "Update a group of domains",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Updating domains: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.UpdateControlGroup()
		},
	}
	return cmd
}

func getDomainIdCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "getdomainid <domains>",
		Short: "Update a group of domains",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("get domain id: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.GetDomainId()
		},
	}
	return cmd
}
func getCertificateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "getcert <domains>",
		Short: "Get certificate of domains",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("get domain id: %s\n", args[0])
			cdnCommand := &cli.CdnCommand{
				DomainName: args[0],
			}
			cdnCommand.GetCertificate()
		},
	}
	return cmd
}
