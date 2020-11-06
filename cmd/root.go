/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "github.com/cihub/seelog"
  "github.com/hfeng101/niwo/database"
  "github.com/hfeng101/niwo/utils/config"
  "github.com/spf13/cobra"
  "os"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

  "github.com/hfeng101/niwo/router"

  "github.com/kataras/iris/v12"
  "github.com/kataras/iris/v12/middleware/logger"
  "github.com/kataras/iris/v12/middleware/recover"
)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "niwo",
  Short: "A brief description of your application",
  Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
  Run: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.niwo.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".niwo" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".niwo")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

func run(cmd *cobra.Command, args []string) {
  //init iris config
  config.InitIrisConfig()

  //init mysql & create table
  database.InitGlobalOrm()
  db := database.GetMysqlDbHandle()
  //if db == nil {
  //  seelog.Errorf("GetMysqlDbHandle failed, mysql init failed")
  //  return
  //}

  defer db.Close()
  //if db.AutoMigrate(&database.UserInfo{}, &database.ThemeCatalog{},&database.EconomicsRecordList{},&database.MilitaryRecordList{},
  //  &database.PersonageRecordList{}, &database.SportRecordList{}, &database.EntertainmentRecordList{}) == nil {
  //  seelog.Errorf("AutoMigrate failed, mysql init failed")
  //  return
  //}

  //start iris app
  if err := startIrisApp(cmd, args);err != nil {
    seelog.Errorf("startIrisApp failed, err is %v", err.Error())
    return
  }
}

func startIrisApp(cmd *cobra.Command, args []string) error{
  app := iris.New()
  app.Use(recover.New())
  app.Use(logger.New())

  //iris.WithConfiguration(iris.YAML("../config/config.yml"))

  Router.RegistryRoutes(app)

  if err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)); err != nil{
    seelog.Errorf("Run app failed, err is %v", err.Error())
    return err
  }

  return nil
}