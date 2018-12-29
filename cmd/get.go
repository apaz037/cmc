// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"github.com/leekchan/accounting"
	api "github.com/miguelmota/go-coinmarketcap/v2"
	"github.com/spf13/cobra"
)

var ticker string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get pulls the latest price for a cryptocurrency by ticker",
	Example: "cmc get eth",
	Long: `Get pulls the latest price for a cryptocurrency by ticker`,
	Args: cobra.ExactArgs(1), // TODO: change validation in future as feature set expands
	RunE: func(cmd *cobra.Command, args []string) error {
		price, err := getPrice(args[0])
		if err != nil && price == 0 {
			return errors.New("Could not retrieve price for:" +  ticker)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getCmd.Flags().StringP("toggle", "t", false, "Help message for toggle")
}

func getPrice(ticker string) (float64, error) {
	priceOptions := &api.PriceOptions{
		Symbol: ticker,
		Convert: "USD",
	}

	quote, err := api.Price(priceOptions)
	if err != nil {
		return 0, err
	}

	ac := accounting.Accounting{
		Symbol: "$",
		Precision: 2,
	}

	fmt.Println(ticker, ":", ac.FormatMoney(quote))
	return quote, nil
}