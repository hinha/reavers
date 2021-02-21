package main

import (
	"fmt"
	"github.com/adlio/trello"
)

func main() {
	appKey := "c48b97cead92fca9f4f180713b6cbfd6"
	appToken := "7017d8dc00f675abc3ca0c086c27772e648f5424725b4ee86e86fbb182951edc"
	client := trello.NewClient(appKey, appToken)
	board, err := client.GetBoard("9fdlFEda", trello.Defaults())
	if err != nil {
		panic(err)
	}
	members, err := board.GetMembers(trello.Defaults())
	if err != nil {
		panic(err)
	}
	fmt.Println(board.IDOrganization)
	for _, member := range members {
		fmt.Println(member.Email, member.FullName, member.Username)
	}

}
