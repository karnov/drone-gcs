package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "gcs plugin"
	app.Usage = "gcs plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "application-credentials",
			Usage:  "google application credentials json file",
			EnvVar: "GOOGLE_APPLICATION_CREDENTIALS_CONTENTS",
		},
		cli.StringFlag{
			Name:   "bucket",
			Usage:  "gcs bucket",
			EnvVar: "PLUGIN_BUCKET,GCS_BUCKET",
		},
		cli.StringFlag{
			Name:   "acl",
			Usage:  "upload files with acl",
			Value:  "private",
			EnvVar: "PLUGIN_ACL",
		},
		cli.StringFlag{
			Name:   "source",
			Usage:  "upload files from source folder",
			EnvVar: "PLUGIN_SOURCE",
		},
		cli.StringFlag{
			Name:   "target",
			Usage:  "upload files to target folder",
			EnvVar: "PLUGIN_TARGET",
		},
		cli.StringFlag{
			Name:   "strip-prefix",
			Usage:  "strip the prefix from the target",
			EnvVar: "PLUGIN_STRIP_PREFIX",
		},
		cli.StringSliceFlag{
			Name:   "exclude",
			Usage:  "ignore files matching exclude pattern",
			EnvVar: "PLUGIN_EXCLUDE",
		},
		cli.BoolFlag{
			Name:   "dry-run",
			Usage:  "dry run for debug purposes",
			EnvVar: "PLUGIN_DRY_RUN",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
		cli.BoolFlag{
			Name:   "compress",
			Usage:  "gzip files before they are uploaded",
			EnvVar: "PLUGIN_COMPRESS",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Credentials: c.String("application-credentials"),
		Bucket:      c.String("bucket"),
		Access:      c.String("acl"),
		Source:      c.String("source"),
		Target:      c.String("target"),
		StripPrefix: c.String("strip-prefix"),
		Exclude:     c.StringSlice("exclude"),
		DryRun:      c.Bool("dry-run"),
		Compress:    c.Bool("compress"),
	}

	return plugin.Exec()
}
