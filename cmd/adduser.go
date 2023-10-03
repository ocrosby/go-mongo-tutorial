/*
Copyright © 2023 Omar Crosby <omar.crosby@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/ocrosby/go-mongo-tutorial/pkg/dal"
	"github.com/ocrosby/go-mongo-tutorial/pkg/models"
	"log"

	"github.com/spf13/cobra"
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		firstName, _ := cmd.Flags().GetString("first-name")
		lastName, _ := cmd.Flags().GetString("last-name")

		userModel := dal.NewUserModel(env)

		user := &models.User{
			FirstName: firstName,
			LastName:  lastName,
		}

		log.Printf("Creating user: %v\n", user)

		if user, err = userModel.Create(user); err != nil {
			panic(err)
		}

		log.Printf("Created user: %v\n", user)
	},
}

func init() {
	rootCmd.AddCommand(adduserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adduserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	adduserCmd.Flags().StringP("first-name", "f", "", "First Name")
	adduserCmd.MarkFlagRequired("first-name")

	adduserCmd.Flags().StringP("last-name", "l", "", "Last Name")
	adduserCmd.MarkFlagRequired("last-name")
}
