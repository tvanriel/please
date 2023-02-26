/*
Copyright © 2023 Ted van Riel

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tvanriel/please/app"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "please",
	Short: "Todo list management with markdown savefile",
	Run: func(cmd *cobra.Command, args []string) {
		app.PrintTodo(viper.GetString("todo_filename"))
	},
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		app.Add(viper.GetString("todo_filename"), strings.Join(args, " "))
	},
}

var finishCommand = &cobra.Command{
	Use:   "finish",
	Short: "Cross off an item from your todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Finish(viper.GetString("todo_filename"), args[0])
	},
}

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Remove an itema from your todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Delete(viper.GetString("todo_filename"), strings.Join(args, " "))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/please.yaml)")
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(finishCommand)
	rootCmd.AddCommand(deleteCommand)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".please" (without extension).
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("$HOME/.config")
		viper.SetConfigType("yaml")
		viper.SetConfigName("please")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
