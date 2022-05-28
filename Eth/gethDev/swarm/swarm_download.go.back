package main

import (
	"fmt"
	"io/ioutil"
	"log"

	bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
)

func swarmDownload() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")
	manifestHash := "c8a1abb17f7f8b5bbb0b979e962c2722c92bc8bb90b63c26feef69049342f302"

	// 先通过调用“DownloadManfest”来下载它，并检查清单的内容
	manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isEncrypted) // false

	// 我们可以遍历清单条目，看看内容类型，大小和内容哈希是什么
	for _, entry := range manifest.Entries {
		fmt.Println(entry.Hash)        // 92672a471f4419b255d7cb0cf313474a6f5856fb347c5ece85fb706d644b630f
		fmt.Println(entry.ContentType) // text/plain; charset=utf-8
		fmt.Println(entry.Size)        // 11
		fmt.Println(entry.Path)        // ""
	}

	// 如果您熟悉swarm url，它们的格式为bzz：/ <hash> / <path>，
	// 因此为了下载文件，我们指定了清单哈希和路径.在这个例子里，路径是一个
	// 空字符串.我们将这些数据传递给Download函数并返回一个文件对象
	file, err := client.Download(manifestHash, "")
	if err != nil {
		log.Fatal(err)
	}

	// 我们现在可以阅读并打印返回的文件阅读器的内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content)) // hello world
}
