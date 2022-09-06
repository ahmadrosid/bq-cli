package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ahmadrosid/bq-cli/service"
	"github.com/google/pprof/driver"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

type Data map[string]interface{}

type bigqueryCcommand struct {
	svc service.BiqueryService
	ui  driver.UI
}

func NewBiqueryCommand(svc service.BiqueryService, ui driver.UI) *bigqueryCcommand {
	return &bigqueryCcommand{
		svc: svc,
		ui:  ui,
	}
}

func (b *bigqueryCcommand) printTable(res []string) {
	json_res := fmt.Sprintf("[%s]", strings.Join(res, ","))
	dec := json.NewDecoder(strings.NewReader(json_res))
	var response []Data
	dec.Decode(&response)

	header := make([]string, 0)

	for key := range response[0] {
		header = append(header, key)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, items := range response {
		data := make([]string, 0)
		for _, key := range header {
			val := items[key]
			value := ""
			if val == nil {
				value = "null"
			} else {
				value = toString(val)
			}
			data = append(data, value)
		}
		table.Append(data)
	}
	table.Render()
}

func (b *bigqueryCcommand) HandleQuery(ctx *cli.Context) error {
	query := ctx.Args().First()
	if query == "" {
		return fmt.Errorf("Query text is required!")
	}

	res, err := b.svc.GetDataFromBQ(query, ctx.Context)
	if err != nil {
		return err
	}

	// Useful for debuging
	// res := []string{
	// 	`{"created_at":"2017-12-15T08:21:15Z","id":"fb1f456b-e1bb-4723-9075-ec67bc74433b"}`,
	// 	`{"created_at":"2017-12-20T05:06:08Z","id":"274670ca-bd6a-40fd-8e3b-9d9d1441c32b"}`,
	// 	`{"created_at":"2020-07-24T08:18:55Z","id":"03c13e25-8fc1-49ee-a8e7-2c0c1fcc94e8"}`,
	// }

	b.printTable(res)

	return nil
}

func (b *bigqueryCcommand) HandleInteractive(ctx *cli.Context) error {
	for {
		query, err := b.ui.ReadLine("> ")
		if err != nil {
			b.ui.PrintErr(err)
			break
		}

		if query == "exit" {
			break
		}

		if query == "" {
			b.ui.Print("> ")
			continue
		}

		res, err := b.svc.GetDataFromBQ(query, ctx.Context)
		if err != nil {
			b.ui.PrintErr(err)
			b.ui.Print("> ")
			continue
		}

		b.printTable(res)
		b.ui.Print("> ")
	}
	return nil
}

func toString(v interface{}) string {
	switch s := v.(type) {
	case string:
		return string(s)
	case int:
		return fmt.Sprintf("%d", s)
	case float32:
		return fmt.Sprintf("%.0f", s)
	case float64:
		return fmt.Sprintf("%.0f", s)
	default:
		return fmt.Sprintf("%v", s)
	}
}
