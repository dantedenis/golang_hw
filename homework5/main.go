package main

import (
	"flag"
	"fmt"
	"homework5/crypto"
	"log"
)

func main() {
	var fileSource, hashFile, outFile, signFile string
	flag.StringVar(&fileSource, "source-file", "", "Source File")
	flag.StringVar(&hashFile, "hash-file", "", "File hash")
	flag.StringVar(&outFile, "out-file", "out", "Output file")
	flag.StringVar(&signFile, "sign-file", "", "File with signature")
	flag.Parse()

	switch flag.Args()[0] {
	case "enc":
		encoder, err := crypto.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = encoder.EncryptSha256()
		if err != nil {
			panic(err)
		}
		err = encoder.SaveToFile(outFile)
		if err != nil {
			panic(err)
		}
	case "dec":
		encoder, err := crypto.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = encoder.EncryptSha256()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(encoder.GetSign().SignatureBytes()))
		_, err = encoder.GetSign().Equal(signFile)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			fmt.Println("File is ok!")
		}
	default:
		log.Fatalln("Please use \"enc\" or \"dec\" param")
	}
}
