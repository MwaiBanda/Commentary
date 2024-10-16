package model

import "strconv"


type Board struct {
	ID          int     `json:"id"`
	Title       string    `json:"title"`
	Prayers    []Prayer  `json:"prayers"`
	Boards	 []Board   `json:"boards"`
	Clients    []Client `json:"clients"`
}

type BoardEvent struct {
	ID		  int     `json:"id"`
}
type BoardData struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Prayers    []Prayer  `json:"prayers"`
	Boards	 []Board   `json:"boards"`
}

func (b *Board) GetBoardData() BoardData {
	return BoardData{
		ID: b.ID,
		Title: b.Title,
		Prayers: b.Prayers,
		Boards: b.Boards,
	}
}

func (b *Board) AddClient(client Client) {
	client.SetBoardId(strconv.Itoa(b.ID))
	b.Clients = append(b.Clients, client)
}

func (b *Board) RemoveClient(client Client) {
	filtered := []Client{}
	for _, c := range b.Clients {
		if c.ID != client.ID {
			filtered = append(filtered, c)
		}
	}
	b.Clients = filtered
}
