package action

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli"
)

// Remove sadakoを実行したrepositoryから削除する
func Remove(c *cli.Context) {
	// When not git repository.
	found := exists(".git")
	if !found {
		fmt.Println("This directory is not git repository yet :(")
		os.Exit(1)
	}

	var sadako = ".git/hooks/post-merge"
	sadakoFound := exists(sadako)
	if !sadakoFound {
		fmt.Println("sadako speak binary is not found.\n" +
			"Where did she go... :p")
		os.Exit(1)
	}
	if err := os.Remove(sadako); err != nil {
		panic(err)
	}

	fmt.Println("Sadako has gone :)")

}

// Set .git/hooks/post-mergeにコピーする
func Set(c *cli.Context) {
	sadakoSrcPath := os.Getenv("GOPATH") + "/src/github.com/konojunya/sadako"

	// When not git repository.
	found := exists(".git")
	if !found {
		fmt.Println("This directory is not git repository yet :(")
		os.Exit(1)
	}

	// When not found source.
	found = exists(sadakoSrcPath)
	if !found {
		fmt.Println("sadako speak binary is not found.\n" +
			"Please `go get github.com/konojunya/sadako`")
		os.Exit(1)
	}

	// first: sadako-speak
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err = copy(sadakoSrcPath+"/speak", path+"/.git/hooks/post-merge"); err != nil {
		panic(err)
	}
	// 実行権限を与える
	os.Chmod(path+"/.git/hooks/post-merge", 0755)

	fmt.Println("Sadako is near you... :)")
}

func copy(srcpath, distpath string) error {
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()

	dist, err := os.Create(distpath)
	if err != nil {
		return err
	}
	defer dist.Close()

	_, err = io.Copy(dist, src)
	if err != nil {
		return err
	}

	return nil
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
