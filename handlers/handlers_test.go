package handlers

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const nearlyFinishedGameToken = "qlYqKlWyUjIxIhBShoawsEK4JNwEYa2hiVItIAAA__8"

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
		{name: "svg", handler: SvgHandler, target: "/game.svg?s=not-a-board", want: "error deserializing game state:"},
		{name: "svg missing token", handler: SvgHandler, target: "/game.svg", want: "error getting token: no game found"},
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
