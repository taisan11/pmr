package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Commands: []*cli.Command{
			{
				Name:      "install",
				Usage:     "install packages",
				ArgsUsage: "package name",
				Action: func(c *cli.Context) error {
					config, err := loadConfig()
					if err != nil {
						return fmt.Errorf("failed to load config: %w", err)
					}
					for _, pmname := range config.Level {
						pm, err := loadPM(pmname)
						if err != nil {
							return fmt.Errorf("failed to load PM: %w", err)
						}
						if pm.Install == nil {
							fmt.Printf("Installing %s...\n", pm.Install)
							continue
						}
						args := append(strings.Split((*pm.Install), " "), c.Args().Get(0))
						cmd := exec.Command(pm.Name, args...)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						if err := cmd.Run(); err != nil {
							return fmt.Errorf("failed to run command: %w", err)
						}
					}
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "update packages",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "Update all packages",
					},
				},
				Action: func(c *cli.Context) error {
					config, err := loadConfig()
					if err != nil {
						return fmt.Errorf("failed to load config: %w", err)
					}
					fmt.Printf("Config: %+v\n", config)
					for _, pmname := range config.Level {
						pm, err := loadPM(pmname)
						if err != nil {
							return fmt.Errorf("failed to load PM: %w", err)
						}
						if c.Bool("all") {
							if pm.UpdateAll == nil {
								fmt.Printf("Updating all packages with %s...\n", pm.UpdateAll)
								continue
							}
							cmd := exec.Command(pm.Name, strings.Split((*pm.UpdateAll), " ")...)
							cmd.Stdout = os.Stdout
							cmd.Stderr = os.Stderr
							if err := cmd.Run(); err != nil {
								return fmt.Errorf("failed to run command: %w", err)
							}
						} else {
							if pm.Update == nil {
								fmt.Printf("Updating package %s...\n", pm.Update)
								continue
							}
							fmt.Println("Updating selected packages...")
						}
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
