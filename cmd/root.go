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
  "context"
  "fmt"
  "github.com/cihub/seelog"
  "github.com/hfeng101/niwo/storage"
  "github.com/hfeng101/niwo/utils/config"
  "github.com/hfeng101/niwo/utils/logger"
  "github.com/spf13/cobra"
  "os"
  "os/signal"
  "runtime"
  "syscall"
  "time"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

  "github.com/hfeng101/niwo/router"

  "github.com/kataras/iris/v12"
  IrisLogger "github.com/kataras/iris/v12/middleware/logger"
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
  seelog.Infof("starting niwo...")

  //init iris config
  config.InitIrisConfig()

  //init mysql & create table
  if err := storage.InitGlobalOrm();err != nil{
    seelog.Errorf("InitGlobalOrm failed, err is %v", err.Error())
    return
  }
  mysqlHandle := storage.GetMysqlDbHandle()
  if mysqlHandle == nil {
    seelog.Errorf("GetMysqlDbHandle failed, mysql init failed")
    return
  }
  defer mysqlHandle.Close()
  if mysqlHandle.AutoMigrate(
      &storage.UserInfo{},
      &storage.ThemeCatalog{},
      &storage.EconomicsRecordList{},
      &storage.MilitaryRecordList{},
      &storage.PersonageRecordList{},
      &storage.SportRecordList{},
      &storage.EntertainmentRecordList{}) == nil {
    seelog.Errorf("AutoMigrate failed")
    return
  }

  seelog.Infof("InitMongoDb starting!")
  //init mongoDB
  if err := storage.InitMongoDb();err != nil {
    seelog.Errorf("InitMongoDb failed, err is %v", err.Error())
    return
  }

  seelog.Infof("InitMongoDb success!")

  //start iris app
  exitChan := make(chan struct{})
  if err := startIrisApp(exitChan);err != nil {
    seelog.Errorf("startIrisApp failed, err is %v", err.Error())
    return
  }

  <-exitChan
}

func startIrisApp(exitChan chan struct{}) error{
  app := iris.New()
  app.Use(recover.New())
  app.Use(IrisLogger.New())

  //iris.WithConfiguration(iris.YAML("../config/config.yml"))
  ctx, cancelFunc := context.WithCancel(context.Background())

  //Catch the INT & TERM signal and overwrite the default action.
  go gracefulExit(cancelFunc, app, ctx, exitChan, 10*time.Second)

  //Catch USR1 signal to all Gs' stack.
  go runtimeStackInfo(ctx)

  //Catch HUP system signal to reload the configuration file.
  go reloadConfiguration(ctx)

  //Run iris server
  Router.RegistryRoutes(app)
  if err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)); err != nil{
    seelog.Errorf("Run app failed, err is %v", err.Error())
    return err
  }

  return nil
}

func reloadConfiguration(ctx context.Context) {
  hupSigChan := make(chan os.Signal)
  //catch HUP signal
  signal.Notify(hupSigChan, syscall.SIGHUP)

  for {
    select {
      case _, ok := <-hupSigChan:
        if ok {
          config.InitIrisConfig()
          //logger setting
          switch config.GetConfig().Other["log-level"].(string) {
            case "debug":
              logger.SwitchLoggerLevel("debug")
            case "info":
              logger.SwitchLoggerLevel("info")
            case "warn":
              logger.SwitchLoggerLevel("warn")
            case "error":
              logger.SwitchLoggerLevel("error")
            default:
              logger.SwitchLoggerLevel("info")
          }

          //reload mysql
          storage.InitGlobalOrm()

          //reload mongodb
          storage.InitMongoDb()
        }
      case <-ctx.Done():
        return
      }
  }
}

//优雅退出
func gracefulExit(cancelFunc context.CancelFunc, app *iris.Application, ctx context.Context, exitChan chan struct{}, waitPeriod time.Duration) {
  defer seelog.Flush()
  var exitSignalCnt int
  exitSignalChan := make(chan os.Signal)
  //catch INT and TERM signal
  signal.Notify(exitSignalChan, syscall.SIGINT, syscall.SIGTERM)
  for {
    select {
    case s, ok := <-exitSignalChan:
      if ok {
        exitSignalCnt++
        seelog.Debugf("The %vth signal recved: %v", exitSignalCnt, s)
        if exitSignalCnt < 2 {
          seelog.Infof("waiting for running process to finish before shutting down, take care to press ctrl+c to quit abruptly")
          go func() {
            cancelFunc()
            seelog.Infof("main process exit %v later", waitPeriod)
            app.Shutdown(ctx)
            time.Sleep(waitPeriod)
            seelog.Infof("quit gracefully")
            exitChan <- struct{}{}
          }()
        } else {
          seelog.Infof("quit abruptly")
          exitChan <- struct{}{}
        }
      }
    }
  }
}

func runtimeStackInfo(ctx context.Context) {
  user1SigChan := make(chan os.Signal)
  //catch USER1 signal
  signal.Notify(user1SigChan, syscall.SIGUSR1)
  for {
    select {
    case _, ok := <-user1SigChan:
      if ok {
        stackBuff := make([]byte, 1<<16)
        n := runtime.Stack(stackBuff, true)
        fmt.Println(string(stackBuff[:n]))
      }
    case <-ctx.Done():
      return
    }
  }
}