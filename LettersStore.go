package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserLetters struct {
	Email string `json:"email"`
	Letters []Letter `json:"letters"`
}

var storePath = "store.json"

func createStoreIfNeed() error {
	fmt.Println("createIfNeed")
	_, err := os.Stat(storePath)
	if err != nil {
		if os.IsNotExist(err) {
			return ioutil.WriteFile(storePath, []byte("[]"), 0644)
		}
		return err
	}
	return nil
}

func GetAllLetters() ([]UserLetters, error) {
	err := createStoreIfNeed()
	fmt.Println("after createIfnedd", err)
	if err != nil {
		return nil, err
	}

	bytesStr, err := ioutil.ReadFile(storePath)
	if err != nil {
		return nil, err
	}

	store := []UserLetters{}
	err = json.Unmarshal(bytesStr, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func GetUserLetters(email string) ([]Letter, int, error) {
	store, err := GetAllLetters()
	if err != nil {
		return nil, -1, err
	}

	for i, userLetters := range store {
		if userLetters.Email == email {
			return userLetters.Letters, i, nil
		}
	}

	return []Letter{}, -1, err
}

func CreateOrUpdateUserLetter(email string, letter Letter) (error) {
	letters, index, err := GetUserLetters(email)
	if err != nil {
		return err
	}

	found := false
	for i, curLetter := range letters {
		if curLetter.ID == letter.ID {
			found = true
			letters[i].CheckedSubEvent = letter.CheckedSubEvent
			letters[i].Checked = true
		}
	}

	if !found {
		letters = append(letters, letter)
	}

	store, err := GetAllLetters()
	userLetters := UserLetters{Email:email, Letters:letters}
	if index == -1 {
		store = append(store, userLetters)
	} else {
		store[index] = userLetters
	}

	bytesToWrite, err := json.Marshal(store)
	if err != nil {
		return err
	}

	err = os.Remove(storePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(storePath, bytesToWrite, 0644)
	if err != nil {
		return err
	}

	return nil
}

