package runeterra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	var out *proto.Deck
	if err := c.get(ctx, c.Config.Address+StaticDecklistEndpoint, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetCardPositions returns the positions of cards at a given point in time.
func (c *Client) GetCardPositions(ctx context.Context) (*proto.CardPositions, error) {
	var out *proto.CardPositions
	if err := c.get(ctx, c.Config.Address+CardPositionsEndpoint, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetGameResult returns a player's most recently completed game.
func (c *Client) GetGameResult(ctx context.Context) (*proto.GameResult, error) {
	var out *proto.GameResult
	if err := c.get(ctx, c.Config.Address+GameResultEndpoint, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) get(ctx context.Context, address string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return errors.Wrap(err, "failed to read request body")
	}

	return nil
}

type Game struct {
	ID           int32
	StartdAt     time.Time
	EndAt        time.Time
	PlayerName   string
	OpponentName string
	Won          bool
	DeckCode     string
}

func (c *Client) RecordGame(ctx context.Context) error {
	const (
		inProgress = "InProgress"
		menus      = "Menus"
	)

	var (
		curG *Game
		// Arbitrary interval.
		ticker = time.NewTicker(time.Second)
	)

	for {
		select {
		case <-ticker.C:
			// Check if in active game.
			cp, err := c.GetCardPositions(ctx)
			if err != nil {
				return err
			}

			switch {
			// Existing game -> continue polling until game completed.
			// Maybe track game state later.
			case cp.GameState == inProgress && curG != nil:
				continue

			// New game -> Gather one time deck and opponent information.
			case cp.GameState == inProgress && curG == nil:
				curG.Player = cp.GetPlayerName()
				curG.Opponent = cp.GetOpponentName()
				curG.StartdAt = time.Now()

				// Get current game info.
				d, err := c.GetStaticDecklist(ctx)
				if err != nil {
					return err
				}

				// Oops, we must have just missed it.
				if d == nil {
					continue
				}

				curG.DeckCode = d.GetDeckCode()

			// Game completed
			case cp.GameState == menus && curG != nil:
				curG.EndAt = time.Now()

				gr, err := c.GetGameResult(ctx)
				if err != nil {
					return err
				}

				curG.ID = gr.GetGameId()
				curG.Won = gr.GetLocalPlayerWon()
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
