package command

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"global"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"log"
	"math"
	"os"
	"tool"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	Base64EncryptionCommand = &cli.Command{
		Name:        "base64",
		Usage:       "Encrypt with base64 algorithm",
		Description: "Encrypt with base64 algorithm",
		Category:    global.CategoryEncrypt,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Usage: "encrypt a specified file"},
			&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "decrypt encoded Base64 string"},
		},
		Action: execBase64EncryptionCommand,
	}

	Md5EncryptionCommand = &cli.Command{
		Name:        "md5",
		Usage:       "Encrypt with md5 algorithm",
		Description: "Encrypt with md5 algorithm",
		Category:    global.CategoryEncrypt,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Usage: "encrypt a specified file"},
		},
		Action: execMd5EncryptionCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================
func execBase64EncryptionCommand(ctx *cli.Context) error {

	argValue := ctx.Args().Get(0)
	// check
	if argValue == "" && ctx.String("file") == "" {
		fmt.Println("Value not specified")
		return nil
	}

	// file compute
	if ctx.String("file") != "" {

		filePath := ctx.String("file")
		if !tool.Exists(filePath) || !tool.IsFile(filePath) {
			return cli.Exit("No such file or not a regular file", -1)
		}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(base64.StdEncoding.EncodeToString(data))
		// string compute
	} else {

		var result string
		// decrypt encoded Base64 string
		if ctx.Bool("reverse") {

			// decrypt a file Base64 string
			if ctx.String("output") != "" {

				outputFilePath := ctx.String("output")
				data, _ := base64.StdEncoding.DecodeString(argValue)
				err := ioutil.WriteFile(outputFilePath, data, 0666)
				if err != nil {
					return cli.Exit(err.Error(), -1)
				}
				fmt.Println(outputFilePath)
				return nil
			}

			// decrypt a pure Base64 string
			bytes, err := base64.StdEncoding.DecodeString(argValue)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			result = string(bytes[:])
			// encode
		} else {
			result = base64.StdEncoding.EncodeToString([]byte(argValue))
		}
		fmt.Println(result)
	}

	return nil
}

const fileChunk = 8192 // 8KB

func execMd5EncryptionCommand(ctx *cli.Context) error {

	argValue := ctx.Args().Get(0)
	// check
	if argValue == "" && ctx.String("file") == "" {
		fmt.Println("Value not specified")
		return nil
	}

	md5Ctx := md5.New()
	// file compute
	if ctx.String("file") != "" {

		filePath := ctx.String("file")
		if !tool.Exists(filePath) || !tool.IsFile(filePath) {
			return cli.Exit("No such file or not a regular file", -1)
		}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		defer file.Close()

		info, _ := file.Stat()
		fileSize := info.Size()
		blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		for i := uint64(0); i < blocks; i++ {
			blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			buf := make([]byte, blockSize)
			file.Read(buf)
			md5Ctx.Write(buf)
		}
		// string compute
	} else {
		data := []byte(argValue)
		md5Ctx.Write(data)
	}

	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(hex.EncodeToString(cipherStr))

	return nil
}
