package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"Week02/service/userService"

	pkgError "github.com/pkg/errors"
	)

func main() {
	_, err := userService.DoService()
	if err != nil {
		if errors.Is(pkgError.Cause(err), sql.ErrNoRows) {
			fmt.Println("Stack Info: \n%+v\n", err)
			os.Exit(1)
		}
	}
	
	fmt.Println("Ending")
}
