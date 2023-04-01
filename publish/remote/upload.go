package remote

import (
	"fmt"
	"log"
	"os"
	"path"
	"publish/config"

	"github.com/pkg/sftp"
)

func Upload() {
	var (
		err        error
		sftpClient *sftp.Client
	)
	sftpClient, err = GetSftpClient()
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()
	var localFilePath = config.FileRoute.Local
	var remoteDir = config.FileRoute.Remote
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()
	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()
	buf := make([]byte, 10*1024*1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	fmt.Println("copy file to remote server finished!")
}
