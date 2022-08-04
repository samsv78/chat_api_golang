package helpers

import (
	"errors"
	"strings"

	"github.com/gorilla/websocket"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}

func RemoveDuplicate[T string | int | uint32](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ContainsString(slice []string, element string) (bool, int) {
	for i, v := range slice {
		if v == element {
			return true, i
		}
	}
	return false, -1
}

func ContainsWSConnection(slice []*websocket.Conn, element *websocket.Conn) (bool, int) {
	for i, v := range slice {
		if v == element {
			return true, i
		}
	}
	return false, -1
}

func RemoveElementByIndex[T any](slice []T, index int) []T {
	sliceLen := len(slice)
	sliceLastIndex := sliceLen - 1

	if index != sliceLastIndex {
		slice[index] = slice[sliceLastIndex]
	}

	return slice[:sliceLastIndex]
}
