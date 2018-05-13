package main

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/fsnotify/fsnotify"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "TaskServer",
		Short: "TaskServer",
		Long:  "TaskServer",
		Run: func(cmd *cobra.Command, args []string) {
			taskMap := viper.GetStringMap("task")
			broker := taskMap["broker"].(string)
			queue := taskMap["queue"].(string)
			resultBackend := taskMap["resultBackend"].(string)
			resultsExpireIn := taskMap["resultsExpireIn"].(int)

			Config := &config.Config{
				Broker:          broker,
				DefaultQueue:    queue,
				ResultBackend:   resultBackend,
				ResultsExpireIn: resultsExpireIn,
			}

			task, err := machinery.NewServer(Config)
			if err != nil {
				fmt.Println(err.Error())
			}

			tasks := map[string]interface{}{
				"add": func() {},
				"multiply": func() {},
				"concat": func() {},
			}
			err = task.RegisterTasks(tasks)
			worker := task.NewWorker(taskMap["consumer_tag"].(string), int(taskMap["max_worker"].(int64)))
			if err := worker.Launch(); err != nil {
				fmt.Println(err.Error())
			}

		},
	}

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func initConfig() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("task")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./h_app/h_task")
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
