package runeterra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alee792/runeterra/pkg/datadragon"
	"github.com/alee792/runeterra/proto"
	"github.com/pkg/errors"
)

const (
	StaticDecklistEndpoint = "/static-decklist"
	CardPositionsEndpoint  = "/positional-rectangles"
	GameResultEndpoint     = "/game-result"
)

// Client interacts with LoR endpoints.
type Client struct {
	hc     *http.Client
	cards  map[string]proto.Card
	Config Config
}

// Config for Client.
type Config struct {
	Address string
}

// NewClient instantiates a client with sensible defaults.
func NewClient(cfg Config) *Client {
	if cfg.Address == "" {
		cfg.Address = "http://localhost:21337"
	}

	c := &Client{
		hc:     http.DefaultClient,
		cards:  datadragon.Cards(),
		Config: cfg,
	}

	return c
}

// GetStaticDeclist describes the player's current deck in an active game.
func (c *Client) GetStaticDecklist(ctx context.Context) (*proto.Deck, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.Config.Address+StaticDecklistEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}

	var out *proto.Deck

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}

	return out, nil
}

// GetCardPositions returns the positions of cards at a given point in time.
func (c *Client) GetCardPositions(ctx context.Context) (*proto.CardPositions, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.Config.Address+CardPositionsEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}

	var out *proto.CardPositions

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}

	return out, nil
}

// GetGameResult returns a player's most recently completed game.
func (c *Client) GetGameResult(ctx context.Context) {}
