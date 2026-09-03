package handlers

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	bgjson "github.com/chandler37/gobackgammon/json"
)

const nearlyFinishedGameToken = "qlYqKlWyUjIxIhBShoawsEK4JNwEYa2hiVItIAAA__8"

// This completed board follows the pinned gobackgammon v0.1.7 dependency's
// TestTakeTurnSingleStakes construction: White has borne off all 15 checkers,
// while Red has borne off 13 and has two on point 17. The handler accepts the
// dependency's uncompressed serialization format as an existing public input.
const deterministicWhiteVictorySerialization = `{"p":"W","p0":"W15","p17":"r2","p25":"r13"}`

func TestRootHandlerContract(t *testing.T) {
	recorder := httptest.NewRecorder()
	RootHandler(recorder, httptest.NewRequest(http.MethodGet, "/", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	for _, want := range []string{
		"<title>Backgammon</title>",
		`href="/game"`,
		"github.com/chandler37/gobackgammond",
	} {
		if !strings.Contains(recorder.Body.String(), want) {
			t.Errorf("response body does not contain %q", want)
		}
	}
}

func TestTokenContract(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		wantToken string
		wantError string
	}{
		{name: "missing", target: "/game", wantError: noGameFoundError.Error()},
		{name: "single", target: "/game?s=board-state", wantToken: "board-state"},
		{name: "empty", target: "/game?s=", wantError: "empty game found"},
		{name: "multiple", target: "/game?s=one&s=two", wantError: "too many games found"},
		{name: "malformed query", target: "/game?s=%zz", wantError: "bad URL:"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := requestWithRawQuery(test.target)
			gotToken, err := token(request)
			if gotToken != test.wantToken {
				t.Errorf("token = %q, want %q", gotToken, test.wantToken)
			}
			if test.wantError == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), test.wantError) {
				t.Fatalf("error = %v, want text containing %q", err, test.wantError)
			}
		})
	}
}

func TestShouldTakeTurnContract(t *testing.T) {
	tests := []struct {
		target string
		want   bool
	}{
		{target: "/game?s=state", want: false},
		{target: "/game?s=state&t=", want: true},
		{target: "/game?s=state&t=anything", want: true},
		{target: "/game?s=state&t=one&t=two", want: true},
		{target: "/game?s=%zz&t=", want: false},
	}

	for _, test := range tests {
		request := requestWithRawQuery(test.target)
		if got := shouldTakeTurn(request); got != test.want {
			t.Errorf("shouldTakeTurn(%q) = %t, want %t", test.target, got, test.want)
		}
	}
}

func requestWithRawQuery(target string) *http.Request {
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	if queryStart := strings.IndexByte(target, '?'); queryStart >= 0 {
		request.URL.RawQuery = target[queryStart+1:]
	}
	return request
}

func TestGameHandlerNewGameContract(t *testing.T) {
	rand.Seed(37)
	recorder := httptest.NewRecorder()
	GameHandler(recorder, httptest.NewRequest(http.MethodGet, "/game", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	for _, want := range []string{
		"<title>Backgammon Game</title>",
		"The current board is",
		"Which of the following is your play?",
		"/game.svg?s=",
	} {
		if !strings.Contains(recorder.Body.String(), want) {
			t.Errorf("response body does not contain %q", want)
		}
	}
}

func TestHandlersRejectInvalidGameState(t *testing.T) {
	for _, test := range []struct {
		name    string
		handler http.HandlerFunc
		target  string
		want    string
	}{
		{name: "game", handler: GameHandler, target: "/game?s=not-a-board", want: "error deserializing game state:"},
		{name: "game multiple tokens", handler: GameHandler, target: "/game?s=one&s=two", want: "error getting token: too many games found"},
		{name: "game malformed query", handler: GameHandler, target: "/game?s=%zz", want: "error getting token: bad URL:"},
		{name: "svg", handler: SvgHandler, target: "/game.svg?s=not-a-board", want: "error deserializing game state:"},
		{name: "svg missing token", handler: SvgHandler, target: "/game.svg", want: "error getting token: no game found"},
		{name: "svg multiple tokens", handler: SvgHandler, target: "/game.svg?s=one&s=two", want: "error getting token: too many games found"},
		{name: "svg malformed query", handler: SvgHandler, target: "/game.svg?s=%zz", want: "error getting token: bad URL:"},
	} {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			test.handler(recorder, httptest.NewRequest(http.MethodGet, test.target, nil))
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
			}
			if !strings.Contains(recorder.Body.String(), test.want) {
				t.Errorf("response body = %q, want text containing %q", recorder.Body.String(), test.want)
			}
		})
	}
}

func TestHandlersAcceptValidUncompressedGameState(t *testing.T) {
	serialization, err := bgjson.Decompress(nearlyFinishedGameToken)
	if err != nil {
		t.Fatalf("decompress fixture: %v", err)
	}
	token := url.QueryEscape(serialization)

	for _, test := range []struct {
		name        string
		handler     http.HandlerFunc
		target      string
		contentType string
		want        string
	}{
		{name: "game", handler: GameHandler, target: "/game?s=" + token, want: "<title>Backgammon Game</title>"},
		{name: "svg", handler: SvgHandler, target: "/game.svg?s=" + token, contentType: "image/svg+xml", want: "<svg"},
	} {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(37)
			recorder := httptest.NewRecorder()
			test.handler(recorder, httptest.NewRequest(http.MethodGet, test.target, nil))

			if recorder.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body = %s", recorder.Code, http.StatusOK, recorder.Body.String())
			}
			if test.contentType != "" && recorder.Header().Get("Content-Type") != test.contentType {
				t.Errorf("Content-Type = %q, want %q", recorder.Header().Get("Content-Type"), test.contentType)
			}
			if !strings.Contains(recorder.Body.String(), test.want) {
				t.Errorf("response body does not contain %q", test.want)
			}
		})
	}
}

