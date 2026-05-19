package main

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher/callback"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

func main() {
	// Register callback
	eventHandler := dispatcher.NewEventDispatcher("", "").
		// Listen for "card action callback card.action.trigger"
		OnP2CardActionTrigger(func(ctx context.Context, event *callback.CardActionTriggerEvent) (*callback.CardActionTriggerResponse, error) {
			fmt.Printf("[ OnP2CardActionTrigger access ], data: %s\n", larkcore.Prettify(event))

			return nil, nil
		}).
		// Listen for "fetch link preview data url.preview.get"
		OnCustomizedEvent("application.bot.menu_v6", func(ctx context.Context, event *larkevent.EventReq) error {
			fmt.Printf("[ OnCustomizedEvent access ], type: message, data: %s\n", string(event.Body))
			return nil
		})
	// Create Client
	cli := larkws.NewClient("cli_a76cab1b75c71013", "dTBxFIYVsMHa1dczZtscXfeVEch63sYn",
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)
	// Establish persistent connection
	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
