package solitairestore

import (
	"fmt"
	"encoding/json"
	"os"
	"errors"

	"solitaire/game"
	"solitaire/game/gamestates"

	"github.com/rs/xid"
)

type fileStore struct {}

func (fs *fileStore) SaveGame(currentGameState game.Game, previousGameStates gamestates.GameStates) (string, error) {
	// Generate a random identifier for the game
	gameId := xid.New().String()
	// Create the filename
	fileName := gameId + ".json"
	// Create GameStateData object
	gameStateData := GameStateData{currentGameState, previousGameStates}
	// Marshal the game data to JSON
	jsonData, err := json.Marshal(gameStateData)
	if err != nil {
		return "", fmt.Errorf("error marshaling game data: %w", err)
	}

	// Write the JSON data to the file
	filePath := fmt.Sprintf("%s/%s", "gameinstances", fileName)
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return "", fmt.Errorf("error writing game data to file: %w", err)
	}

	return gameId, nil
}

func (fs *fileStore) LoadGame(gameId string) (game.Game, gamestates.GameStates, error) {
	// Create the filename
	fileName := gameId + ".json"
	// Read the JSON data from the file
	filePath := fmt.Sprintf("%s/%s", "gameinstances", fileName)
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return game.Game{}, gamestates.GameStates{}, fmt.Errorf("game with ID %s not found", gameId)
		}
		return game.Game{}, gamestates.GameStates{}, fmt.Errorf("error reading game data from file: %w", err)
	}

	// Unmarshal the JSON data into a GameStateData object
	var gameStateData GameStateData
	err = json.Unmarshal(jsonData, &gameStateData)
	if err != nil {
		return game.Game{}, gamestates.GameStates{}, fmt.Errorf("error unmarshalling game data: %w", err)
	}

	return gameStateData.CurrentGameState, gameStateData.PreviousGameStates, nil
}