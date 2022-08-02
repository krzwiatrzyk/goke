/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"goke/config"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var module string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	runCmd.PersistentFlags().StringVarP(&module, "module", "m", "gin", "supported values: gin, http")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	switch module {
	case "gin":
		runGin()
	case "http":
		runHttp()
	}
}

func runGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/headers", returnHeaders)

	r.Run(fmt.Sprintf(":%d", serverConfig.Port))
}

func returnHeaders(c *gin.Context) {
	c.JSON(http.StatusOK, c.Request.Header)
}

func runHttp() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/templates", templateFile)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(fmt.Sprintf(":%d", serverConfig.Port), nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// func hello2(w http.ResponseWriter, req *http.Request) {

//     test := config.Config{"test","test123"}
//     fmt.Println(test)

//     ctx := req.Context()
//     fmt.Println("server: hello handler started")
//     defer fmt.Println("server: hello handler ended")

//     select {
//     case <-time.After(10 * time.Second):
//         fmt.Fprintf(w, "hello\n")
//     case <-ctx.Done():

//         err := ctx.Err()
//         fmt.Println("server:", err)
//         internalError := http.StatusInternalServerError
//         http.Error(w, err.Error(), internalError)
//     }
// }

func templateFile(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html.gotpl"))
	cfg := config.Config{"test_tile", "0.0.1-dev"}
	tmpl.Execute(w, cfg)
}