func TestSmartBoardsAIAndSerializationContract(t *testing.T) {
	rand.Seed(37)
	serialization, err := bgjson.Decompress(nearlyFinishedGameToken)
	if err != nil {
		t.Fatalf("decompress fixture: %v", err)
	}
	board, err := bgjson.Deserialize(serialization)
	if err != nil {
		t.Fatalf("deserialize fixture: %v", err)
	}

	got, err := smartBoards(board.LegalContinuations())
	if err != nil {
		t.Fatalf("smartBoards: %v", err)
	}
	if len(got) == 0 {
		t.Fatal("smartBoards returned no AI-ranked continuations")
	}
	if !strings.HasPrefix(got[0].Hint, "PlayerConservative's choice") {
		t.Errorf("first hint = %q, want PlayerConservative choice marker", got[0].Hint)
	}
	for i, candidate := range got {
		if candidate.Board == nil {
			t.Fatalf("candidate %d has a nil board", i)
		}
		serialized, err := bgjson.Decompress(candidate.Serialization)
		if err != nil {
			t.Fatalf("candidate %d is not valid compressed serialization: %v", i, err)
		}
		decoded, err := bgjson.Deserialize(serialized)
		if err != nil || decoded == nil {
			t.Fatalf("candidate %d does not deserialize: board=%v err=%v", i, decoded, err)
		}
	}
}

func TestDynamicTemplateValuesAreEscaped(t *testing.T) {
	malicious := `"><script>alert("x")</script>`

	var game bytes.Buffer
	err := gameTemplate.Execute(&game, state{
		Board: smartBoard{Serialization: malicious},
		LegalContinuations: []smartBoard{{
			Serialization: malicious,
			Hint:          malicious,
		}},
	})
	if err != nil {
		t.Fatalf("execute game template: %v", err)
	}
	if strings.Contains(game.String(), "<script>") {
		t.Fatalf("game template emitted executable input: %s", game.String())
	}

	var victory bytes.Buffer
	err = victoryTemplate.Execute(&victory, victoryState{
		Serialization: malicious,
		Stakes:        malicious,
		Victor:        malicious,
		Score:         malicious,
	})
	if err != nil {
		t.Fatalf("execute victory template: %v", err)
	}
	if strings.Contains(victory.String(), "<script>") {
		t.Fatalf("victory template emitted executable input: %s", victory.String())
	}
}

func TestHandlersAcceptValidCompressedGameState(t *testing.T) {
	rand.Seed(37)

	t.Run("game", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		GameHandler(recorder, httptest.NewRequest(http.MethodGet, "/game?s="+nearlyFinishedGameToken, nil))

		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d; body = %s", recorder.Code, http.StatusOK, recorder.Body.String())
		}
		for _, want := range []string{"<title>Backgammon Game</title>", "/game.svg?s="} {
			if !strings.Contains(recorder.Body.String(), want) {
				t.Errorf("response body does not contain %q", want)
			}
		}
	})

	t.Run("svg", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		SvgHandler(recorder, httptest.NewRequest(http.MethodGet, "/game.svg?s="+nearlyFinishedGameToken, nil))

		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d; body = %s", recorder.Code, http.StatusOK, recorder.Body.String())
		}
		if got := recorder.Header().Get("Content-Type"); got != "image/svg+xml" {
			t.Errorf("Content-Type = %q, want %q", got, "image/svg+xml")
		}
		for _, want := range []string{"<?xml version=\"1.0\"?>", "<svg", "</svg>"} {
			if !strings.Contains(recorder.Body.String(), want) {
				t.Errorf("response body does not contain %q", want)
			}
		}
	})
}

func TestGameHandlerTakeTurnContract(t *testing.T) {
	rand.Seed(37)
	recorder := httptest.NewRecorder()
	GameHandler(recorder, httptest.NewRequest(http.MethodGet, "/game?s="+nearlyFinishedGameToken+"&t=", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	body := recorder.Body.String()
	if !strings.Contains(body, "<title>Backgammon Game</title>") &&
		!strings.Contains(body, "<title>Backgammon Game Results</title>") {
		t.Errorf("response body is neither a continued game nor a victory result: %q", body)
	}
}

func TestGameHandlerDeterministicVictoryContract(t *testing.T) {
	recorder := httptest.NewRecorder()
	target := "/game?s=" + url.QueryEscape(deterministicWhiteVictorySerialization) + "&t="
	GameHandler(recorder, httptest.NewRequest(http.MethodGet, target, nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	for _, want := range []string{
		"<title>Backgammon Game Results</title>",
		"The final score is White:1, Red:0.",
		"Congratulations on winning 1 point, White!",
		"/game.svg?s=",
	} {
		if !strings.Contains(recorder.Body.String(), want) {
			t.Errorf("response body does not contain %q", want)
		}
	}
}
