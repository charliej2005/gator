package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/charliej2005/gator/internal/config"
	"github.com/charliej2005/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, _ := config.Read()

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	s := state{db: dbQueries, config: &cfg}

	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

	args := os.Args
	if len(args) < 2 {
		err := errors.New("command not provided")
		fmt.Println(err)
		os.Exit(1)
	}
	userCmdName := args[1]
	var userArgs []string
	if len(args) > 2 {
		userArgs = args[2:]
	}

	userCmd := command{name: userCmdName, args: userArgs}
	err = cmds.run(&s, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type state struct {
	db     *database.Queries
	config *config.Config
}
