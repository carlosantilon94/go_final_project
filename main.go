package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Item struct {
	Id string
	Name string
}

func main()  {
	ctx := context.Background()
	id := uuid.NewString()
	fmt.Println(id)
	item := Item{
		Id:   id,
		Name: "Carlos Prueba 4",
	}
	InsertData(ctx,item)
	fmt.Printf("%+#v",GetData(ctx, id))

}
