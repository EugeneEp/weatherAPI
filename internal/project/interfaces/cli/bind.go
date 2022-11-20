package cli

import (
	"github.com/spf13/cobra"
	"projectname/internal/project"
	"projectname/internal/project/interfaces/cli/cmd/service"
)

func Bind(root *cobra.Command, app project.Project) {
	srv := &cobra.Command{Use: "service"}
	srv.AddCommand(service.Start(app))
	srv.AddCommand(service.Stop(app))
	srv.AddCommand(service.Restart(app))
	srv.AddCommand(service.Install(app))
	srv.AddCommand(service.Uninstall(app))
	root.AddCommand(srv)
}
