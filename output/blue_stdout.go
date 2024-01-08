package output

import (
	"bytes"
	"context"

	"github.com/benthosdev/benthos/v4/public/service"
	storage_go "github.com/supabase-community/storage-go"
)

func init() {
	err := service.RegisterOutput(
		"s_supabase", service.NewConfigSpec(),
		func(conf *service.ParsedConfig, mgr *service.Resources) (out service.Output, maxInFlight int, err error) {
			return &supabaseOutput{}, 1, nil
		})
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------

type supabaseOutput struct {
	client *storage_go.Client
}

func (output *supabaseOutput) Connect(ctx context.Context) error {
	// URL da API do Supabase
	supabaseUrl := "https://oodqfgyraszhlkxfpyix.supabase.co"
	// Chave de serviço do Supabase
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im9vZHFmZ3lyYXN6aGxreGZweWl4Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDM3ODgxMzgsImV4cCI6MjAxOTM2NDEzOH0.rgxURjxGIjw1kceuM3WywUY_-GrWwwtMdHmWK9gTL4s"

	client := storage_go.NewClient(supabaseUrl, supabaseKey, nil)
	output.client = client

	// fileBody := ... // load your file here

	return nil
}

func (output *supabaseOutput) Write(ctx context.Context, msg *service.Message) error {
	content, err := msg.AsBytes()
	if err != nil {
		return err
	}
	// convert byte slice to io.Reader
	reader := bytes.NewReader(content)

	_, err = output.client.UploadFile("bemthos", "test.json", reader)
	if err != nil {
		return err
	}
	return nil
}

func (output *supabaseOutput) Close(ctx context.Context) error {
	return nil
}
