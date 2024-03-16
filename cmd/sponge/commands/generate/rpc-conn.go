package generate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/zhufuyi/sponge/pkg/replacer"
)

// GRPCConnectionCommand generate grpc connection code
func GRPCConnectionCommand() *cobra.Command {
	var (
		moduleName      string // module name for go.mod
		outPath         string // output directory
		grpcServerNames string // grpc service names

		serverName             string // server name
		isSupportLargeCodeRepo bool   // is support large code repository
	)

	cmd := &cobra.Command{
		Use:   "rpc-conn",
		Short: "Generate grpc connection code",
		Long: `generate grpc connection code.

Examples:
  # generate grpc connection code
  sponge micro rpc-conn --module-name=yourModuleName --rpc-server-name=yourGrpcName

  # generate grpc connection code with multiple names.
  sponge micro rpc-conn --module-name=yourModuleName --rpc-server-name=name1,name2

  # generate grpc connection code and specify the server directory, Note: code generation will be canceled when the latest generated file already exists.
  sponge micro rpc-conn --rpc-server-name=user --out=./yourServerDir

  # if you want the generated code to support large code repository, you need to specify the parameter --support-large-code-repo=true --serverName=yourServerName
`,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			mdName, srvName, isLCR := getNamesFromOutDir(outPath)
			if mdName != "" {
				moduleName = mdName
				serverName = srvName
				isSupportLargeCodeRepo = isLCR
			} else if moduleName == "" {
				return errors.New(`required flag(s) "module-name" not set, use "sponge micro rpc-conn -h" for help`)
			}
			if isSupportLargeCodeRepo {
				if serverName == "" {
					return errors.New(`required flag(s) "server-name" not set, use "sponge micro rpc-conn -h" for help`)
				}
				serverName = convertServerName(serverName)
			}

			grpcNames := strings.Split(grpcServerNames, ",")
			for _, grpcName := range grpcNames {
				if grpcName == "" {
					continue
				}

				var err error
				var g = &grpcConnectionGenerator{
					moduleName: moduleName,
					grpcName:   grpcName,
					outPath:    outPath,

					serverName:             serverName,
					isSupportLargeCodeRepo: isSupportLargeCodeRepo,
				}
				outPath, err = g.generateCode()
				if err != nil {
					return err
				}
			}

			fmt.Printf(`
using help:
  move the folder "internal" to your project code folder.

`)
			fmt.Printf("generate \"rpc-conn\" code successfully, out = %s\n", outPath)
			return nil
		},
	}

	cmd.Flags().StringVarP(&moduleName, "module-name", "m", "", "module-name is the name of the module in the go.mod file")
	cmd.Flags().StringVarP(&grpcServerNames, "rpc-server-name", "r", "", "rpc service name, multiple names separated by commas")
	_ = cmd.MarkFlagRequired("rpc-server-name")
	cmd.Flags().StringVarP(&serverName, "server-name", "s", "", "server name")
	cmd.Flags().BoolVarP(&isSupportLargeCodeRepo, "support-large-code-repo", "l", false, "whether to support large code repository")
	cmd.Flags().StringVarP(&outPath, "out", "o", "", "output directory, default is ./rpc-conn_<time>,"+
		" if you specify the directory where the web or microservice generated by sponge, the module-name flag can be ignored")

	return cmd
}

type grpcConnectionGenerator struct {
	moduleName string
	grpcName   string
	outPath    string

	serverName             string
	isSupportLargeCodeRepo bool
}

func (g *grpcConnectionGenerator) generateCode() (string, error) {
	subTplName := "rpc-conn"
	r := Replacers[TplNameSponge]
	if r == nil {
		return "", errors.New("replacer is nil")
	}

	// setting up template information
	subDirs := []string{ // only the specified subdirectory is processed, if empty or no subdirectory is specified, it means all files
		"internal/rpcclient",
	}
	ignoreDirs := []string{} // specify the directory in the subdirectory where processing is ignored
	ignoreFiles := []string{ // specify the files in the subdirectory to be ignored for processing
		"doc.go", "serverNameExample_test.go",
	}

	r.SetSubDirsAndFiles(subDirs)
	r.SetIgnoreSubDirs(ignoreDirs...)
	r.SetIgnoreSubFiles(ignoreFiles...)
	_ = r.SetOutputDir(g.outPath, subTplName)
	fields := g.addFields(r)
	r.SetReplacementFields(fields)
	if err := r.SaveFiles(); err != nil {
		return "", err
	}

	return r.GetOutputDir(), nil
}

func (g *grpcConnectionGenerator) addFields(r replacer.Replacer) []replacer.Field {
	var fields []replacer.Field

	fields = append(fields, []replacer.Field{
		{
			Old: "github.com/zhufuyi/sponge/configs",
			New: g.moduleName + "/configs",
		},
		{
			Old: "github.com/zhufuyi/sponge/internal/config",
			New: g.moduleName + "/internal/config",
		},
		{
			Old:             "serverNameExample",
			New:             g.grpcName,
			IsCaseSensitive: true,
		},
	}...)

	if g.isSupportLargeCodeRepo {
		fs := SubServerCodeFields(r.GetOutputDir(), g.moduleName, g.serverName)
		fields = append(fields, fs...)
	}

	return fields
}
